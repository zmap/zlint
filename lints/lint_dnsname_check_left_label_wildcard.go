package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameLeftLabelWildcardCheck struct{}

func (l *DNSNameLeftLabelWildcardCheck) Initialize() error {
	return nil
}

func (l *DNSNameLeftLabelWildcardCheck) CheckApplies(c *x509.Certificate) bool {
	return true
}

func wildcardInLeftLabelIncorrect(domain string) bool {
	labels := strings.Split(domain, ".")
	if len(labels) >= 1 {
		leftLabel := labels[0]
		if strings.Contains(leftLabel, "*") && leftLabel != "*" {
			return true
		}
	}
	return false
}

func (l *DNSNameLeftLabelWildcardCheck) Execute(c *x509.Certificate) ResultStruct {
	if wildcardInLeftLabelIncorrect(c.Subject.CommonName) {
		return ResultStruct{Result: Error}
	}
	for _, dns := range c.DNSNames {
		if wildcardInLeftLabelIncorrect(dns) {
			return ResultStruct{Result: Error}
		}
	}
	return ResultStruct{Result: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_left_label_wildcard_correct",
		Description:   "Wildcards in the left label of DNSName should only be *",
		Source:        "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Lint:          &DNSNameLeftLabelWildcardCheck{},
	})
}
