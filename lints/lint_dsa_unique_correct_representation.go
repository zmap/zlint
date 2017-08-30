package lints

import (
	"crypto/dsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"math/big"
)

type dsaUniqueCorrectRepresentation struct {
	// Internal data here
}

func (l *dsaUniqueCorrectRepresentation) Initialize() error {
	return nil
}

func (l *dsaUniqueCorrectRepresentation) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaUniqueCorrectRepresentation) RunTest(c *x509.Certificate) (ResultStruct, error) {
	dsaKey := c.PublicKey.(*dsa.PublicKey)
	twoConst := big.NewInt(2)
	if dsaKey.Y.Cmp(twoConst) == 0 || dsaKey.Y.Cmp(twoConst) == 1 {
		if dsaKey.Y.Cmp(dsaKey.P) == 0 || dsaKey.Y.Cmp(dsaKey.P) == -1 {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_unique_correct_representation",
		Description:   "DSA: Public key value has the unique correct representation in the field, and that the key has the correct order in the subgroup",
		Provenance:    "BRs: 6.1.6",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &dsaUniqueCorrectRepresentation{},
	})
}
