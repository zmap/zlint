// lint_old_sub_cert_rsa_mod_less_than_1024_bits.go
// CHANGE THIS COMMENT TO MATCH PROVIDENCE TEXT

package lints

import (
	"crypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subModSize struct {
	// Internal data here
}

func (l *subModSize) Initialize() error {
	return nil
}

func (l *subModSize) CheckApplies(c *x509.Certificate) bool {
	endDate := c.NotAfter
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA && !util.IsCaCert(c) && endDate.Before(util.RsaDate3)
}

func (l *subModSize) RunTest(c *x509.Certificate) (ResultStruct, error) {
	key := c.PublicKey.(*rsa.PublicKey)
	if key.N.BitLen() < 1024 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:        "e_old_sub_cert_rsa_mod_less_than_1024_bits",
		Description: "In a validity period ending on or before 31 Dec 2013, subscriber certificates using RSA public key algorithm MUST use a 1024 bit modulus",
		Providence:  "CAB: 6.1.5",
		// since effective date should be checked against end date in this specific case, putting time check into checkApplies instead, ZeroDate here to automatically pass NE test
		EffectiveDate: util.ZeroDate,
		Test:          &subModSize{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.EOldSubCertRsaModLessThan_1024Bits = result },
	})
}
