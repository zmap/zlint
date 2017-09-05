// lint_ev_locality_name_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evLocalityMissing struct{}

func (l *evLocalityMissing) Initialize() error {
	return nil
}

func (l *evLocalityMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
}

func (l *evLocalityMissing) Execute(c *x509.Certificate) *LintResult {
	if util.TypeInName(&c.Subject, util.LocalityNameOID) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_locality_name_missing",
		Description:   "EV certificates must include localityName in subject",
		Source:        "BRs: 7.1.6.1",
		Type:          CABFBaselineRequirements,
		EffectiveDate: util.ZeroDate,
		Lint:          &evLocalityMissing{},
	})
}
