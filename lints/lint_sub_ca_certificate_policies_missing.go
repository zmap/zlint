// lint_sub_ca_certificate_policies_missing.go
/************************************************
BRs: 7.1.2.2a certificatePolicies
This extension MUST be present and SHOULD NOT be marked critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCACertPolicyMissing struct{}

func (l *subCACertPolicyMissing) Initialize() error {
	return nil
}

func (l *subCACertPolicyMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c)
}

func (l *subCACertPolicyMissing) Execute(c *x509.Certificate) *LintResult {
	if util.IsExtInCert(c, util.CertPolicyOID) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_certificate_policies_missing",
		Description:   "Subordinate CA certificates must have a certificatePolicies extension",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCACertPolicyMissing{},
	})
}
