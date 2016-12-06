// lint_public_key_type_not_allowed.go

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
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
		Name:          "public_key_type_not_allowed",
		Description:   "Certificates must have RSA, DSA, or ECDSA public key type.",
		Providence:    "CAB: 6.1.5",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &publicKeyAllowed{}})
}
