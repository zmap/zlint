package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rootCAKeyUsagePresent struct{}

func (l *rootCAKeyUsagePresent) Initialize() error {
	return nil
}

func (l *rootCAKeyUsagePresent) CheckApplies(c *x509.Certificate) bool {
	return util.IsRootCA(c)
}

func (l *rootCAKeyUsagePresent) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	if util.IsExtInCert(c, util.KeyUsageOID) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_root_ca_key_usage_present",
		Description:   "Root CA certificates MUST have Key Usage Extension Present",
		Citation:      "BRs: 7.1.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.RFC2459Date,
		Lint:          &rootCAKeyUsagePresent{},
	})
}
