// lint_cert_policy_conflicts_with_org.go
// If the Certificate asserts the policy identifier of 2.23.140.1.2.1, then it MUST NOT include
// organizationName, streetAddress, localityName, stateOrProvinceName, or postalCode in the Subject field.

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type certPolicyConflictsWithOrg struct {
	// Internal data here
}

func (l *certPolicyConflictsWithOrg) Initialize() error {
	return nil
}

func (l *certPolicyConflictsWithOrg) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRDomainValidatedOID) && !util.IsCACert(cert)
}

func (l *certPolicyConflictsWithOrg) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	var out ResultStruct
	if util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		out.Result = Error
	} else {
		out.Result = Pass
	}
	return out, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cab_dv_conflicts_with_org",
		Description:   "If certificate policy 2.23.140.1.2.1 (CA/B BR domain validated) is included, organization name MUST NOT be included in subject",
		Source:        "BRs: 7.1.6.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &certPolicyConflictsWithOrg{},
	})
}
