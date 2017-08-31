// lint_ext_ian_uri_not_ia5.go
/************************************************
When the issuerAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"unicode"
)

type IANURIIA5String struct{}

func (l *IANURIIA5String) Initialize() error {
	return nil
}

func (l *IANURIIA5String) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANURIIA5String) Execute(c *x509.Certificate) *LintResult {
	for _, uri := range c.IANURIs {
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
		Name:          "e_ext_ian_uri_not_ia5",
		Description:   "When subjectAltName contains a URI, the name MUST be an IA5 string",
		Source:        "RFC5280: 4.2.1.7",
		EffectiveDate: util.RFC5280Date,
		Lint:          &IANURIIA5String{},
	})
}
