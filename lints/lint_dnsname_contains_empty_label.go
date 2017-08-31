package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameEmptyLabel struct{}

func (l *DNSNameEmptyLabel) Initialize() error {
	return nil
}

func (l *DNSNameEmptyLabel) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func domainHasEmptyLabel(domain string) bool {
	labels := strings.Split(domain, ".")
	for _, elem := range labels {
		if elem == "" {
			return true
		}
	}
	return false
}

func (l *DNSNameEmptyLabel) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		if domainHasEmptyLabel(c.Subject.CommonName) {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		if domainHasEmptyLabel(dns) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_empty_label",
		Description:   "DNSNames should not have an empty label.",
		Source:        "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Lint:          &DNSNameEmptyLabel{},
	})
}
