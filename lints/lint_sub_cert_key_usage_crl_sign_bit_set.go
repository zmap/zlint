// lint_sub_key_usage_crl_sign_bit_set.go
/**************************************************************************
BRs: 7.1.2.3
keyUsage (optional)
If present, bit positions for keyCertSign and cRLSign MUST NOT be set.
***************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCrlSignAllowed struct{}

func (l *subCrlSignAllowed) Initialize() error {
	return nil
}

func (l *subCrlSignAllowed) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID) && !util.IsCACert(c)
}

func (l *subCrlSignAllowed) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	if (c.KeyUsage & x509.KeyUsageCRLSign) == x509.KeyUsageCRLSign {
		return &LintResult{Status: Error}
	} else { //key usage doesn't allow cert signing or isn't present
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_key_usage_crl_sign_bit_set",
		Description:   "Subscriber Certificate: keyUsage if present, bit positions for keyCertSign and cRLSign MUST NOT be set.",
		Citation:      "BRs: 7.1.2.3",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCrlSignAllowed{},
	})
}
