// lint_ca_key_cert_sign_not_set.go
/************************************************
BRs: 7.1.2.1b
This extension MUST be present and MUST be marked critical. Bit positions for keyCertSign and cRLSign MUST be set.
If the Root CA Private Key is used for signing OCSP responses, then the digitalSignature bit MUST be set.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type caKeyCertSignNotSet struct{}

func (l *caKeyCertSignNotSet) Initialize() error {
	return nil
}

func (l *caKeyCertSignNotSet) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *caKeyCertSignNotSet) Execute(c *x509.Certificate) *LintResult {
	if c.KeyUsage&x509.KeyUsageCertSign != 0 {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_key_cert_sign_not_set",
		Description:   "Root CA Certificate: Bit positions for keyCertSign and cRLSign MUST be set.",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caKeyCertSignNotSet{},
	})
}
