/*******************************************************************************************************
"BRs: 6.1.6"
RSA: The CA SHALL confirm that the value of the public exponent is an odd number equal to 3 or more. Additionally, the public exponent SHOULD be in the range between 2^16+1 and 2^256-1. The modulus SHOULD also have the following characteristics: an odd number, not the power of a prime, and have no factors smaller than 752. [Citation: Section 5.3.3, NIST SP 800-89].
*******************************************************************************************************/

package lints

import (
	"crypto/rsa"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaParsedTestsExpInRange struct {
	upperBound *big.Int
}

func (l *rsaParsedTestsExpInRange) Initialize() error {
	l.upperBound = &big.Int{}
	l.upperBound.Exp(big.NewInt(2), big.NewInt(256), nil)
	return nil
}

func (l *rsaParsedTestsExpInRange) CheckApplies(c *x509.Certificate) bool {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaParsedTestsExpInRange) Execute(c *x509.Certificate) *LintResult {
	key := c.PublicKey.(*rsa.PublicKey)
	exponent := key.E
	const lowerBound = 65536 // 2^16 + 1
	if exponent > lowerBound && l.upperBound.Cmp(big.NewInt(int64(exponent))) == 1 {
		return &LintResult{Status: Pass}
	}
	return &LintResult{Status: Warn}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_rsa_public_exponent_not_in_range",
		Description:   "RSA: Public exponent SHOULD be in the range between 2^16 + 1 and 2^256 - 1",
		Citation:      "BRs: 6.1.6",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABV113Date,
		Lint:          &rsaParsedTestsExpInRange{},
	})
}
