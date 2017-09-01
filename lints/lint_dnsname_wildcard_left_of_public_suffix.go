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

func wildcardLeftOfPublicSuffix(domain string) (bool, error) {
	parsedDomain, err := util.ICANNPublicSuffixParse(domain)
	if err != nil {
		return true, err
	}
	if parsedDomain.SLD == "*" {
		return true, nil
	}
	return false, nil
}

func (l *DNSNameWildcardLeftofPublicSuffix) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		wildcardFound, err := wildcardLeftOfPublicSuffix(c.Subject.CommonName)
		if err != nil {
			return &LintResult{Status: NA}
		}
		if wildcardFound {
			return &LintResult{Status: Warn}
		}
	}
	for _, dns := range c.DNSNames {
		wildcardFound, err := wildcardLeftOfPublicSuffix(dns)
		if err != nil {
			return &LintResult{Status: NA}
		}
		if wildcardFound {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_dnsname_wildcard_left_of_public_suffix",
		Description:   "the CA MUST establish and follow a documented procedure[^pubsuffix] that determines if the wildcard character occurs in the first label position to the left of a “registry‐controlled” label or “public suffix”",
		Source:        "BRs: 3.2.2.6",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &DNSNameWildcardLeftofPublicSuffix{},
	})
}
