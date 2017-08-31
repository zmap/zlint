package lints

import (
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameHyphenInSLD struct{}

func (l *DNSNameHyphenInSLD) Initialize() error {
	return nil
}

func (l *DNSNameHyphenInSLD) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func hyphenAtStartOrEndOfSLD(domain string) (bool, error) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, err
	}
	if strings.HasPrefix(domainName.SLD, "-") || strings.HasSuffix(domainName.SLD, "-") {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *DNSNameHyphenInSLD) Execute(c *x509.Certificate) LintResult {
	if c.Subject.CommonName != "" {
		hyphenFound, err := hyphenAtStartOrEndOfSLD(c.Subject.CommonName)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if hyphenFound {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		hyphenFound, err := hyphenAtStartOrEndOfSLD(dns)
		if err != nil {
			return &LintResult{Status: Fatal}
		}
		if hyphenFound {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_hyphen_in_sld",
		Description:   "DNSName should not have a hyphen beginning or ending the SLD",
		Source:        "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Lint:          &DNSNameHyphenInSLD{},
	})
}
