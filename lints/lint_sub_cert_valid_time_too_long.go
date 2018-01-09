// lint_sub_cert_valid_time_too_long.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertValidTimeTooLong struct{}

func (l *subCertValidTimeTooLong) Initialize() error {
	return nil
}

func (l *subCertValidTimeTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertValidTimeTooLong) Execute(c *x509.Certificate) *LintResult {
	if c.NotBefore.AddDate(0, 39, 0).Before(c.NotAfter) {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_valid_time_too_long",
		Description:   "Subscriber Certificates issued after 1 March 2018 MUST have a Validity Period no greater than 825 days. Subscriber Certificates issued after 1 July 2016 but prior to 1 March 2018 MUST have a Validity Period no greater than 39 months.",
		Citation:      "BRs: 6.3.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.SubCert39Month,
		Lint:          &subCertValidTimeTooLong{},
	})
}
