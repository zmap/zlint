// lint_subject_org_without_locality_or_province.go
/*******************************************************************************************************
Required/Optional: The organization name is OPTIONAL. If organization name is present, then localityName, stateOrProvinceName (where applicable), and countryName are REQUIRED and streetAddress and postalCode are OPTIONAL. If organization name is absent, then the Certificate MUST NOT contain a streetAddress, localityName, stateOrProvinceName, or postalCode attribute. The CA MAY include the Subject’s countryName field without including other Subject Identity Information pursuant to Section 9.2.5.
*******************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type orgNoLocalOrProvince struct {
	// Internal data here
}

func (l *orgNoLocalOrProvince) Initialize() error {
	return nil
}

func (l *orgNoLocalOrProvince) CheckApplies(cert *x509.Certificate) bool {
	return !util.IsCACert(cert)
}

func (l *orgNoLocalOrProvince) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if !util.TypeInName(&cert.Subject, util.LocalityNameOID) && !util.TypeInName(&cert.Subject, util.StateOrProvinceNameOID) && util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		return ResultStruct{Result: Error}, nil
	} else { //if no organization, local/province can be nil, only one of the two is required if org is preasent
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_org_without_locality_or_province",
		Description:   "If organization is included in a subscriber certificate, either stateOrProvince or locality MUST be included",
		Source:        "BRs: 7.1.4.2.2 (d&e)",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &orgNoLocalOrProvince{},
	})
}
