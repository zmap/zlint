package lints

import (
	"crypto/dsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaImproperSize struct {
	// Internal data here
}

func (l *dsaImproperSize) Initialize() error {
	return nil
}

func (l *dsaImproperSize) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaImproperSize) RunTest(c *x509.Certificate) (ResultStruct, error) {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return ResultStruct{Result: NA}, nil
	}
	L := dsaKey.Parameters.P.BitLen()
	N := dsaKey.Parameters.Q.BitLen()
	if (L == 2048 && N == 224) || (L == 2048 && N == 256) || (L == 3072 && N == 256) {
		return ResultStruct{Result: Pass}, nil
	}
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_improper_modulus_or_divisor_size",
		Description:   "Certificates MUST meet the following requirements for algorithm type and key size: L=2048, N=224,256 minimum DSA",
		Source:        "BRs: 6.1.5",
		EffectiveDate: util.ZeroDate,
		Test:          &dsaImproperSize{},
	})
}
