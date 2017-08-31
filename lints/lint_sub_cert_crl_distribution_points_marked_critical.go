// lint_sub_cert_crl_distribution_points_marked_critical.go
/*******************************************************************************************************
BRs: 7.1.2.3
cRLDistributionPoints
This extension MAY be present. If present, it MUST NOT be marked critical, and it MUST contain the HTTP
URL of the CA’s CRL service.
*******************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCrlDistCrit struct{}

func (l *subCrlDistCrit) Initialize() error {
	return nil
}

func (l *subCrlDistCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *subCrlDistCrit) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	e := util.GetExtFromCert(c, util.CrlDistOID)
	if e.Critical == false {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_crl_distribution_points_marked_critical",
		Description:   "Subscriber Certiifcate: cRLDistributionPoints MUST NOT be marked critical, and MUST contain the HTTP URL of the CA's CRL service.",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCrlDistCrit{},
	})
}
