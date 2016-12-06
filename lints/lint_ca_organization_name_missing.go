// lint_ca_organization_name_missing.go
/************************************************
CAB: 7.1.2.1e
The Certificate Subject MUST contain the following: organizationName (OID 2.5.4.10): This field MUST be present and the contents MUST contain either the Subject CA’s name or DBA as verified under Section 3.2.2.2.
************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type caOrganizationNameMissing struct {
	// Internal data here
}

func (l *caOrganizationNameMissing) Initialize() error {
	return nil
}

func (l *caOrganizationNameMissing) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.IsCA
}

func (l *caOrganizationNameMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.Organization != nil && c.Subject.Organization[0] != "" {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "ca_organization_name_missing",
		Description:   "Root & Subordinate CA certificates must have a organizationName present in subject information",
		Providence:    "CAB: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &caOrganizationNameMissing{}})
}
