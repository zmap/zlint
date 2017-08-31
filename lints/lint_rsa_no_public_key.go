// lint_rsa_no_public_key.go
package lints

import (
	"crypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaParsedPubKeyExist struct{}

func (l *rsaParsedPubKeyExist) Initialize() error {
	return nil
}

func (l *rsaParsedPubKeyExist) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaParsedPubKeyExist) Execute(c *x509.Certificate) ResultStruct {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		return ResultStruct{Result: Error}
	} else {
		return ResultStruct{Result: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_rsa_no_public_key",
		Description:   "The RSA public key should be present",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &rsaParsedPubKeyExist{},
	})
}
