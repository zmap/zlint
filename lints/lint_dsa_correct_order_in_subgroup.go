// lint_dsa_improper_modulus_or_divisor_size.go
/************************************************
BRs: 6.1.5
Certificates MUST meet the following requirements for algorithm type and key size.
Minimum DSA modulus and divisor size (bits)***: L=2048,	N=224 or L=2048, N=256.
**As a note, this points to FIPS 186-4 for further clarification**
************************************************/

package lints

import (
	"crypto/dsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"math/big"
)

type dsaSubgroup struct {
	// Internal data here
}

func (l *dsaSubgroup) Initialize() error {
	return nil
}

func (l *dsaSubgroup) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaSubgroup) RunTest(c *x509.Certificate) (ResultStruct, error) {
	dsaKey := c.PublicKey.(*dsa.PublicKey)
	output := big.Int{}
	output.Exp(dsaKey.Y, dsaKey.Q, dsaKey.P)
	if output.Cmp(big.NewInt(1)) == 0 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_improper_modulus_or_divisor_size",
		Description:   "DSA: Public key value has the unique correct representation in the field, and that the key has the correct order in the subgroup",
		Provenance:    "BRs: 6.1.6",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &dsaSubgroup{},
	})
}
