package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameLeftLabelWildcardCheck struct {
	// Internal data here
}

func (l *DNSNameLeftLabelWildcardCheck) Initialize() error {
	return nil
}

func (l *DNSNameLeftLabelWildcardCheck) CheckApplies(c *x509.Certificate) bool {
	return true
}

func wildcardInLeftLabelInorrect(domain string) bool {
	labels := strings.Split(domain, ".")
	if len(labels) > 1 {
		for _, label := range labels {
			if strings.Contains(label, "*") {
				if label != "*" {
					return true
				}
			}
		}
	}
	return false
}

func (l *DNSNameLeftLabelWildcardCheck) RunTest(c *x509.Certificate) (ResultStruct, error) {
	result := ResultStruct{Result: Pass}
	if wildcardInLeftLabelInorrect(c.Subject.CommonName) {
		result = ResultStruct{Result: Error}
	}
	for _, dns := range c.DNSNames {
		if wildcardInLeftLabelInorrect(dns) {
			result = ResultStruct{Result: Error}
		}
	}
	return result, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_left_label_wildcard_correct",
		Description:   "Wildcards in the left label of DNSName should only be *",
		Provenance:    "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameLeftLabelWildcardCheck{},
	})
}
