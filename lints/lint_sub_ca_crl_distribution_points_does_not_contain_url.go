// lint_sub_ca_crl_distribution_points_does_not_contain_url.go
/************************************************
BRs: 7.1.2.2b cRLDistributionPoints
This extension MUST be present and MUST NOT be marked critical.
It MUST contain the HTTP URL of the CA’s CRL service.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type subCACRLDistNoUrl struct {
	// Internal data here
}

func (l *subCACRLDistNoUrl) Initialize() error {
	return nil
}

func (l *subCACRLDistNoUrl) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.CrlDistOID)
}

func (l *subCACRLDistNoUrl) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, s := range c.CRLDistributionPoints {
		if strings.HasPrefix(s, "http://") {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_crl_distribution_points_does_not_contain_url",
		Description:   "Subordinate CA Certificate: cRLDistributionPoints MUST contain the HTTP URL of the CA's CRL service.",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCACRLDistNoUrl{},
	})
}
