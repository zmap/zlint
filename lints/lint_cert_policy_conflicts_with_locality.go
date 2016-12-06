// lint_cert_policy_conflicts_with_locality.go
// If the Certificate asserts the policy identifier of 2.23.140.1.2.1, then it MUST NOT include
// organizationName, streetAddress, localityName, stateOrProvinceName, or postalCode in the Subject field.

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type certPolicyConflictsWithLocality struct {
	// Internal data here
}

func (l *certPolicyConflictsWithLocality) Initialize() error {
	return nil
}

func (l *certPolicyConflictsWithLocality) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRDomainValidatedOID)
}

func (l *certPolicyConflictsWithLocality) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	var out ResultStruct
	if util.TypeInName(&cert.Subject, util.LocalityNameOID) {
		out.Result = Error
	} else {
		out.Result = Pass
	}
	return out, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "cert_policy_conflicts_with_locality",
		Description:   "If certificate policy 2.23.140.1.2.1 (CA/B BR domain validated) is included, locality name must not be included in subject.",
		Providence:    "CAB: 7.1.6.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &certPolicyConflictsWithLocality{}})
}
