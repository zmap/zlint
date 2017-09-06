// lint_ian_dns_name_includes_null_char.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IANDNSNull struct{}

func (l *IANDNSNull) Initialize() error {
	return nil
}

func (l *IANDNSNull) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANDNSNull) Execute(c *x509.Certificate) *LintResult {
	for _, dns := range c.IANDNSNames {
		for i := 0; i < len(dns); i++ {
			if dns[i] == 0 {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ian_dns_name_includes_null_char",
		Description:   "DNSName MUST NOT include a null character",
		Citation:      "awslabs certlint",
		Source:        AWSLabs,
		EffectiveDate: util.ZeroDate,
		Lint:          &IANDNSNull{},
	})
}
