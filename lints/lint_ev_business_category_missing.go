// lint_ev_business_category_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evNoBiz struct{}

func (l *evNoBiz) Initialize() error {
	return nil
}

func (l *evNoBiz) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
}

func (l *evNoBiz) Execute(c *x509.Certificate) *LintResult {
	if util.TypeInName(&c.Subject, util.BusinessOID) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_business_category_missing",
		Description:   "EV certificates must include businessCategory in subject",
		Source:        "BRs: 7.1.6.1",
		Type:          BRs,
		EffectiveDate: util.ZeroDate,
		Lint:          &evNoBiz{},
	})
}
