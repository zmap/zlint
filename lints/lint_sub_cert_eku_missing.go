// lint_sub_cert_eku_missing.go
/*******************************************************************************************************
CAB: 7.1.2.3
extKeyUsage (required)
Either the value id-kp-serverAuth [RFC5280] or id-kp-clientAuth [RFC5280] or both values MUST be present. id-kp-emailProtection [RFC5280] MAY be present. Other values SHOULD NOT be present.
*******************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subExtKeyUsage struct {
	// Internal data here
}

func (l *subExtKeyUsage) Initialize() error {
	return nil
}

func (l *subExtKeyUsage) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return !util.IsCaCert(c)
}

func (l *subExtKeyUsage) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if util.IsExtInCert(c, util.EkuSynOid) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_eku_missing",
		Description:   "Subscriber certificates must have the extended key usage extension present",
		Providence:    "CAB: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subExtKeyUsage{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESubCertEkuMissing = result },
	})
}
