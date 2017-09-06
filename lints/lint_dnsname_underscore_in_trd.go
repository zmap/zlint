package lints

import (
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameUnderscoreInTRD struct{}

func (l *DNSNameUnderscoreInTRD) Initialize() error {
	return nil
}

func (l *DNSNameUnderscoreInTRD) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func underscoreInTRD(domain string) (bool, error) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, err
	}
	if strings.Contains(domainName.TRD, "_") {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *DNSNameUnderscoreInTRD) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		underscoreFound, err := underscoreInTRD(c.Subject.CommonName)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if underscoreFound {
			return &LintResult{Status: Warn}
		}
	}
	for _, dns := range c.DNSNames {
		underscoreFound, err := underscoreInTRD(dns)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if underscoreFound {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "w_dnsname_underscore_in_trd",
		Description:    "DNSName should not have an underscore in labels left of the ETLD+1",
		ReadableSource: "RFC 5280",
		Source:         RFC5280,
		EffectiveDate:  util.RFC5280Date,
		Lint:           &DNSNameUnderscoreInTRD{},
	})
}
