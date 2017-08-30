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

type subCertPolicyCrit struct {
	// Internal data here
}

func (l *subCertPolicyCrit) Initialize() error {
	return nil
}

func (l *subCertPolicyCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *subCertPolicyCrit) RunTest(c *x509.Certificate) (ResultStruct, error) {
	e := util.GetExtFromCert(c, util.CertPolicyOID)
	if e.Critical == false {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_cert_certificate_policies_marked_critical",
		Description:   "Subscriber Certificate: certificatePolicies MUST be present and SHOULD NOT be marked critical.",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertPolicyCrit{},
	})
}
