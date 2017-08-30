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

type subCRLDistNoURL struct {
	// Internal data here
}

func (l *subCRLDistNoURL) Initialize() error {
	return nil
}

func (l *subCRLDistNoURL) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *subCRLDistNoURL) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	for _, s := range c.CRLDistributionPoints {
		if strings.HasPrefix(s, "http://") {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_crl_distribution_points_does_not_contain_url",
		Description:   "Subscriber certificate cRLDistributionPoints extension must contain the HTTP URL of the CA’s CRL service",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCRLDistNoURL{},
	})
}
