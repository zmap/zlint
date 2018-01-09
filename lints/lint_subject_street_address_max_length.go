/************************************************
ITU-T X.520 (02/2001) UpperBounds
ub-street-address INTEGER ::= 128

************************************************/

package lints

import (
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectStreetAddressMaxLength struct{}

func (l *subjectStreetAddressMaxLength) Initialize() error {
	return nil
}

func (l *subjectStreetAddressMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectStreetAddressMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.StreetAddress {
		if utf8.RuneCountInString(j) > 128 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_street_address_max_length",
		Description:   "The 'StreetAddress' field of the subject MUST be less than 129 characters",
		Citation:      "ITU-T X.520 (02/2001) UpperBounds",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectStreetAddressMaxLength{},
	})
}
