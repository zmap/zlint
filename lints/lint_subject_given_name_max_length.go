/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-given-name-length INTEGER ::= 16

************************************************/

package lints

import (
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectGivenNameMaxLength struct{}

func (l *subjectGivenNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectGivenNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectGivenNameMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.GivenName {
		if utf8.RuneCountInString(j) > 16 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_given_name_max_length",
		Description:   "The 'GivenName' field of the subject MUST be less than 17 characters",
		Citation:      "RFC 5280: A.1",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectGivenNameMaxLength{},
	})
}
