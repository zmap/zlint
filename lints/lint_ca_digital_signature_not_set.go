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

type caDigSignNotSet struct{}

func (l *caDigSignNotSet) Initialize() error {
	return nil
}

func (l *caDigSignNotSet) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *caDigSignNotSet) Execute(c *x509.Certificate) *LintResult {
	if c.KeyUsage&x509.KeyUsageDigitalSignature != 0 {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Notice}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_ca_digital_signature_not_set",
		Description:   "Root and Subordinate CA Certificates that wish to use their private key for signing OCSP responses will not be able to without their digital signature set",
		Citation:      "BRs: 7.1.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caDigSignNotSet{},
	})
}
