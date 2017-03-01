// lint_ext_san_uri_not_ia5.go
/************************************************
When the subjectAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).
************************************************/

package lints

import (
	"github.com/zmap/zgrab/ztools/x509"
	"github.com/zmap/zlint/util"
	"unicode"
)

type extSanURINotIA5 struct {
	// Internal data here
}

func (l *extSanURINotIA5) Initialize() error {
	return nil
}

func (l *extSanURINotIA5) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SanOID)
}

func (l *extSanURINotIA5) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, uri := range c.URIs {
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
		Name:          "ext_san_uri_not_ia5",
		Description:   "When SAN contains a URI, the name must be an IA5 string",
		Providence:    "RFC5280: 4.2.1.6",
		EffectiveDate: util.RFC5280Date,
		Test:          &extSanURINotIA5{}})
}
