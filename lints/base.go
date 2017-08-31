package lints

import (
	"time"

	"github.com/zmap/zcrypto/x509"
)

// global
var (
	Lints map[string]*Lint = make(map[string]*Lint)
)

const ZLintVersion = 3

type ZLintResult struct {
	ZLintVersion    int64                   `json:"version"`
	Timestamp       int64                   `json:"timestamp"`
	ZLint           map[string]ResultStruct `json:"lints"`
	NoticesPresent  bool                    `json:"notices_present"`
	WarningsPresent bool                    `json:"warnings_present"`
	ErrorsPresent   bool                    `json:"errors_present"`
	FatalsPresent   bool                    `json:"fatals_present"`
}

func (result *ZLintResult) Execute(cert *x509.Certificate) error {
	result.ZLint = make(map[string]ResultStruct, len(Lints))
	for name, l := range Lints {
		res, _ := l.ExecuteTest(cert)
		result.ZLint[name] = res
		result.updateErrorStatePresent(res)
	}
	return nil
}

func (zlintResult *ZLintResult) updateErrorStatePresent(result ResultStruct) {
	switch result.Result {
	case Notice:
		zlintResult.NoticesPresent = true
	case Warn:
		zlintResult.WarningsPresent = true
	case Error:
		zlintResult.ErrorsPresent = true
	case Fatal:
		zlintResult.FatalsPresent = true
	}
}

type LintTest interface {
	// runs once globally
	Initialize() error
	CheckApplies(c *x509.Certificate) bool
	RunTest(c *x509.Certificate) (ResultStruct, error)
}

type Lint struct {
	Name          string
	Description   string
	Source        string
	EffectiveDate time.Time
	Test          LintTest
}

func (l *Lint) ExecuteTest(cert *x509.Certificate) (ResultStruct, error) {
	if !l.Test.CheckApplies(cert) {
		return ResultStruct{Result: NA}, nil
	} else if !l.EffectiveDate.IsZero() && l.EffectiveDate.After(cert.NotBefore) {
		return ResultStruct{Result: NE}, nil
	}

	return l.Test.RunTest(cert)
}

func RegisterLint(l *Lint) {
	if Lints == nil {
		Lints = make(map[string]*Lint)
	}
	if err := l.Test.Initialize(); err != nil {
		panic("could not initialize lint: " + l.Name + ": " + err.Error())
	}
	Lints[l.Name] = l
}
