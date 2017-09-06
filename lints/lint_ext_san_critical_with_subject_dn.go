// lint_ext_san_critical_with_subject_dn.go
/************************************************
Further, if the only subject identity included in the certificate is an
 alternative name form (e.g., an electronic mail address), then the subject
 distinguished name MUST be empty (an empty sequence), and the subjectAltName
 extension MUST be present. If the subject field contains an empty sequence,
 then the issuing CA MUST include a subjectAltName extension that is marked as
 critical. When including the subjectAltName extension in a certificate that
 has a non-empty subject distinguished name, conforming CAs SHOULD mark the
 subjectAltName extension as non-critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtSANCriticalWithSubjectDN struct{}

func (l *ExtSANCriticalWithSubjectDN) Initialize() error {
	return nil
}

func (l *ExtSANCriticalWithSubjectDN) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.SubjectAlternateNameOID)
}

func (l *ExtSANCriticalWithSubjectDN) Execute(cert *x509.Certificate) *LintResult {
	san := util.GetExtFromCert(cert, util.SubjectAlternateNameOID)
	if san.Critical && util.NotAllNameFieldsAreEmpty(&cert.Subject) {
		return &LintResult{Status: Warn}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "w_ext_san_critical_with_subject_dn",
		Description:    "If the subject contains a distinguished name, subjectAlternateName SHOULD be non-critical",
		ReadableSource: "RFC 5280: 4.2.1.6",
		Source:         CABFBaselineRequirements,
		EffectiveDate:  util.RFC5280Date,
		Lint:           &ExtSANCriticalWithSubjectDN{},
	})
}
