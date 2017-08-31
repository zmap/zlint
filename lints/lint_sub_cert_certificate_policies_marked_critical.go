// lint_sub_cert_certificate_policies_marked_critical.go
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

type subCertPolicyCrit struct{}

func (l *subCertPolicyCrit) Initialize() error {
	return nil
}

func (l *subCertPolicyCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *subCertPolicyCrit) Execute(c *x509.Certificate) * LintResult{
	e := util.GetExtFromCert(c, util.CertPolicyOID)
	if e.Critical == false {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Warn}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_cert_certificate_policies_marked_critical",
		Description:   "Subscriber Certificate: certificatePolicies MUST be present and SHOULD NOT be marked critical.",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCertPolicyCrit{},
	})
}
