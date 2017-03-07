// lint_rsa_exp_negative.go

package lints

import (
	"crypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
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
	key, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		return ResultStruct{Result: Error}, nil
	}
	if key.E < 0 {
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
