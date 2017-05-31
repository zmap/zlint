// lint_issuer_multiple_attr_in_rdn.go

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
)

type IssuerRDNHasMultipleAttribute struct {
	// Internal data here
}

func (l *IssuerRDNHasMultipleAttribute) Initialize() error {
	return nil
}

func (l *IssuerRDNHasMultipleAttribute) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerRDNHasMultipleAttribute) RunTest(c *x509.Certificate) (ResultStruct, error) {
	var issuer pkix.RDNSequence
	_, err := asn1.Unmarshal(c.RawIssuer, &issuer)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	for _, rdn := range issuer {
		if len(rdn) > 1 {
			return ResultStruct{Result: Warn}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_multiple_attr_in_rdn",
		Description:   "Certificates should not have multiple attributes in a single RDN in the subject name",
		Providence:    "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerRDNHasMultipleAttribute{}})
}
