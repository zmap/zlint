// lint_rsa_mod_less_than_2048_bits.go
/************************************************
Change this to match provenance TEXT
************************************************/

package lints

import (
	"crypto/rsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaParsedTestsKeySize struct {
	// Internal data here
}

func (l *rsaParsedTestsKeySize) Initialize() error {
	return nil
}

func (l *rsaParsedTestsKeySize) CheckApplies(c *x509.Certificate) bool {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA && c.NotAfter.After(util.NoRSA1024Date.Add(-1))
}

func (l *rsaParsedTestsKeySize) RunTest(c *x509.Certificate) (ResultStruct, error) {
	key := c.PublicKey.(*rsa.PublicKey)
	if key.N.BitLen() < 2048 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_rsa_mod_less_than_2048_bits",
		Description:   "For certificates valid after 31 Dec 2013, all certificates using RSA public key algorithm MUST have 2048 bits of modulus",
		Source:        "BRs: 6.1.5",
		EffectiveDate: util.ZeroDate,
		Test:          &rsaParsedTestsKeySize{},
	})
}
