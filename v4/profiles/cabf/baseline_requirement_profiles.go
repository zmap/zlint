package cabf

import (
	"github.com/zmap/zcrypto/x509"

	"github.com/zmap/zlint/v4/lint"
	"github.com/zmap/zlint/v4/util"
)

// 7.1.2.1 Root CA Certificate Profile
// Subordinate CAs
//   Cross Certificates
//     7.1.2.2 Cross-Certified Subordinate CA Certificate Profile
//   Technically Constrainted CA Certificates
//     7.1.2.3 - Technically-Constrained Non-TLS Subordinate CA Certificate Profile
//     7.1.2.4 - Technically-Constrained Precertificate Signing CA Certificate Profile
//     7.1.2.5 - Technically-Constrained TLS Subordinate CA Certificate Profile
//   7.1.2.6 - TLS Subordinate CA Certificate Profile

// 7.1.2.7 - Subscriber (End-Entity) Certificate Profile
// 7.1.2.8 - OCSP Responder Certificate Profile
// 7.1.2.9 - Precertificate Profile

const (
	RootCA                                        = 1
	CrossCertificateSubordinateCA                 = 2
	TechnicallyConstrainedNonTLSSubordinateCA     = 3
	TechnicallyConstrainedPrecertificateSigningCA = 4
	TechnicallyConstraintTLSSubordinateCA         = 5
	TLSSubordinateCA                              = 6
	Subscriber                                    = 7
	OCSP                                          = 8
	Precertificate                                = 9
)

type rootCAMatcher struct{}

func (m *rootCAMatcher) CheckMatches(c *x509.Certificate) bool {
	return util.IsCA(c) && util.HasTBSCertificate(c)
}

func (m *rootCAMatcher) CheckFields()

var RootCACertificateProfile = lint.Profile{
	Name:    "Root CA Certificate",
	Handle:  RootCA,
	Matcher: &rootCAMatcher{},
}
