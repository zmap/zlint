// lint_sub_cert_valid_time_too_long.go

package lints

import (
	"time"
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
	july_1_2016 := time.Date(2016, time.July, 1, 0, 0, 0, 0, time.UTC)
	march_1_2018 := time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC)
	if c.NotBefore.Before(july_1_2016) {
		// Vacuously passed
		return &LintResult{Status: Pass}
	}
	if c.NotBefore.After(march_1_2018) {
		if c.NotBefore.AddDate(0, 0, 825).Before(c.NotAfter) {
			// NotAfter - NotBefore > 825 days
			return &LintResult{Status: Error}
		} else {
			return &LintResult{Status: Pass}
		}
	} else {
		// Between 2016/July/01 and 2018/March/01: 39 months
		if c.NotBefore.AddDate(0, 39, 0).Before(c.NotAfter) {
			return &LintResult{Status: Error}
		}
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_valid_time_too_long",
		Description:   "Subscriber Certificates issued after 1 March 2018 MUST have a Validity Period no greater than 825 days. Subscriber Certificates issued after 1 July 2016 but prior to 1 March 2018 MUST have a Validity Period no greater than 39 months. ",
		Citation:      "BRs: 6.3.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.SubCert39Month,
		Lint:          &subCertValidTimeTooLong{},
	})
}
