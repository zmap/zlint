package lints

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameWildcardOnlyInLeftlabel struct{}

func (l *DNSNameWildcardOnlyInLeftlabel) Initialize() error {
	return nil
}

func (l *DNSNameWildcardOnlyInLeftlabel) CheckApplies(c *x509.Certificate) bool {
	return true
}

func wildcardNotInLeftLabel(domain string) bool {
	labels := strings.Split(domain, ".")
	if len(labels) > 1 {
		labels = labels[1:]
		for _, label := range labels {
			if strings.Contains(label, "*") {
				return true
			}
		}
	}
	return false
}

func (l *DNSNameWildcardOnlyInLeftlabel) Execute(c *x509.Certificate) *LintResult {
	if wildcardNotInLeftLabel(c.Subject.CommonName) {
		return &LintResult{Status: Error}
	}
	for _, dns := range c.DNSNames {
		if wildcardNotInLeftLabel(dns) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_wildcard_only_in_left_label",
		Description:   "DNSName should not have wildcards except in the left-most label",
		Citation:      "BRs: 7.1.4.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &DNSNameWildcardOnlyInLeftlabel{},
	})
}
