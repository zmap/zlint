// lint_br_san_bare_wildcard.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type brSANBareWildcard struct{}

func (l *brSANBareWildcard) Initialize() error {
	return nil
}

func (l *brSANBareWildcard) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *brSANBareWildcard) Execute(c *x509.Certificate) *LintResult {
	for _, dns := range c.DNSNames {
		if strings.HasSuffix(dns, "*") {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_san_bare_wildcard",
		Description:   "A wildcard MUST be accompanied by other data to its right (Only checks DNSName)",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &brSANBareWildcard{},
	})
}
