/************************************************
BRs: 7.1.2.1e
The Certificate Subject MUST contain the following: organizationName (OID 2.5.4.10): This field MUST be present and the contents MUST contain either the Subject CA’s name or DBA as verified under Section 3.2.2.2.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type caOrganizationNameMissing struct{}

func (l *caOrganizationNameMissing) Initialize() error {
	return nil
}

func (l *caOrganizationNameMissing) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

func (l *caOrganizationNameMissing) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.Organization != nil && c.Subject.Organization[0] != "" {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_organization_name_missing",
		Description:   "Root and Subordinate CA certificates MUST have a organizationName present in subject information",
		Citation:      "BRs: 7.1.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caOrganizationNameMissing{},
	})
}
