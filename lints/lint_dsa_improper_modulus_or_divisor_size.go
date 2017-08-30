// lint_dsa_improper_modulus_or_divisor_size.go
/************************************************
BRs: 6.1.5
Certificates MUST meet the following requirements for algorithm type and key size.
Minimum DSA modulus and divisor size (bits)***: L=2048,	N=224 or L=2048, N=256.
**As a note, this points to FIPS 186-4 for further clarification**
************************************************/

package lints

import (
	"crypto/dsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaImproperSize struct {
	// Internal data here
}

func (l *dsaImproperSize) Initialize() error {
	return nil
}

func (l *dsaImproperSize) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaImproperSize) RunTest(c *x509.Certificate) (ResultStruct, error) {
	theKey := c.PublicKey.(*dsa.PublicKey)
	lbit := theKey.Parameters.P.BitLen()
	nbit := theKey.Parameters.Q.BitLen()
	if lbit == 2048 && nbit == 224 ||
		lbit == 2048 && nbit == 256 ||
		lbit == 3072 && nbit == 256 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_improper_modulus_or_divisor_size",
		Description:   "Certificates MUST meet the following requirements for algorithm type and key size: L=2048, N=224,256 minimum DSA",
		Source:        "BRs: 6.1.5",
		EffectiveDate: util.ZeroDate,
		Test:          &dsaImproperSize{},
	})
}
