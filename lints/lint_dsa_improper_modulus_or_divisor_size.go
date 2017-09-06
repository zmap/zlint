package lints

import (
	"crypto/dsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaImproperSize struct{}

func (l *dsaImproperSize) Initialize() error {
	return nil
}

func (l *dsaImproperSize) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaImproperSize) Execute(c *x509.Certificate) *LintResult {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return &LintResult{Status: NA}
	}
	L := dsaKey.Parameters.P.BitLen()
	N := dsaKey.Parameters.Q.BitLen()
	if (L == 2048 && N == 224) || (L == 2048 && N == 256) || (L == 3072 && N == 256) {
		return &LintResult{Status: Pass}
	}
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_dsa_improper_modulus_or_divisor_size",
		Description:    "Certificates MUST meet the following requirements for algorithm type and key size: L=2048, N=224,256 minimum DSA",
		ReadableSource: "BRs: 6.1.5",
		Source:         CABFBaselineRequirements,
		EffectiveDate:  util.ZeroDate,
		Lint:           &dsaImproperSize{},
	})
}
