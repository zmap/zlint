package lints

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
)

type SubjectRDNHasMultipleAttribute struct{}

func (l *SubjectRDNHasMultipleAttribute) Initialize() error {
	return nil
}

func (l *SubjectRDNHasMultipleAttribute) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectRDNHasMultipleAttribute) Execute(c *x509.Certificate) *LintResult {
	var subject pkix.RDNSequence
	if _, err := asn1.Unmarshal(c.RawSubject, &subject); err != nil {
		return &LintResult{Status: Fatal}
	}
	for _, rdn := range subject {
		if len(rdn) > 1 {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_multiple_subject_rdn",
		Description:   "Certificates should not have multiple attributes in a single RDN (subject)",
		Citation:      "AWSLabs certlint",
		Source:        AWSLabs,
		EffectiveDate: util.ZeroDate,
		Lint:          &SubjectRDNHasMultipleAttribute{},
	})
}
