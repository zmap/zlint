// lint_ext_ian_uri_not_ia5.go
/************************************************
When the issuerAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).
************************************************/

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
	"unicode"
)

type ianUriIa5 struct {
	// Internal data here
}

func (l *ianUriIa5) Initialize() error {
	return nil
}

func (l *ianUriIa5) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *ianUriIa5) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, uri := range c.IANURIs {
		for _, c := range uri {
			if c > unicode.MaxASCII {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_ian_uri_not_ia5",
		Description:   "When SAN contains a URI, the name must be an IA5 string",
		Providence:    "RFC5280: 4.2.1.7",
		EffectiveDate: util.RFC5280Date,
		Test:          &ianUriIa5{}})
}
