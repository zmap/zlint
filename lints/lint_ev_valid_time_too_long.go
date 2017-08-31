// lint_ev_valid_time_too_long.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evValidTooLong struct{}

func (l *evValidTooLong) Initialize() error {
	return nil
}

func (l *evValidTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
}

func (l *evValidTooLong) Execute(c *x509.Certificate) * LintResult{
	if c.NotBefore.AddDate(2, 3, 0).Before(c.NotAfter) {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_valid_time_too_long",
		Description:   "EV certificates must be 27 months in validity or less",
		Source:        "CAB 6.3.2",
		EffectiveDate: util.ZeroDate,
		Lint:          &evValidTooLong{},
	})
}
