// lint_root_ca_extended_key_usage_present.go
/************************************************
BRs: 7.1.2.1d extendedKeyUsage
This extension MUST NOT be present.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rootCAContainsEKU struct{}

func (l *rootCAContainsEKU) Initialize() error {
	return nil
}

func (l *rootCAContainsEKU) CheckApplies(c *x509.Certificate) bool {
	return util.IsRootCA(c)
}

func (l *rootCAContainsEKU) Execute(c *x509.Certificate) * LintResult{
	// Add actual lint here
	if util.IsExtInCert(c, util.EkuSynOid) {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_root_ca_extended_key_usage_present",
		Description:   "Root CA Certificate: extendedKeyUsage MUST NOT be present.t",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &rootCAContainsEKU{},
	})
}
