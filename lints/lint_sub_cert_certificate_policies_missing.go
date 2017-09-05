// lint_sub_cert_certificate_policies_missing.go
/******************************************************************************
BRs: 7.1.2.3
certificatePolicies
This extension MUST be present and SHOULD NOT be marked critical.
******************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertPolicy struct{}

func (l *subCertPolicy) Initialize() error {
	return nil
}

func (l *subCertPolicy) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c)
}

func (l *subCertPolicy) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	if util.IsExtInCert(c, util.CertPolicyOID) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_certificate_policies_missing",
		Description:   "Subscriber Certificate: certificatePolicies MUST be present and SHOULD NOT be marked critical.",
		Source:        "BRs: 7.1.2.2",
		Type:          BRs,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCertPolicy{},
	})
}
