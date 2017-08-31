package lints

import (
	"crypto/dsa"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaSubgroup struct{}

func (l *dsaSubgroup) Initialize() error {
	return nil
}

func (l *dsaSubgroup) CheckApplies(c *x509.Certificate) bool {
	if c.PublicKeyAlgorithm != x509.DSA {
		return false
	}
	if _, ok := c.PublicKey.(*dsa.PublicKey); !ok {
		return false
	}
	return true
}

func (l *dsaSubgroup) RunTest(c *x509.Certificate) (ResultStruct, error) {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return ResultStruct{Result: NA}, nil
	}
	output := big.Int{}

	// Enforce that Y^Q == 1 mod P, e.g. that Order(Y) == Q mod P.
	output.Exp(dsaKey.Y, dsaKey.Q, dsaKey.P)
	if output.Cmp(big.NewInt(1)) == 0 {
		return ResultStruct{Result: Pass}, nil
	}
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_correct_order_in_subgroup",
		Description:   "DSA: Public key value has the unique correct representation in the field, and that the key has the correct order in the subgroup",
		Source:        "BRs: 6.1.6",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &dsaSubgroup{},
	})
}
