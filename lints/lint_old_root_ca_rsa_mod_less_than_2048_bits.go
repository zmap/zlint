// lint_old_root_ca_rsa_mod_less_than_2048_bits.go
/************************************************
Change this to match source TEXT
************************************************/

package lints

import (
	"crypto/rsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rootCaModSize struct{}

func (l *rootCaModSize) Initialize() error {
	return nil
}

func (l *rootCaModSize) CheckApplies(c *x509.Certificate) bool {
	issueDate := c.NotBefore
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA && util.IsRootCA(c) && issueDate.Before(util.NoRSA1024RootDate)
}

func (l *rootCaModSize) Execute(c *x509.Certificate) *LintResult {
	key := c.PublicKey.(*rsa.PublicKey)
	if key.N.BitLen() < 2048 {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_old_root_ca_rsa_mod_less_than_2048_bits",
		Description:   "In a validity period beginning on or before 31 Dec 2010, root CA certificates using RSA public key algorithm MUST use a 2048 bit modulus",
		Source:        "BRs: 6.1.5",
		EffectiveDate: util.ZeroDate,
		Lint:          &rootCaModSize{},
	})
}
