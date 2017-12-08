// lint_subject_surname_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-surname-length INTEGER ::= 40

************************************************/

package lints

import (
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectSurnameMaxLength struct{}

func (l *subjectSurnameMaxLength) Initialize() error {
	return nil
}

func (l *subjectSurnameMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectSurnameMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.Surname {
		if utf8.RuneCountInString(j) > 40 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_surname_max_length",
		Description:   "The 'Surname' field of the subject MUST be less than 41 characters",
		Citation:      "RFC 5280: A.1",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectSurnameMaxLength{},
	})
}
