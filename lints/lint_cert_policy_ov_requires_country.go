// lint_cert_policy_ov_requires_country.go
/*If the Certificate asserts the policy identifier of 2.23.140.1.2.2, then it MUST also include organizationName, localityName (to the extent such field is required under Section 7.1.4.2.2), stateOrProvinceName (to the extent such field is required under Section 7.1.4.2.2), and countryName in the Subject field.*/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type CertPolicyOVRequiresCountry struct{}

func (l *CertPolicyOVRequiresCountry) Initialize() error {
	return nil
}

func (l *CertPolicyOVRequiresCountry) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BROrganizationValidatedOID)
}

func (l *CertPolicyOVRequiresCountry) Execute(cert *x509.Certificate) *LintResult {
	var out LintResult
	if util.TypeInName(&cert.Subject, util.CountryNameOID) {
		out.Status = Pass
	} else {
		out.Status = Error
	}
	return &out
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cert_policy_ov_requires_country",
		Description:   "If certificate policy 2.23.140.1.2.2 is included, countryName MUST be included in subject",
		Citation:      "BRs: 7.1.6.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &CertPolicyOVRequiresCountry{},
	})
}
