// lint_rsa_exp_negative.go

package lints

import (

	"crypto/rsa"
	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type rsaExpNegative struct {
	// Internal data here
}

func (l *rsaExpNegative) Initialize() error {
	return nil
}

func (l *rsaExpNegative) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaExpNegative) RunTest(c *x509.Certificate) (ResultStruct, error) {
	key, found := c.PublicKey.(*rsa.PublicKey)
	if !found {
		return ResultStruct{Result: Error}, nil
	}
	pubKey := key.E
	if pubKey < 0 {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "rsa_exp_negative",
		Description:   "RSA public key exponent must be positive",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &rsaExpNegative{}})
}
