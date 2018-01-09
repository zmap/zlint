/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-emailaddress-length INTEGER ::= 128
************************************************/

package lints

import (
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectEmailMaxLength struct{}

func (l *subjectEmailMaxLength) Initialize() error {
	return nil
}

func (l *subjectEmailMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectEmailMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.EmailAddress {
		if utf8.RuneCountInString(j) > 128 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_email_max_length",
		Description:   "The 'Email' field of the subject MUST be less than 129 characters",
		Citation:      "RFC 5280: A.1",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectEmailMaxLength{},
	})
}
