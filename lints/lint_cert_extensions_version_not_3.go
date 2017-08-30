// lint_cert_extensions_verson_not_3.go
/************************************************
4.1.2.1.  Version
   This field describes the version of the encoded certificate. When
   extensions are used, as expected in this profile, version MUST be 3
   (value is 2). If no extensions are present, but a UniqueIdentifier
   is present, the version SHOULD be 2 (value is 1); however, the version
   MAY be 3.  If only basic fields are present, the version SHOULD be 1
   (the value is omitted from the certificate as the default value);
   however, the version MAY be 2 or 3.

   Implementations SHOULD be prepared to accept any version certificate.
   At a minimum, conforming implementations MUST recognize version 3 certificates.
4.1.2.9.  Extensions
   This field MUST only appear if the version is 3 (Section 4.1.2.1).
   If present, this field is a SEQUENCE of one or more certificate
   extensions. The format and content of certificate extensions in the
   Internet PKI are defined in Section 4.2.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type CertExtensionsVersonNot3 struct {
	// Internal data here
}

func (l *CertExtensionsVersonNot3) Initialize() error {
	return nil
}

func (l *CertExtensionsVersonNot3) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *CertExtensionsVersonNot3) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if cert.Version != 3 && len(cert.Extensions) != 0 {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_cert_extensions_version_not_3",
		Description:   "The extensions field MUST only appear in version 3 certificates",
		Source:        "RFC 5280: 4.1.2.9",
		EffectiveDate: util.RFC2459Date,
		Test:          &CertExtensionsVersonNot3{},
	})
}
