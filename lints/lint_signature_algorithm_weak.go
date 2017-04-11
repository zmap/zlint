// lint_serial_number_too_short.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"math/big"
)

type SNTooShort struct {
	// Internal data here
}

func (l *SNTooShort) Initialize() error {
	return nil
}

func (l *SNTooShort) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SNTooShort) RunTest(c *x509.Certificate) (ResultStruct, error) {
	z := big.NewInt(0)
	if c.SignatureAlgorithm == x509.SHA1WithRSA || c.SignatureAlgorithm == x509.DSAWithSHA1 || c.SignatureAlgorithm == x509.ECDSAWithSHA1 {
		z.Exp(big.NewInt(2), big.NewInt(64), nil)
		if c.SerialNumber.Cmp(z) == -1 {
			return ResultStruct{Result: Warn}, nil
		}
	}
	if c.SignatureAlgorithm == x509.SHA256WithRSA || c.SignatureAlgorithm == x509.SHA384WithRSA || c.SignatureAlgorithm == x509.SHA512WithRSA ||
		c.SignatureAlgorithm == x509.DSAWithSHA256 || c.SignatureAlgorithm == x509.ECDSAWithSHA256 || c.SignatureAlgorithm == x509.ECDSAWithSHA384 ||
		c.SignatureAlgorithm == x509.ECDSAWithSHA512 {
		z.Exp(big.NewInt(2), big.NewInt(20), nil)
		if c.SerialNumber.Cmp(z) == -1 {
			return ResultStruct{Result: Warn}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_serial_number_too_short",
		Description:   "A weak algorithm should have at least 64 bits of entropy in its serial number. A good algorithm should have at least 20 bits of entropy in its serial number.",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SNTooShort{}})
}
