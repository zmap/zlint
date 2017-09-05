// lint_sub_ca_certificate_policies_marked_critical.go
/************************************************
BRs: 7.1.2.2a certificatePolicies
This extension MUST be present and SHOULD NOT be marked critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCACertPolicyCrit struct{}

func (l *subCACertPolicyCrit) Initialize() error {
	return nil
}

func (l *subCACertPolicyCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *subCACertPolicyCrit) Execute(c *x509.Certificate) *LintResult {
	if e := util.GetExtFromCert(c, util.CertPolicyOID); e.Critical {
		return &LintResult{Status: Warn}
	} else {
		return &LintResult{Status: Pass}
	}

}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_certificate_policies_marked_critical",
		Description:   "Subordinate CA certificates certificatePolicies extension should not be marked as critical",
		Source:        "BRs: 7.1.2.2",
		Type:          CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCACertPolicyCrit{},
	})
}
