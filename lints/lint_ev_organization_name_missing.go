// lint_ev_organization_name_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evOrgMissing struct{}

func (l *evOrgMissing) Initialize() error {
	return nil
}

func (l *evOrgMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
}

func (l *evOrgMissing) Execute(c *x509.Certificate) *LintResult {
	if util.TypeInName(&c.Subject, util.OrganizationNameOID) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_organization_name_missing",
		Description:   "EV certificates must include organizationName in subject",
		Source:        "BRs: 7.1.6.1",
		Type:          CABFBaselineRequirements,
		EffectiveDate: util.ZeroDate,
		Lint:          &evOrgMissing{},
	})
}
