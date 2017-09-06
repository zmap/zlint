// lint_validity_time_not_positive.go
/************************************************
Change this to match source TEXT
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type validityNegative struct{}

func (l *validityNegative) Initialize() error {
	return nil
}

func (l *validityNegative) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *validityNegative) Execute(c *x509.Certificate) *LintResult {
	if c.NotBefore.After(c.NotAfter) {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_validity_time_not_positive",
		Description:    "Certificates MUST have a positive time for which they are valid",
		ReadableSource: "AWSLabs certlint",
		Source:         AWSLabs,
		EffectiveDate:  util.ZeroDate,
		Lint:           &validityNegative{},
	})
}
