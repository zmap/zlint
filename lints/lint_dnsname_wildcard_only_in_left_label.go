package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameWildcardOnlyInLeftlabel struct {
	// Internal data here
}

func (l *DNSNameWildcardOnlyInLeftlabel) Initialize() error {
	return nil
}

func (l *DNSNameWildcardOnlyInLeftlabel) CheckApplies(c *x509.Certificate) bool {
	return true
}

func wildcardNotInLeftLabel(domain string) bool {
	labels := strings.Split(domain, ".")
	labels = labels[1:]
	for _, label := range labels {
		if strings.Contains(label, "*") {
			return true
		}
	}
	return false
}

func (l *DNSNameWildcardOnlyInLeftlabel) RunTest(c *x509.Certificate) (ResultStruct, error) {
	result := ResultStruct{Result: Pass}
	if wildcardNotInLeftLabel(c.Subject.CommonName) {
		result = ResultStruct{Result: Error}
	}
	for _, dns := range c.DNSNames {
		if wildcardNotInLeftLabel(dns) {
			result = ResultStruct{Result: Error}
		}
	}
	return result, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_wildcard_only_in_left_label",
		Description:   "DNSName should not have wildcards except in the left-most label",
		Provenance:    "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameWildcardOnlyInLeftlabel{},
	})
}
