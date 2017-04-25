// lint_sub_cert_crl_distribution_points_marked_critical.go
/*******************************************************************************************************
CAB: 7.1.2.3
cRLDistributionPoints
This extension MAY be present. If present, it MUST NOT be marked critical, and it MUST contain the HTTP
URL of the CA’s CRL service.
*******************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCrlDistCrit struct {
	// Internal data here
}

func (l *subCrlDistCrit) Initialize() error {
	return nil
}

func (l *subCrlDistCrit) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *subCrlDistCrit) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	e := util.GetExtFromCert(c, util.CrlDistOID)
	if e.Critical == false {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_crl_distribution_points_marked_critical",
		Description:   "Subscriber certificate cRLDistributionPoints extension must not be marked critical if present",
		Providence:    "CAB: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCrlDistCrit{},
		updateReport: func(report *LintReport, result ResultStruct) {
			report.ESubCertCrlDistributionPointsMarkedCritical = result
		},
	})
}
