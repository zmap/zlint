// lint_ca_crl_sign_not_set.go
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

type caCRLSignNotSet struct{}

func (l *caCRLSignNotSet) Initialize() error {
	return nil
}

func (l *caCRLSignNotSet) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *caCRLSignNotSet) Execute(c *x509.Certificate) ResultStruct {
	if c.KeyUsage&x509.KeyUsageCRLSign != 0 {
		return ResultStruct{Result: Pass}
	} else {
		return ResultStruct{Result: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_crl_sign_not_set",
		Description:   "Root and Subordinate CA certificate keyUsage extension's crlSign bit MUST be set",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caCRLSignNotSet{},
	})
}
