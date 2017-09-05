// lint_sub_cert_crl_distribution_points_does_not_contain_url.go
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
	"strings"
)

type subCRLDistNoURL struct{}

func (l *subCRLDistNoURL) Initialize() error {
	return nil
}

func (l *subCRLDistNoURL) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *subCRLDistNoURL) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	for _, s := range c.CRLDistributionPoints {
		if strings.HasPrefix(s, "http://") {
			return &LintResult{Status: Pass}
		}
	}
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_crl_distribution_points_does_not_contain_url",
		Description:   "Subscriber certificate cRLDistributionPoints extension must contain the HTTP URL of the CA’s CRL service",
		Source:        "BRs: 7.1.2.3",
		Type:          BRs,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCRLDistNoURL{},
	})
}
