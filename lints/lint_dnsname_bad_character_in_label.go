package lints

import (
	"regexp"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameProperCharacters struct {
	CompiledExpression *regexp.Regexp
}

func (l *DNSNameProperCharacters) Initialize() error {
	const dnsNameRegexp = `^(\*\.)?(\?\.)?(A-Za-z0-9*_-]+\.)*[A-Za-z0-9*_-]*$`
	var err error
	l.CompiledExpression, err = regexp.Compile(dnsNameRegexp)

	return err
}

func (l *DNSNameProperCharacters) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *DNSNameProperCharacters) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		if !l.CompiledExpression.MatchString(c.Subject.CommonName) {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		if !l.CompiledExpression.MatchString(dns) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_bad_character_in_label",
		Description:   "Characters in labels of DNSNames MUST be alphanumeric, - , _ or *",
		Citation:      "BRs: 7.1.4.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &DNSNameProperCharacters{},
	})
}
