// lint_sub_ca_certificate_policies_missing.go
/************************************************
CAB: 7.1.2.2a certificatePolicies
This extension MUST be present and SHOULD NOT be marked critical.
************************************************/

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
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
		Name:          "sub_ca_certificate_policies_missing",
		Description:   "Subordinate CA certificates must have a certificatePolicies extension",
		Providence:    "CAB: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCACertPolicyMissing{}})
}
