package lints

import (
	"crypto/dsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaTooShort struct {
	// Internal data here
}

func (l *dsaTooShort) Initialize() error {
	return nil
}

func (l *dsaTooShort) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaTooShort) RunTest(c *x509.Certificate) (ResultStruct, error) {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return ResultStruct{Result: NA}, nil
	}
	dsaParams := dsaKey.Parameters
	L := dsaParams.P.BitLen()
	N := dsaParams.Q.BitLen()
	if L >= 2048 && N >= 244 {
		return ResultStruct{Result: Pass}, nil
	}
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:        "e_dsa_shorter_than_2048_bits",
		Description: "DSA modulus size must be at least 2048 bits",
		Source:      "BRs: 6.1.5",
		// Refer to BRs: 6.1.5, taking the statement "Before 31 Dec 2010" literally
		EffectiveDate: util.ZeroDate,
		Test:          &dsaTooShort{},
	})
}
