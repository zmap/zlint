// lint_sub_ca_crl_distribution_points_missing.go
/************************************************
BRs: 7.1.2.2b cRLDistributionPoints
This extension MUST be present and MUST NOT be marked critical.
It MUST contain the HTTP URL of the CA’s CRL service.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCACRLDistMissing struct {
	// Internal data here
}

func (l *subCACRLDistMissing) Initialize() error {
	return nil
}

func (l *subCACRLDistMissing) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubCA(c)
}

func (l *subCACRLDistMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(c, util.CrlDistOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_crl_distribution_points_missing",
		Description:   "Subordinate CA Certificate: cRLDistributionPoints MUST be present and MUST NOT be marked critical.",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCACRLDistMissing{},
	})
}
