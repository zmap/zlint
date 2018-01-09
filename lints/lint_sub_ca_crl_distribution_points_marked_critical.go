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

type subCACRLDistCrit struct{}

func (l *subCACRLDistCrit) Initialize() error {
	return nil
}

func (l *subCACRLDistCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.CrlDistOID)
}

func (l *subCACRLDistCrit) Execute(c *x509.Certificate) *LintResult {
	if e := util.GetExtFromCert(c, util.CrlDistOID); e.Critical {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_crl_distribution_points_marked_critical",
		Description:   "Subordinate CA Certificate: cRLDistributionPoints MUST be present and MUST NOT be marked critical.",
		Citation:      "BRs: 7.1.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCACRLDistCrit{},
	})
}
