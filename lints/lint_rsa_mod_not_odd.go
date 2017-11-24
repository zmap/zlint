// lint_rsa_mod_not_odd.go
/*******************************************************************************************************
"BRs: 6.1.6"
RSA: The CA SHALL confirm that the value of the public exponent is an odd number equal to 3 or	more. Additionally,	the public exponent SHOULD be in the range between 2^16+1 and 2^256-1. The modulus SHOULD also have the following characteristics: an odd number, not the power of a prime, and have no factors smaller than 752. [Citation: Section 5.3.3, NIST SP 800-89].
*******************************************************************************************************/

package lints

import (
	"crypto/rsa"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaParsedTestsKeyModOdd struct{}

func (l *rsaParsedTestsKeyModOdd) Initialize() error {
	return nil
}

func (l *rsaParsedTestsKeyModOdd) CheckApplies(c *x509.Certificate) bool {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaParsedTestsKeyModOdd) Execute(c *x509.Certificate) *LintResult {
	key := c.PublicKey.(*rsa.PublicKey)
	z := big.NewInt(0)
	if (z.Mod(key.N, big.NewInt(2)).Cmp(big.NewInt(1))) == 0 {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Warn}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_rsa_mod_not_odd",
		Description:   "RSA: Modulus SHOULD also have the following characteristics: an odd number",
		Citation:      "BRs: 6.1.6",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABV113Date,
		Lint:          &rsaParsedTestsKeyModOdd{},
	})
}
