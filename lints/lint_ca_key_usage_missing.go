// lint_ca_key_usage_missing.go
/************************************************
RFC 5280: 4.2.1.3
Conforming CAs MUST include this extension in certificates that
   contain public keys that are used to validate digital signatures on
   other public key certificates or CRLs.  When present, conforming CAs
   SHOULD mark this extension as critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type caKeyUsageMissing struct{}

func (l *caKeyUsageMissing) Initialize() error {
	return nil
}

func (l *caKeyUsageMissing) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

func (l *caKeyUsageMissing) Execute(c *x509.Certificate) ResultStruct {
	if c.KeyUsage != x509.KeyUsage(0) {
		return ResultStruct{Result: Pass}
	} else {
		return ResultStruct{Result: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_key_usage_missing",
		Description:   "Root and Subordinate CA certificate keyUsage extension MUST be present",
		Source:        "BRs: 7.1.2.1, RFC 5280: 4.2.1.3",
		EffectiveDate: util.RFC3280Date,
		Lint:          &caKeyUsageMissing{},
	})
}
