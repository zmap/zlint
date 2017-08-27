package lints

import (
	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameUnderscoreInTRD struct {
	// Internal data here
}

func (l *DNSNameUnderscoreInTRD) Initialize() error {
	return nil
}

func (l *DNSNameUnderscoreInTRD) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func underscoreInTRD(domain string) bool {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return false
	}
	if strings.Contains(domainName.TRD, "_") {
		return true
	} else {
		return false
	}
}

func (l *DNSNameUnderscoreInTRD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	result := ResultStruct{Result: Pass}
	if underscoreInTRD(c.Subject.CommonName) {
		result = ResultStruct{Result: Warn}
	}
	for _, dns := range c.DNSNames {
		if underscoreInTRD(dns) {
			result = ResultStruct{Result: Warn}
		}
	}
	return result, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_dnsname_underscore_in_trd",
		Description:   "DNSName should not have an underscore in labels left of the ETLD+1",
		Provenance:    "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameUnderscoreInTRD{},
	})
}
