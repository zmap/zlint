// lint_rsa_public_exponent_not_in_range.go
/*******************************************************************************************************
"BRs: 6.1.6"
RSA: The CA SHALL confirm that the value of the public exponent is an odd number equal to 3 or more. Additionally, the public exponent SHOULD be in the range between 2^16+1 and 2^256-1. The modulus SHOULD also have the following characteristics: an odd number, not the power of a prime, and have no factors smaller than 752. [Source: Section 5.3.3, NIST SP 800-89].
*******************************************************************************************************/

package lints

import (
	"crypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"math/big"
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

func (l *rsaParsedTestsExpInRange) RunTest(c *x509.Certificate) (ResultStruct, error) {
	key := c.PublicKey.(*rsa.PublicKey)
	exponent := key.E
	const lowerBound = 65536 // 2^16 + 1
	//	if l.upperBound.Cmp(big.NewInt(0)) == 0 {
	//		l.upperBound.Exp(big.NewInt(2), big.NewInt(256), nil)
	//	}
	if exponent > lowerBound && l.upperBound.Cmp(big.NewInt(int64(exponent))) == 1 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_rsa_public_exponent_not_in_range",
		Description:   "RSA: Public exponent SHOULD be in the range between 2^16 + 1 and 2^256 - 1",
		Source:        "BRs: 6.1.6",
		EffectiveDate: util.CABV113Date,
		Test:          &rsaParsedTestsExpInRange{},
	})
}
