// lint_subject_org_without_country.go
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

type orgNoCountry struct {
	// Internal data here
}

func (l *orgNoCountry) Initialize() error {
	return nil
}

func (l *orgNoCountry) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *orgNoCountry) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if !util.TypeInName(&cert.Subject, util.CountryNameOID) && util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		return ResultStruct{Result: Error}, nil
	} else { //if no organization, country can be nil
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_org_without_country",
		Description:   "The organization name field MUST not be included without a country name",
		Source:        "BRs: 7.1.4.2.2 (d&e)",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &orgNoCountry{},
	})
}
