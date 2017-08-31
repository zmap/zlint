// lint_subject_state_name_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-state-name INTEGER ::= 128
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectStateNameMaxLength struct{}

func (l *subjectStateNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectStateNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectStateNameMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.Province {
		if len(j) > 128 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_state_name_max_length",
		Description:   "The 'State Name' field of the subject MUST be less than 128 characters",
		Source:        "RFC 5280: A.1",
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectStateNameMaxLength{},
	})
}
