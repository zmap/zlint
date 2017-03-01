package lints

import (
	"github.com/zmap/zgrab/ztools/x509"
	"math/big"
	"time"
)

// global
var (
	Lints      map[string]*Lint = make(map[string]*Lint)
	upperBound                  = &big.Int{}
)

type LintTest interface {
	// runs once globally
	Initialize() error
	CheckApplies(c *x509.Certificate) bool
	RunTest(c *x509.Certificate) (ResultStruct, error)
}

type Lint struct {
	Name          string
	Description   string
	Providence    string
	EffectiveDate time.Time
	Test          LintTest
}

func Init() {
	upperBound.Exp(big.NewInt(2), big.NewInt(256), nil)
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
	l.Test.Initialize()
	Lints[l.Name] = l
}
