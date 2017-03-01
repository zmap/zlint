// lint_ca_dig_sign_not_set.go
/************************************************
CAB: 7.1.2.1b
This extension MUST be present and MUST be marked critical. Bit positions for keyCertSign and cRLSign MUST be set.
If the Root CA Private Key is used for signing OCSP responses, then the digitalSignature bit MUST be set.
************************************************/

package lints

import (
	"github.com/zmap/zgrab/ztools/x509"
	"github.com/zmap/zlint/util"
)

type caDigSignNotSet struct {
	// Internal data here
}

func (l *caDigSignNotSet) Initialize() error {
	return nil
}

func (l *caDigSignNotSet) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.IsCA && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *caDigSignNotSet) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.KeyUsage&x509.KeyUsageDigitalSignature != 0 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "ca_dig_sign_not_set",
		Description:   "Root & Subordinate CA Certificates that wish to use their private key for signing OCSP responses will not be able to with out digital signature set",
		Providence:    "CAB: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &caDigSignNotSet{}})
}
