// lint_ext_subject_key_identifier_missing_ca.go
/************************************************
   To facilitate certification path construction, this extension MUST
   appear in all conforming CA certificates, that is, all certificates
   including the basic constraints extension (Section 4.2.1.9) where the
   value of cA is TRUE.  In conforming CA certificates, the value of the
   subject key identifier MUST be the value placed in the key identifier
   field of the authority key identifier extension (Section 4.2.1.1) of
   certificates issued by the subject of this certificate.  Applications
   are not required to verify that key identifiers match when performing
   certification path validation.
   ...
   For end entity certificates, the subject key identifier extension provides
   a means for identifying certificates containing the particular public key
   used in an application. Where an end entity has obtained multiple certificates,
   especially from multiple CAs, the subject key identifier provides a means to
   quickly identify the set of certificates containing a particular public key.
   To assist applications in identifying the appropriate end entity certificate,
   this extension SHOULD be included in all end entity certificates.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectKeyIdMissingCA struct {
	// Internal data here
}

func (l *subjectKeyIdMissingCA) Initialize() error {
	return nil
}

func (l *subjectKeyIdMissingCA) CheckApplies(cert *x509.Certificate) bool {
	return util.IsCACert(cert)
}

func (l *subjectKeyIdMissingCA) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(cert, util.SubjectKeyIdentityOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_subject_key_identifier_missing_ca",
		Description:   "CAs MUST include a Subject Key Identifier in all CA certificates",
		Source:        "RFC 5280: 4.2 & 4.2.1.2",
		EffectiveDate: util.RFC2459Date,
		Test:          &subjectKeyIdMissingCA{},
	})
}
