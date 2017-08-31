// lint_rsa_public_exponent_too_small.go
/*******************************************************************************************************
"BRs: 6.1.6"
RSA: The CA SHALL confirm that the value of the public exponent is an odd number equal to 3 or more. Additionally, the public exponent SHOULD be in the range between 2^16+1 and 2^256-1. The modulus SHOULD also have the following characteristics: an odd number, not the power of a prime, and have no factors smaller than 752. [Source: Section 5.3.3, NIST SP 800-89].
*******************************************************************************************************/

package lints

import (
	"crypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaParsedTestsExpBounds struct{}

func (l *rsaParsedTestsExpBounds) Initialize() error {
	return nil
}

func (l *rsaParsedTestsExpBounds) CheckApplies(c *x509.Certificate) bool {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaParsedTestsExpBounds) Execute(c *x509.Certificate) ResultStruct {
	key := c.PublicKey.(*rsa.PublicKey)
	if key.E >= 3 { //If Cmp returns 1, means N > E
		return ResultStruct{Result: Pass}
	} else {
		return ResultStruct{Result: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_rsa_public_exponent_too_small",
		Description:   "RSA: Value of public exponent is an odd number equal to 3 or more.",
		Source:        "BRs: 6.1.6",
		EffectiveDate: util.CABV113Date,
		Lint:          &rsaParsedTestsExpBounds{},
	})
}
