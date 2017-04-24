package lints

import (
	"time"

	"github.com/zmap/zcrypto/x509"
)

// global
var (
	Lints map[string]*Lint = make(map[string]*Lint)
)

type LintTest interface {
	// runs once globally
	Initialize() error
	CheckApplies(c *x509.Certificate) bool
	RunTest(c *x509.Certificate) (ResultStruct, error)
}

type lintReportUpdater func(*LintReport, *ResultStruct)

type Lint struct {
	Name          string
	Description   string
	Providence    string
	EffectiveDate time.Time
	Test          LintTest
	updateReport  lintReportUpdater
}

// LintReport holds the results of all lints ran on a certificate.
type LintReport struct {
	EBasicConstraintsNotCritical *ResultStruct
	EIanBareWildcard             *ResultStruct
	// etc.
}

// Execute runs all lints and records the results in this LintReport. If any
// lint returns an error, Execute will not run any remaining lints and returns
// the error.
func (report *LintReport) Execute(c *x509.Certificate) error {
	for _, lint := range Lints {
		res, err := lint.ExecuteTest(c)
		if err != nil {
			return err
		}
		lint.updateReport(report, &res)
	}
	return nil
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
