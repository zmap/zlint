// lint_issuer_dn_trailing_whitespace.go

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
	"strings"
)

type IssuerDNTrailingSpace struct {
	// Internal data here
}

func (l *IssuerDNTrailingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNTrailingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	var issuer pkix.RDNSequence
	if _, err := asn1.Unmarshal(c.RawIssuer, &issuer); err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	for _, rdn := range issuer {
		if len(rdn) == 0 {
			continue
		}
		atv := rdn[0]
		value, ok := atv.Value.(string)
		if !ok {
			continue
		}

		t := atv.Type
		if len(t) == 4 && t[0] == 2 && t[1] == 5 && t[2] == 4 && util.IsAttributeInList(t[3]) {
			if strings.HasSuffix(value, " ") {
				return ResultStruct{Result: Warn}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_trailing_whitespace",
		Description:   "AttributeValue in issuer RelativeDistinguishedName sequence should not have trailing whitespace",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerDNTrailingSpace{}})
}
