// lint_subject_postal_without_org.go
/*********************************************************************************************************
Required/Optional: The organization name is OPTIONAL. If organization name is present, then localityName,
stateOrProvinceName (where applicable), and countryName are REQUIRED and streetAddress and postalCode are
OPTIONAL. If organization name is absent, then the Certificate MUST NOT contain a streetAddress,
localityName, stateOrProvinceName, or postalCode attribute. The CA MAY include the Subject’s countryName
field without including other Subject Identity Information pursuant to Section 9.2.5.
**********************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type postalNoOrg struct {
	// Internal data here
}

func (l *postalNoOrg) Initialize() error {
	return nil
}

func (l *postalNoOrg) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *postalNoOrg) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if util.TypeInName(&cert.Subject, util.PostalCodeOID) && !util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		return ResultStruct{Result: Error}, nil
	} else { //if no Postal code, Organization can be omitted
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_postal_without_org",
		Description:   "The postal code MUST NOT be included without an organization name",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &postalNoOrg{},
	})
}
