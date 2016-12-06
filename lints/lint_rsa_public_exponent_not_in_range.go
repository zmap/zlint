// lint_rsa_public_exponent_not_in_range.go
/*******************************************************************************************************
"CAB: 6.1.6"
RSA: The CA SHALL confirm that the value of the public exponent is an odd number equal to 3 or more. Additionally, the public exponent SHOULD be in the range between 2^16+1 and 2^256-1. The modulus SHOULD also have the following characteristics: an odd number, not the power of a prime, and have no factors smaller than 752. [Source: Section 5.3.3, NIST SP 800-89].
*******************************************************************************************************/

package lints

import (

	"crypto/rsa"
	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
	"math/big"
)

type rsaParsedTestsExpInRange struct {
	// Internal data here
}

func (l *rsaParsedTestsExpInRange) Initialize() error {
	return nil
}

func (l *rsaParsedTestsExpInRange) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaParsedTestsExpInRange) RunTest(c *x509.Certificate) (ResultStruct, error) {
	pubKey := c.PublicKey.(*rsa.PublicKey).E
	lowerBound := 65536 // 2^16 + 1
	var upperBound big.Int
	upperBound.Exp(big.NewInt(2), big.NewInt(256), nil)

	if pubKey > lowerBound && upperBound.Cmp(big.NewInt(int64(pubKey))) == 1 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "rsa_public_exponent_not_in_range",
		Description:   "The RSA public exponent SHOULD be in the range between 2^16 + 1 and 2^256 - 1",
		Providence:    "CAB: 6.1.6",
		EffectiveDate: util.CABV113Date,
		Test:          &rsaParsedTestsExpInRange{}})
}
