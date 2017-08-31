// lint_br_ian_bare_wildcard.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type brIANBareWildcard struct{}

func (l *brIANBareWildcard) Initialize() error {
	return nil
}

func (l *brIANBareWildcard) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *brIANBareWildcard) Execute(c *x509.Certificate) LintResult {
	for _, dns := range c.IANDNSNames {
		if strings.HasSuffix(dns, "*") {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ian_bare_wildcard",
		Description:   "A wildcard MUST be accompanied by other data to its right (Only checks DNSName)",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &brIANBareWildcard{},
	})
}
