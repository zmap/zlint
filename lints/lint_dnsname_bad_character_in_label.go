package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"regexp"
	"strings"
)

type DNSNameProperCharacters struct {
	// Internal data here
	CompiledExpression *regexp.Regexp
}

func (l *DNSNameProperCharacters) Initialize() error {
	dnsLabelRegex := "^[A-Za-z0-9*_-]+$"
	var err error
	l.CompiledExpression, err = regexp.Compile(dnsLabelRegex)
	if err != nil {
		return err
	}
	return nil
}

func (l *DNSNameProperCharacters) CheckApplies(c *x509.Certificate) bool {
	return true
}

func labelContainsBadCharacters(domain string, compiledExpression *regexp.Regexp) bool {
	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if !compiledExpression.MatchString(label) {
			return true
		}
	}
	return false
}

func (l *DNSNameProperCharacters) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName != "" {
		badCharacterFound := labelContainsBadCharacters(c.Subject.CommonName, l.CompiledExpression)
		if badCharacterFound {
			return ResultStruct{Result: Error}, nil
		}
	}
	for _, dns := range c.DNSNames {
		badCharacterFound := labelContainsBadCharacters(dns, l.CompiledExpression)
		if badCharacterFound {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_bad_character_in_label",
		Description:   "Characters in labels of DNSNames MUST be alphanumeric, - , _ or *",
		Provenance:    "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameProperCharacters{},
	})
}
