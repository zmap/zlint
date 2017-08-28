package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"regexp"
	"strings"
)

type DNSNameProperCharacters struct {
	// Internal data here
}

func (l *DNSNameProperCharacters) Initialize() error {
	return nil
}

func (l *DNSNameProperCharacters) CheckApplies(c *x509.Certificate) bool {
	return true
}

func labelContainsBadCharacters(domain string) (bool, error) {
	dnsLabelRegex := "^[A-Za-z0-9*_-]+$"
	re, err := regexp.Compile(dnsLabelRegex)
	if err != nil {
		return true, err
	}
	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if !re.MatchString(label) {
			return true, nil
		}
	}
	return false, nil
}

func (l *DNSNameProperCharacters) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName != "" {
		badCharacterFound, err := labelContainsBadCharacters(c.Subject.CommonName)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
		if badCharacterFound {
			return ResultStruct{Result: Error}, nil
		}
	}
	for _, dns := range c.DNSNames {
		badCharacterFound, err := labelContainsBadCharacters(dns)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
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
