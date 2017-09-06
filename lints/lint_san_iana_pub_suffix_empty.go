// lint_san_iana_pub_suffix_empty.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/net/publicsuffix"
)

type pubSuffix struct{}

func (l *pubSuffix) Initialize() error {
	return nil
}

func (l *pubSuffix) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *pubSuffix) Execute(c *x509.Certificate) *LintResult {
	for _, dns := range c.DNSNames {
		suffix, _ := publicsuffix.PublicSuffix(dns)
		if suffix == dns {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "w_san_iana_pub_suffix_empty",
		Description:    "The domain SHOULD NOT have a bare public suffix",
		ReadableSource: "awslabs certlint",
		Source:         AWSLabs,
		EffectiveDate:  util.ZeroDate,
		Lint:           &pubSuffix{},
	})
}
