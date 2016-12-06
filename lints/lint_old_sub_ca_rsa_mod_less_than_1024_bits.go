// lint_old_sub_ca_rsa_mod_less_than_1024_bits.go
// CHANGE THIS COMMENT TO MATCH PROVIDENCE TEXT

package lints

import (

	"crypto/rsa"
	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type subCaModSize struct {
	// Internal data here
}

func (l *subCaModSize) Initialize() error {
	return nil
}

func (l *subCaModSize) CheckApplies(c *x509.Certificate) bool {
	issueDate := c.NotBefore
	endDate := c.NotAfter
	return (util.IsSubCA(c) && issueDate.Before(util.RsaDate2) && endDate.Before(util.RsaDate3))
}

func (l *subCaModSize) RunTest(c *x509.Certificate) (ResultStruct, error) {
	mod := c.PublicKey.(*rsa.PublicKey).N
	if mod.BitLen() < 1024 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:        "old_sub_ca_rsa_mod_less_than_1024_bits",
		Description: "In a validity period beginning on or before 31 dec 2010 and ending on or before 31 dec 2013, subordinate CA certificates using RSA public key algorithm must have 1024 bits of modulus",
		Providence:  "CAB: 6.1.5",
		// since effective date should be checked against end date in this specific case, putting time check into checkApplies instead, ZeroDate here to automatically pass NE test
		EffectiveDate: util.ZeroDate,
		Test:          &subCaModSize{}})
}
