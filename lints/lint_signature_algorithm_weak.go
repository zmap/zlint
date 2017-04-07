// lint_signature_algorithm_weak.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"math/big"
)

type WeakAlgo struct {
	// Internal data here
}

func (l *WeakAlgo) Initialize() error {
	return nil
}

func (l *WeakAlgo) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *WeakAlgo) RunTest(c *x509.Certificate) (ResultStruct, error) {
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
		Name:          "w_signature_algorithm_weak",
		Description:   "A weak alg should have at least 64 bits on entropy. A good alg should have at least 20 bits",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &WeakAlgo{}})
}
