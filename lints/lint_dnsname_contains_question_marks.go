package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameRedacted struct{}

func (l *DNSNameRedacted) Initialize() error {
	return nil
}

func (l *DNSNameRedacted) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func isRedactedCertificate(domain string) bool {
	domain = util.RemoveWildcardFromDomain(domain)
	return strings.HasPrefix(domain, "?.")
}

func (l *DNSNameRedacted) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		if isRedactedCertificate(c.Subject.CommonName) {
			return &LintResult{Status: Notice}
		}
	}
	for _, domain := range c.DNSNames {
		if isRedactedCertificate(domain) {
			return &LintResult{Status: Notice}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_contains_redacted_dnsname",
		Description:   "Some Precerts are prepended with question marks.",
		Source:        CABFBaselineRequirements,
		Citation:      "MDSP",
		EffectiveDate: util.ZeroDate,
		Lint:          &DNSNameRedacted{},
	})
}
