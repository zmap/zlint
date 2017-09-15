package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameValidTLD struct{}

func (l *DNSNameValidTLD) Initialize() error {
	return nil
}

func (l *DNSNameValidTLD) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *DNSNameValidTLD) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		if !util.HasValidTLD(c.Subject.CommonName) {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		if !util.HasValidTLD(dns) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_not_valid_tld",
		Description:   "DNSNames must have a valid TLD.",
		Citation:      "BRs: 7.1.4.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &DNSNameValidTLD{},
	})
}
