// lint_subject_empty_without_san.go
/*************************************************************************
RFC 5280: 4.2 & 4.2.1.6
Further, if the only subject identity included in the certificate is
an alternative name form (e.g., an electronic mail address), then the
subject distinguished name MUST be empty (an empty sequence), and the
subjectAltName extension MUST be present.  If the subject field
contains an empty sequence, then the issuing CA MUST include a
subjectAltName extension that is marked as critical.  When including
the subjectAltName extension in a certificate that has a non-empty
subject distinguished name, conforming CAs SHOULD mark the
subjectAltName extension as non-critical.
*************************************************************************/
package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type emptyWithoutSAN struct {
	// Internal data here
}

func (l *emptyWithoutSAN) Initialize() error {
	return nil
}

func (l *emptyWithoutSAN) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *emptyWithoutSAN) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if subjectIsEmpty(cert) && !util.IsExtInCert(cert, util.SubjectAlternateNameOID) {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func subjectIsEmpty(cert *x509.Certificate) bool {
	if cert.Subject.Names == nil {
		return true
	}
	return false
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_empty_without_san",
		Description:   "CAs MUST support subject alternative name if the subject field is an empty sequence",
		Source:        "RFC 5280: 4.2 & 4.2.1.6",
		EffectiveDate: util.RFC2459Date,
		Test:          &emptyWithoutSAN{},
	})
}
