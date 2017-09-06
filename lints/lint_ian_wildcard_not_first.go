// lint_br_ian_wildcard_not_first.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type brIANWildcardFirst struct{}

func (l *brIANWildcardFirst) Initialize() error {
	return nil
}

func (l *brIANWildcardFirst) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *brIANWildcardFirst) Execute(c *x509.Certificate) *LintResult {
	for _, dns := range c.IANDNSNames {
		for i := 1; i < len(dns); i++ {
			if dns[i] == '*' {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_ian_wildcard_not_first",
		Description:    "A wildcard MUST be in the first label of FQDN (ie not: www.*.com) (Only checks DNSName)",
		ReadableSource: "awslabs certlint",
		Source:         AWSLabs,
		EffectiveDate:  util.ZeroDate,
		Lint:           &brIANWildcardFirst{},
	})
}
