// lint_ext_cert_policy_duplicate.go
/************************************************
  The certificate policies extension contains a sequence of one or more
  policy information terms, each of which consists of an object identifier
  (OID) and optional qualifiers. Optional qualifiers, which MAY be present,
  are not expected to change the definition of the policy. A certificate
  policy OID MUST NOT appear more than once in a certificate policies extension.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtCertPolicyDuplicate struct {
	// Internal data here
}

func (l *ExtCertPolicyDuplicate) Initialize() error {
	return nil
}

func (l *ExtCertPolicyDuplicate) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.CertPolicyOID)
}

func (l *ExtCertPolicyDuplicate) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	// O(n^2) is not terrible here because n is small
	for i := 0; i < len(cert.PolicyIdentifiers); i++ {
		for j := i + 1; j < len(cert.PolicyIdentifiers); j++ {
			if i != j && cert.PolicyIdentifiers[i].Equal(cert.PolicyIdentifiers[j]) {
				temp := ResultStruct{Result: Error}
				//temp.Details = fmt.Sprintf("%v", cert.PolicyIdentifiers[i])
				return temp, nil // Any one duplicate fails the test, so return here
			}
		}
	}
	// Nested loop will return if it finds a duplicate, so it's safe to assume pass
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_cert_policy_duplicate",
		Description:   "A certificate policy OID must not appear more than once in the extension",
		Source:        "RFC 5280: 4.2.1.4",
		EffectiveDate: util.RFC5280Date,
		Test:          &ExtCertPolicyDuplicate{},
	})
}
