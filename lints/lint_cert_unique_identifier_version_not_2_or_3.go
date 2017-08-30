// lint_cert_unique_identifier_version_not_2_or_3.go
/**************************************************************************
RFC 5280: 4.1.2.8
 These fields MUST only appear if the version is 2 or 3 (Section 4.1.2.1).
 These fields MUST NOT appear if the version is 1. The subject and issuer
 unique identifiers are present in the certificate to handle the possibility
 of reuse of subject and/or issuer names over time. This profile RECOMMENDS
 that names not be reused for different entities and that Internet certificates
 not make use of unique identifiers. CAs conforming to this profile MUST NOT
 generate certificates with unique identifiers. Applications conforming to
 this profile SHOULD be capable of parsing certificates that include unique
 identifiers, but there are no processing requirements associated with the
 unique identifiers.
****************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type certUniqueIdVersion struct {
	// Internal data here
}

func (l *certUniqueIdVersion) Initialize() error {
	return nil
}

func (l *certUniqueIdVersion) CheckApplies(c *x509.Certificate) bool {
	return c.IssuerUniqueId.Bytes != nil || c.SubjectUniqueId.Bytes != nil
}

func (l *certUniqueIdVersion) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if (c.Version) != 2 && (c.Version) != 3 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cert_unique_identifier_version_not_2_or_3",
		Description:   "Unique identifiers MUST only appear if the X.509 version is 2 or 3",
		Source:        "RFC 5280: 4.1.2.8",
		EffectiveDate: util.RFC5280Date,
		Test:          &certUniqueIdVersion{},
	})
}
