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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type streetNoOrg struct{}

func (l *streetNoOrg) Initialize() error {
	return nil
}

func (l *streetNoOrg) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *streetNoOrg) Execute(cert *x509.Certificate) * LintResult{
	if util.TypeInName(&cert.Subject, util.StreetAddressOID) && !util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		return &LintResult{Status: Error}
	} else { //if no Street address, Organization can be omitted
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_street_without_org",
		Description:   "The 'Street Address' field MUST NOT be included without an organization name",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &streetNoOrg{},
	})
}
