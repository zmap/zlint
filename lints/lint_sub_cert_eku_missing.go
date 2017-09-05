// lint_sub_cert_eku_missing.go
/*******************************************************************************************************
BRs: 7.1.2.3
extKeyUsage (required)
Either the value id-kp-serverAuth [RFC5280] or id-kp-clientAuth [RFC5280] or both values MUST be present. id-kp-emailProtection [RFC5280] MAY be present. Other values SHOULD NOT be present.
*******************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subExtKeyUsage struct{}

func (l *subExtKeyUsage) Initialize() error {
	return nil
}

func (l *subExtKeyUsage) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c)
}

func (l *subExtKeyUsage) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	if util.IsExtInCert(c, util.EkuSynOid) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_eku_missing",
		Description:   "Subscriber certificates MUST have the extended key usage extension present",
		Source:        "BRs: 7.1.2.3",
		Type:          BRs,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subExtKeyUsage{},
	})
}
