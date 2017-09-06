// lint_ian_dns_name_starts_with_period.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type IANDNSPeriod struct{}

func (l *IANDNSPeriod) Initialize() error {
	return nil
}

func (l *IANDNSPeriod) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANDNSPeriod) Execute(c *x509.Certificate) *LintResult {
	for _, dns := range c.IANDNSNames {
		if strings.HasPrefix(dns, ".") {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_ian_dns_name_starts_with_period",
		Description:    "DNSName MUST NOT start with a period",
		ReadableSource: "awslabs certlint",
		Source:         AWSLabs,
		EffectiveDate:  util.ZeroDate,
		Lint:           &IANDNSPeriod{},
	})
}
