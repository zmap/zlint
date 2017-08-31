// lint_br_san_wildcard_not_first.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANWildCardFirst struct{}

func (l *SANWildCardFirst) Initialize() error {
	return nil
}

func (l *SANWildCardFirst) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANWildCardFirst) Execute(c *x509.Certificate) * LintResult{
	for _, dns := range c.DNSNames {
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
		Name:          "e_san_wildcard_not_first",
		Description:   "A wildcard MUST be in the first label of FQDN (ie not: www.*.com) (Only checks DNSName)",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &SANWildCardFirst{},
	})
}
