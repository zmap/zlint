/**********************************************************************************************************************
BRs: 7.1.4.2.2
Other Subject Attributes
With the exception of the subject:organizationalUnitName (OU) attribute, optional attributes, when present within
the subject field, MUST contain information that has been verified by the CA. Metadata such as ‘.’, ‘-‘, and ‘ ‘ (i.e.
space) characters, and/or any other indication that the value is absent, incomplete, or not applicable, SHALL NOT
be used.
**********************************************************************************************************************/
package lints

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type illegalChar struct{}

func (l *illegalChar) Initialize() error {
	return nil
}

func (l *illegalChar) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *illegalChar) Execute(c *x509.Certificate) *LintResult {
	domain := c.Subject.DomainComponent
	serial := c.Subject.SerialNumber
	names := c.Subject.Names
	for _, j := range names {
		tempStr, ok := j.Value.(string)
		if !ok {
			continue //TODO: change this?
		}
		if tempStr == "-" || tempStr == "." || tempStr == " " {
			return &LintResult{Status: Error}
		}
	}
	if serial == "-" || serial == "." || serial == " " {
		return &LintResult{Status: Error}
	}
	for _, j := range domain {
		if strings.Compare(j, "-") == 0 ||
			strings.Compare(j, ".") == 0 ||
			strings.Compare(j, " ") == 0 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_contains_noninformational_value",
		Description:   "Subject name fields must not contain '.','-',' ' or any other indication that the field has been omitted",
		Citation:      "BRs: 7.1.4.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &illegalChar{},
	})
}
