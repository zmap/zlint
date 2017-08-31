package lints

import (
	"crypto/dsa"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaUniqueCorrectRepresentation struct{}

func (l *dsaUniqueCorrectRepresentation) Initialize() error {
	return nil
}

func (l *dsaUniqueCorrectRepresentation) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaUniqueCorrectRepresentation) RunTest(c *x509.Certificate) (ResultStruct, error) {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return ResultStruct{Result: NA}, nil
	}
	// Verify that 2 ≤ y ≤ p-2.
	two := big.NewInt(2)
	pMinusTwo := big.NewInt(0)
	pMinusTwo.Sub(dsaKey.P, two)
	if two.Cmp(dsaKey.Y) > 0 || dsaKey.Y.Cmp(pMinusTwo) > 0 {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_unique_correct_representation",
		Description:   "DSA: Public key value has the unique correct representation in the field, and that the key has the correct order in the subgroup",
		Source:        "BRs: 6.1.6",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &dsaUniqueCorrectRepresentation{},
	})
}
