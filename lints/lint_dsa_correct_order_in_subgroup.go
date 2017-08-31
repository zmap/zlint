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

func (l *dsaSubgroup) Execute(c *x509.Certificate) *LintResult {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return &LintResult{Status: NA}
	}
	output := big.Int{}

	// Enforce that Y^Q == 1 mod P, e.g. that Order(Y) == Q mod P.
	output.Exp(dsaKey.Y, dsaKey.Q, dsaKey.P)
	if output.Cmp(big.NewInt(1)) == 0 {
		return &LintResult{Status: Pass}
	}
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_correct_order_in_subgroup",
		Description:   "DSA: Public key value has the unique correct representation in the field, and that the key has the correct order in the subgroup",
		Source:        "BRs: 6.1.6",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &dsaSubgroup{},
	})
}
