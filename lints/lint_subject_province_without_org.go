// lint_subject_province_without_org.go
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

type provinceNoOrg struct{}

func (l *provinceNoOrg) Initialize() error {
	return nil
}

func (l *provinceNoOrg) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *provinceNoOrg) Execute(cert *x509.Certificate) LintResult {
	if util.TypeInName(&cert.Subject, util.StateOrProvinceNameOID) && !util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		return &LintResult{Status: Error}
	} else { //if no Province, Organization omitted
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_province_without_org",
		Description:   "The stateOrProvince name MUST NOT be included without an organization name",
		Source:        "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &provinceNoOrg{},
	})
}
