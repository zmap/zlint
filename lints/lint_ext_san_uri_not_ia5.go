// lint_ext_san_uri_not_ia5.go
/************************************************
When the subjectAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"unicode"
)

type extSANURINotIA5 struct{}

func (l *extSANURINotIA5) Initialize() error {
	return nil
}

func (l *extSANURINotIA5) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *extSANURINotIA5) Execute(c *x509.Certificate) *LintResult {
	for _, uri := range c.URIs {
		for _, c := range uri {
			if c > unicode.MaxASCII {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_uri_not_ia5",
		Description:   "When subjectAlternateName contains a URI, the name MUST be an IA5 string",
		Source:        "RFC5280: 4.2.1.6",
		EffectiveDate: util.RFC5280Date,
		Lint:          &extSANURINotIA5{},
	})
}
