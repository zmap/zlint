// lint_cert_policy_conflicts_with_locality.go
// If the Certificate asserts the policy identifier of 2.23.140.1.2.1, then it MUST NOT include
// organizationName, streetAddress, localityName, stateOrProvinceName, or postalCode in the Subject field.

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type certPolicyConflictsWithLocality struct{}

func (l *certPolicyConflictsWithLocality) Initialize() error {
	return nil
}

func (l *certPolicyConflictsWithLocality) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRDomainValidatedOID) && !util.IsCACert(cert)
}

func (l *certPolicyConflictsWithLocality) Execute(cert *x509.Certificate) ResultStruct {
	var out ResultStruct
	if util.TypeInName(&cert.Subject, util.LocalityNameOID) {
		out.Result = Error
	} else {
		out.Result = Pass
	}
	return out
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cab_dv_conflicts_with_locality",
		Description:   "If certificate policy 2.23.140.1.2.1 (CA/B BR domain validated) is included, locality name MUST NOT be included in subject",
		Source:        "BRs: 7.1.6.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &certPolicyConflictsWithLocality{},
	})
}
