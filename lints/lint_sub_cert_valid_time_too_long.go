// lint_ev_valid_time_too_long.go

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
		Name:           "e_sub_cert_valid_time_too_long",
		Description:    "CAs MUST NOT issue subscriber certificates with validity periods longer than 39 months regardless of circumstance.",
		ReadableSource: "BRs: 6.3.2",
		Source:         CABFBaselineRequirements,
		EffectiveDate:  util.SubCert39Month,
		Lint:           &subCertValidTimeTooLong{},
	})
}
