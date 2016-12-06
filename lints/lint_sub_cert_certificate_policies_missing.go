// lint_sub_cert_certificate_policies_missing.go
/******************************************************************************
CAB: 7.1.2.3
certificatePolicies
This extension MUST be present and SHOULD NOT be marked critical.
******************************************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type subCertPolicy struct {
	// Internal data here
}

func (l *subCertPolicy) Initialize() error {
	return nil
}

func (l *subCertPolicy) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return !util.IsCaCert(c)
}

func (l *subCertPolicy) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if util.IsExtInCert(c, util.CertPolicyOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "sub_cert_certificate_policies_missing",
		Description:   "Subscriber certificates should have certificates policies extension present",
		Providence:    "CAB: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertPolicy{}})
}
