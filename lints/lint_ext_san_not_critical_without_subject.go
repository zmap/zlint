// lint_ext_san_not_critical_without_subject.go
/************************************************
RFC 5280: 4.2.1.6
Further, if the only subject identity included in the certificate is
   an alternative name form (e.g., an electronic mail address), then the
   subject distinguished name MUST be empty (an empty sequence), and the
   subjectAltName extension MUST be present.  If the subject field
   contains an empty sequence, then the issuing CA MUST include a
   subjectAltName extension that is marked as critical.  When including
   the subjectAltName extension in a certificate that has a non-empty
   subject distinguished name, conforming CAs SHOULD mark the
   subjectAltName extension as non-critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type extSANNotCritNoSubject struct {
	// Internal data here
}

func (l *extSANNotCritNoSubject) Initialize() error {
	return nil
}

func (l *extSANNotCritNoSubject) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SANOID)
}

func (l *extSANNotCritNoSubject) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if e := util.GetExtFromCert(c, util.SANOID); !util.NotAllNameFieldsAreEmpty(&c.Subject) && !e.Critical {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_not_critical_without_subject",
		Description:   "If there is an empty subject field, then the SAN extension MUST be critical",
		Providence:    "RFC 5280: 4.2.1.6",
		EffectiveDate: util.RFC2459Date,
		Test:          &extSANNotCritNoSubject{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.EExtSanNotCriticalWithoutSubject = result },
	})
}
