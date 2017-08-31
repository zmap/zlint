// lint_ian_iana_pub_suffix_empty.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type IANPubSuffix struct{}

func (l *IANPubSuffix) Initialize() error {
	return nil
}

func (l *IANPubSuffix) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANPubSuffix) Execute(c *x509.Certificate) * LintResult{
	for _, dns := range c.IANDNSNames {
		if len(strings.Split(dns, ".")) < 3 {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ian_iana_pub_suffix_empty",
		Description:   "Domain SHOULD NOT have a bare public suffix",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &IANPubSuffix{},
	})
}
