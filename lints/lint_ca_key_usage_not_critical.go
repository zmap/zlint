// lint_ca_key_usage_not_critical.go
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

type caKeyUsageNotCrit struct{}

func (l *caKeyUsageNotCrit) Initialize() error {
	return nil
}

func (l *caKeyUsageNotCrit) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *caKeyUsageNotCrit) Execute(c *x509.Certificate) * LintResult{
	if e := util.GetExtFromCert(c, util.KeyUsageOID); e.Critical {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_key_usage_not_critical",
		Description:   "Root and Subordinate CA certificate keyUsage extension MUST be marked as critical",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caKeyUsageNotCrit{},
	})
}
