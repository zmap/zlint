// lint_sub_cert_valid_time_longer_than_825_days.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertValidTimeLongerThan825Days struct{}

func (l *subCertValidTimeLongerThan825Days) Initialize() error {
	return nil
}

func (l *subCertValidTimeLongerThan825Days) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertValidTimeLongerThan825Days) Execute(c *x509.Certificate) *LintResult {
	if c.NotBefore.AddDate(0, 0, 825).Before(c.NotAfter) {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_valid_time_longer_than_825_days",
		Description:   "Subscriber Certificates issued after 1 March 2018 MUST have a Validity Period no greater than 825 days.",
		Citation:      "BRs: 6.3.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.SubCert825Days,
		Lint:          &subCertValidTimeLongerThan825Days{},
	})
}
