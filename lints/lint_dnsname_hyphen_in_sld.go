package lints

import (
	"strings"

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

func (l *DNSNameHyphenInSLD) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" && !util.CommonNameIsIP(c) {
		domainInfo := c.GetParsedSubjectCommonName(false)
		if domainInfo.ParseError != nil {
			return &LintResult{Status: NA}
		}
		if strings.HasPrefix(domainInfo.ParsedDomain.SLD, "-") || strings.HasSuffix(domainInfo.ParsedDomain.SLD, "-") {
			return &LintResult{Status: Error}
		}
	}
	parsedSANDNSNames := c.GetParsedDNSNames(false)
	for i := range c.GetParsedDNSNames(false) {
		if parsedSANDNSNames[i].ParseError != nil {
			return &LintResult{Status: NA}
		}
		if strings.HasPrefix(parsedSANDNSNames[i].ParsedDomain.SLD, "-") ||
			strings.HasSuffix(parsedSANDNSNames[i].ParsedDomain.SLD, "-") {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_hyphen_in_sld",
		Description:   "DNSName should not have a hyphen beginning or ending the SLD",
		Citation:      "BRs 7.1.4.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.RFC5280Date,
		Lint:          &DNSNameHyphenInSLD{},
	})
}
