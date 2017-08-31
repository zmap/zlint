// lint_public_key_type_not_allowed.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type publicKeyAllowed struct {
	// Internal data here
}

func (l *publicKeyAllowed) Initialize() error {
	return nil
}

func (l *publicKeyAllowed) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *publicKeyAllowed) RunTest(c *x509.Certificate) (ResultStruct, error) {
	alg := c.PublicKeyAlgorithm
	if alg != x509.UnknownPublicKeyAlgorithm {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_public_key_type_not_allowed",
		Description:   "Certificates MUST have RSA, DSA, or ECDSA public key type",
		Source:        "BRs: 6.1.5",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &publicKeyAllowed{},
	})
}
