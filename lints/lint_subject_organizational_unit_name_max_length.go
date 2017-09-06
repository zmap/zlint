// lint_subject_organizational_unit_name_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-organizational-unit-name INTEGER ::= 64
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"unicode/utf8"
)

type subjectOrganizationalUnitNameMaxLength struct{}

func (l *subjectOrganizationalUnitNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectOrganizationalUnitNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectOrganizationalUnitNameMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.OrganizationalUnit {
		if utf8.RuneCountInString(j) > 64 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_subject_organizational_unit_name_max_length",
		Description:    "The 'Organizational Unit Name' field of the subject MUST be less than 64 characters",
		ReadableSource: "RFC 5280: A.1",
		Source:         RFC5280,
		EffectiveDate:  util.RFC2459Date,
		Lint:           &subjectOrganizationalUnitNameMaxLength{},
	})
}
