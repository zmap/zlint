// lint_rsa_no_public_key.go
package lints

import (
	"crypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaParsedPubKeyExist struct {
	// Internal data here
}

func (l *rsaParsedPubKeyExist) Initialize() error {
	return nil
}

func (l *rsaParsedPubKeyExist) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaParsedPubKeyExist) RunTest(c *x509.Certificate) (ResultStruct, error) {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_rsa_no_public_key",
		Description:   "The RSA public key should be present",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &rsaParsedPubKeyExist{},
	})
}
