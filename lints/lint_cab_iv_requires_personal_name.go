// lint_cert_policy_requires_personal_name.go
/*If the Certificate asserts the policy identifier of 2.23.140.1.2.3, then it MUST also include (i) either organizationName or givenName and surname, (ii) localityName (to the extent such field is required under Section 7.1.4.2.2), (iii) stateOrProvinceName (to the extent required under Section 7.1.4.2.2), and (iv) countryName in the Subject field.*/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type CertPolicyRequiresPersonalName struct {
	// Internal data here
}

func (l *CertPolicyRequiresPersonalName) Initialize() error {
	return nil
}

func (l *CertPolicyRequiresPersonalName) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRIndividualValidatedOID) && !util.IsCACert(cert)
}

func (l *CertPolicyRequiresPersonalName) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	var out ResultStruct
	if util.TypeInName(&cert.Subject, util.OrganizationNameOID) || (util.TypeInName(&cert.Subject, util.GivenNameOID) && util.TypeInName(&cert.Subject, util.SurnameOID)) {
		out.Result = Pass
	} else {
		out.Result = Error
	}
	return out, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cab_iv_requires_personal_name",
		Description:   "If certificate policy 2.23.140.1.2.3 is included, either organizationName or givenName and surname MUST be included in subject",
		Source:        "BRs: 7.1.6.1",
		EffectiveDate: util.CABV131Date,
		Test:          &CertPolicyRequiresPersonalName{},
	})
}
