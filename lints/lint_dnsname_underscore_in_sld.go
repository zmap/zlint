package lints

import (
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameUnderscoreInSLD struct{}

func (l *DNSNameUnderscoreInSLD) Initialize() error {
	return nil
}

func (l *DNSNameUnderscoreInSLD) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func underscoreInSLD(domain string) (bool, error) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, err
	}
	if strings.Contains(domainName.SLD, "_") {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *DNSNameUnderscoreInSLD) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		underscoreFound, err := underscoreInSLD(c.Subject.CommonName)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if underscoreFound {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		underscoreFound, err := underscoreInSLD(dns)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if underscoreFound {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_underscore_in_sld",
		Description:   "DNSName should not have underscore in SLD",
		Source:        "RFC 5280",
		Type:          RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &DNSNameUnderscoreInSLD{},
	})
}
