// lint_old_root_ca_rsa_mod_less_than_2048_bits.go
/************************************************
Change this to match providence TEXT
************************************************/

package lints

import (

	"crypto/rsa"
	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type rootCaModSize struct {
	// Internal data here
}

func (l *rootCaModSize) Initialize() error {
	return nil
}

func (l *rootCaModSize) CheckApplies(c *x509.Certificate) bool {
	issueDate := c.NotBefore
	return (c.PublicKeyAlgorithm == x509.RSA && util.IsRootCA(c) && issueDate.Before(util.RsaDate2))
}

func (l *rootCaModSize) RunTest(c *x509.Certificate) (ResultStruct, error) {
	mod := c.PublicKey.(*rsa.PublicKey).N
	if mod.BitLen() < 2048 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "old_root_ca_rsa_mod_less_than_2048_bits",
		Description:   "In a validity period beginning on or before 31 dec 2010, root CA certificates using RSA public key algorithm must have 2048 bits of modulus",
		Providence:    "CAB: 6.1.5",
		EffectiveDate: util.ZeroDate,
		Test:          &rootCaModSize{}})
}
