// lint_subject_dn_serial_number_not_printable_string.go
package lints

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNSerialNumberNotPrintableString struct{}

func (l *SubjectDNSerialNumberNotPrintableString) Initialize() error {
	return nil
}

func (l *SubjectDNSerialNumberNotPrintableString) CheckApplies(c *x509.Certificate) bool {
	return len(c.Subject.SerialNumber) > 0
}

func (l *SubjectDNSerialNumberNotPrintableString) Execute(c *x509.Certificate) *LintResult {
	rdnSequence := util.RawRDNSequence{}
	rest, err := asn1.Unmarshal(c.RawSubject, &rdnSequence)
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	if len(rest) > 0 {
		return &LintResult{Status: Fatal}
	}

	for _, attrTypeAndValueSet := range rdnSequence {
		for _, attrTypeAndValue := range attrTypeAndValueSet {
			if attrTypeAndValue.Type.Equal(util.SerialOID) && attrTypeAndValue.Value.Tag != asn1.TagPrintableString {
				return &LintResult{Status: Error}
			}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_dn_serial_number_not_printable_string",
		Description:   "X520 Distinguished Name SerialNumber MUST be encoded as PrintableString",
		Citation:      "RFC 5280: Appendix A",
		Source:        RFC5280,
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectDNSerialNumberNotPrintableString{},
	})
}
