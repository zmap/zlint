package lints

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubjectDNCountryNotPrintableString struct{}

func (l *SubjectDNCountryNotPrintableString) Initialize() error {
	return nil
}

func (l *SubjectDNCountryNotPrintableString) CheckApplies(c *x509.Certificate) bool {
	return len(c.Subject.Country) > 0
}

func (l *SubjectDNCountryNotPrintableString) Execute(c *x509.Certificate) *LintResult {
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
			if attrTypeAndValue.Type.Equal(util.CountryNameOID) && attrTypeAndValue.Value.Tag != asn1.TagPrintableString {
				return &LintResult{Status: Error}
			}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_dn_country_not_printable_string",
		Description:   "X520 Distinguished Name Country MUST be encoded as PrintableString",
		Citation:      "RFC 5280: Appendix A",
		Source:        RFC5280,
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectDNCountryNotPrintableString{},
	})
}
