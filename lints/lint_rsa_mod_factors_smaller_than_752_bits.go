// lint_rsa_mod_factors_smaller_than_752.go
/**************************************************************************************************
6.1.6. Public Key Parameters Generation and Quality Checking
RSA: The CA SHALL confirm that the value of the public exponent is an odd number equal to 3 or more. Additionally, the public exponent SHOULD be in the range between 216+1 and 2256-1. The modulus SHOULD also have the following characteristics: an odd number, not the power of a prime, and have no factors smaller than 752. [Source: Section 5.3.3, NIST SP 800‐89].
**************************************************************************************************/

package lints

import (
	"crypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaModSmallFactor struct {
	// Internal data here
}

func (l *rsaModSmallFactor) Initialize() error {
	return nil
}

func (l *rsaModSmallFactor) CheckApplies(c *x509.Certificate) bool {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaModSmallFactor) RunTest(c *x509.Certificate) (ResultStruct, error) {
	key := c.PublicKey.(*rsa.PublicKey)
	if util.PrimeNoSmallerThan752(key.N) {
		return ResultStruct{Result: Pass}, nil
	}
	return ResultStruct{Result: Warn}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_rsa_mod_factors_smaller_than_752",
		Description:   "RSA: Modulus SHOULD also have the following characteristics: no factors smaller than 752",
		Source:        "BRs: 6.1.6",
		EffectiveDate: util.CABV113Date,
		Test:          &rsaModSmallFactor{},
	})
}
