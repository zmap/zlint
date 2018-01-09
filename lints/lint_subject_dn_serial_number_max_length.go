package lints

import (
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNSerialNumberMaxLength struct{}

func (l *SubjectDNSerialNumberMaxLength) Initialize() error {
	return nil
}

func (l *SubjectDNSerialNumberMaxLength) CheckApplies(c *x509.Certificate) bool {
	return len(c.Subject.SerialNumber) > 0
}

func (l *SubjectDNSerialNumberMaxLength) Execute(c *x509.Certificate) *LintResult {
	if utf8.RuneCountInString(c.Subject.SerialNumber) > 64 {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_dn_serial_number_max_length",
		Description:   "The 'Serial Number' field of the subject MUST be less than 64 characters",
		Citation:      "RFC 5280: Appendix A",
		Source:        RFC5280,
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectDNSerialNumberMaxLength{},
	})
}
