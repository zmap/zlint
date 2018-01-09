// If the Certificate asserts the policy identifier of 2.23.140.1.2.1, then it MUST NOT include
// organizationName, streetAddress, localityName, stateOrProvinceName, or postalCode in the Subject field.

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type certPolicyConflictsWithStreet struct{}

func (l *certPolicyConflictsWithStreet) Initialize() error {
	return nil
}

func (l *certPolicyConflictsWithStreet) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRDomainValidatedOID) && !util.IsCACert(cert)
}

func (l *certPolicyConflictsWithStreet) Execute(cert *x509.Certificate) *LintResult {
	var out LintResult
	if util.TypeInName(&cert.Subject, util.StreetAddressOID) {
		out.Status = Error
	} else {
		out.Status = Pass
	}
	return &out
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cab_dv_conflicts_with_street",
		Description:   "If certificate policy 2.23.140.1.2.1 (CA/B BR domain validated) is included, streetAddress MUST NOT be included in subject",
		Citation:      "BRs: 7.1.6.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &certPolicyConflictsWithStreet{},
	})
}
