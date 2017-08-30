package lints

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
)

type IssuerRDNHasMultipleAttribute struct {
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
		return ResultStruct{Result: Fatal}, err
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
		Name:          "w_multiple_issuer_rdn",
		Description:   "Certificates should not have multiple attributes in a single RDN (issuer)",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerRDNHasMultipleAttribute{},
	})
}
