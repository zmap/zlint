// lint_subject_common_name_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-common-name INTEGER ::= 64
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectCommonNameMaxLength struct{}

func (l *subjectCommonNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectCommonNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectCommonNameMaxLength) Execute(c *x509.Certificate) *LintResult {
	if len(c.Subject.CommonName) > 64 {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_common_name_max_length",
		Description:   "The commonName field of the subject MUST be less than 64 characters",
		Source:        "RFC 5280: A.1",
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectCommonNameMaxLength{},
	})
}
