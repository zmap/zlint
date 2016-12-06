// lint_subject_street_without_org.go
/*********************************************************************************************************
Required/Optional: The organization name is OPTIONAL. If organization name is present, then localityName,
stateOrProvinceName (where applicable), and countryName are REQUIRED and streetAddress and postalCode are
OPTIONAL. If organization name is absent, then the Certificate MUST NOT contain a streetAddress,
localityName, stateOrProvinceName, or postalCode attribute. The CA MAY include the Subject’s countryName
field without including other Subject Identity Information pursuant to Section 9.2.5.
**********************************************************************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type streetNoOrg struct {
	// Internal data here
}

func (l *streetNoOrg) Initialize() error {
	return nil
}

func (l *streetNoOrg) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *streetNoOrg) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if util.TypeInName(&cert.Subject, util.StreetAddressOID) && !util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		return ResultStruct{Result: Error}, nil
	} else { //if no Street address, Organization can be ommited
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "subject_street_without_org",
		Description:   "The street address field must not be included without an organization name.",
		Providence:    "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &streetNoOrg{}})
}
