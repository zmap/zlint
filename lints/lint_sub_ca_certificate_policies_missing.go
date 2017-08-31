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

type subCACertPolicyMissing struct {
	// Internal data here
}

func (l *subCACertPolicyMissing) Initialize() error {
	return nil
}

func (l *subCACertPolicyMissing) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubCA(c)
}

func (l *subCACertPolicyMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(c, util.CertPolicyOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_certificate_policies_missing",
		Description:   "Subordinate CA certificates must have a certificatePolicies extension",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCACertPolicyMissing{},
	})
}
