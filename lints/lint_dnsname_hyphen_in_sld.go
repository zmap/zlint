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

func hyphenAtStartOrEndOfSLD(domain string) (bool, error) {
	domainName, err := util.ICANNPublicSuffixParse(domain)
	if err != nil {
		return true, err
	}
	if strings.HasPrefix(domainName.SLD, "-") || strings.HasSuffix(domainName.SLD, "-") {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *DNSNameHyphenInSLD) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		hyphenFound, err := hyphenAtStartOrEndOfSLD(c.Subject.CommonName)
		if err != nil {
			return &LintResult{Status: NA}
		}
		if hyphenFound {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		hyphenFound, err := hyphenAtStartOrEndOfSLD(dns)
		if err != nil {
			return &LintResult{Status: NA}
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
		Citation:      "BRs 7.1.4.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.RFC5280Date,
		Lint:          &DNSNameHyphenInSLD{},
	})
}
