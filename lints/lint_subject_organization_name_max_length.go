// lint_subject_organization_name_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-organization-name INTEGER ::= 64
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"unicode/utf8"
)

type subjectOrganizationNameMaxLength struct{}

func (l *subjectOrganizationNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectOrganizationNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectOrganizationNameMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.Organization {
		if utf8.RuneCountInString(j) > 64 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_subject_organization_name_max_length",
		Description:    "The 'Organization Name' field of the subject MUST be less than 64 characters",
		ReadableSource: "RFC 5280: A.1",
		Source:         RFC5280,
		EffectiveDate:  util.RFC2459Date,
		Lint:           &subjectOrganizationNameMaxLength{},
	})
}
