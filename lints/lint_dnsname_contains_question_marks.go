// lint_dnsname_contains_question_marks.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameContainsQuestionMarks struct{}

func (l *DNSNameContainsQuestionMarks) Initialize() error {
	return nil
}

func (l *DNSNameContainsQuestionMarks) CheckApplies(c *x509.Certificate) bool {
	return true
}

func isPrependedByQuestionMarks(domain string) bool {
	if strings.HasPrefix(domain, "?.") {
		return true
	}
	return false
}

func (l *DNSNameContainsQuestionMarks) Execute(c *x509.Certificate) *LintResult {
	if isPrependedByQuestionMarks(c.Subject.CommonName) {
		return &LintResult{Status: Notice}
	}
	for _, domain := range c.DNSNames {
		if isPrependedByQuestionMarks(domain) {
			return &LintResult{Status: Notice}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_dnsname_contains_question_marks",
		Description:   "Some Precerts are prepended with question marks.",
		Source:        "MDSP",
		EffectiveDate: util.ZeroDate,
		Lint:          &DNSNameContainsQuestionMarks{},
	})
}
