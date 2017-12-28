package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameWildcardLeftofPublicSuffix struct{}

func (l *DNSNameWildcardLeftofPublicSuffix) Initialize() error {
	return nil
}

func (l *DNSNameWildcardLeftofPublicSuffix) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *DNSNameWildcardLeftofPublicSuffix) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" && !util.CommonNameIsIP(c) {
		domainInfo := c.GetParsedSubjectCommonName(false)
		if domainInfo.ParseError != nil {
			return &LintResult{Status: NA}
		}

		if domainInfo.ParsedDomain.SLD == "*" {
			return &LintResult{Status: Warn}
		}
	}

	parsedSANDNSNames := c.GetParsedDNSNames(false)
	for i := range c.GetParsedDNSNames(false) {
		if parsedSANDNSNames[i].ParseError != nil {
			return &LintResult{Status: NA}
		}

		if parsedSANDNSNames[i].ParsedDomain.SLD == "*" {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_dnsname_wildcard_left_of_public_suffix",
		Description:   "the CA MUST establish and follow a documented procedure[^pubsuffix] that determines if the wildcard character occurs in the first label position to the left of a “registry‐controlled” label or “public suffix”",
		Citation:      "BRs: 3.2.2.6",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &DNSNameWildcardLeftofPublicSuffix{},
	})
}
