/*If the Certificate asserts the policy identifier of 2.23.140.1.2.3, then it MUST also include (i) either organizationName or givenName and surname, (ii) localityName (to the extent such field is required under Section 7.1.4.2.2), (iii) stateOrProvinceName (to the extent required under Section 7.1.4.2.2), and (iv) countryName in the Subject field.*/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type CertPolicyRequiresPersonalName struct{}

func (l *CertPolicyRequiresPersonalName) Initialize() error {
	return nil
}

func (l *CertPolicyRequiresPersonalName) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRIndividualValidatedOID) && !util.IsCACert(cert)
}

func (l *CertPolicyRequiresPersonalName) Execute(cert *x509.Certificate) *LintResult {
	var out LintResult
	if util.TypeInName(&cert.Subject, util.OrganizationNameOID) || (util.TypeInName(&cert.Subject, util.GivenNameOID) && util.TypeInName(&cert.Subject, util.SurnameOID)) {
		out.Status = Pass
	} else {
		out.Status = Error
	}
	return &out
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cab_iv_requires_personal_name",
		Description:   "If certificate policy 2.23.140.1.2.3 is included, either organizationName or givenName and surname MUST be included in subject",
		Citation:      "BRs: 7.1.6.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABV131Date,
		Lint:          &CertPolicyRequiresPersonalName{},
	})
}
