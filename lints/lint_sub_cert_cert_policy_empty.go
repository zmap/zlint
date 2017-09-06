// lint_sub_cert_cert_policy_empty.go
/********************************************************************************************************************
BRs: 7.1.6.4
Subscriber Certificates
A Certificate issued to a Subscriber MUST contain one or more policy identifier(s), defined by the Issuing CA, in
the Certificate’s certificatePolicies extension that indicates adherence to and complIANce with these Requirements.
CAs complying with these Requirements MAY also assert one of the reserved policy OIDs in such Certificates.
*********************************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertPolicyEmpty struct{}

func (l *subCertPolicyEmpty) Initialize() error {
	return nil
}

func (l *subCertPolicyEmpty) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c)
}

func (l *subCertPolicyEmpty) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	if util.IsExtInCert(c, util.CertPolicyOID) && c.PolicyIdentifiers != nil {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_cert_policy_empty",
		Description:   "Subscriber certificates must contain at least one policy identifier that indicates adherence to CAB standards",
		Citation:      "BRs: 7.1.6.4",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCertPolicyEmpty{},
	})
}
