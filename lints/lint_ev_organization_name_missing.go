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

func (l *evOrgMissing) Execute(c *x509.Certificate) ResultStruct {
	if util.TypeInName(&c.Subject, util.OrganizationNameOID) {
		return ResultStruct{Result: Pass}
	} else {
		return ResultStruct{Result: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_organization_name_missing",
		Description:   "EV certificates must include organizationName in subject",
		Source:        "CAB 7.1.6.1",
		EffectiveDate: util.ZeroDate,
		Lint:          &evOrgMissing{},
	})
}
