// lint_subject_common_name_not_lowercase.go
/************************************************
This lint will emit a notice if the common name contains capital letters.
This case is not explicitly covered by the BR, but the common name would not be
byte by byte equal to any of the SAN DNS names.
************************************************/

package lints

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectCommonNameNotLowercase struct{}

func (l *subjectCommonNameNotLowercase) Initialize() error {
	return nil
}

func (l *subjectCommonNameNotLowercase) CheckApplies(c *x509.Certificate) bool {
	return c.Subject.CommonName != "" && !util.IsCACert(c)
}

func (l *subjectCommonNameNotLowercase) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != strings.ToLower(c.Subject.CommonName) {
		return &LintResult{Status: Notice}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_lint_subject_common_name_not_lowercase",
		Description:   "The common name field in not in lower case",
		Citation:      "This case is not explicitly addressed in the CA/B BR",
		Source:        ZLint,
		EffectiveDate: util.ZeroDate,
		Lint:          &subjectCommonNameNotLowercase{},
	})
}
