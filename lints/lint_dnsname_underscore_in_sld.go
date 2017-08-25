package lints

import (
	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameUnderscoreInSLD struct {
	// Internal data here
}

func (l *DNSNameUnderscoreInSLD) Initialize() error {
	return nil
}

func (l *DNSNameUnderscoreInSLD) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func underscoreInSLD(domain string) bool {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return false
	}
	if strings.Contains(domainName.SLD, "_") {
		return true
	} else {
		return false
	}
}

func (l *DNSNameUnderscoreInSLD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	result := ResultStruct{Result: Pass}
	if underscoreInSLD(c.Subject.CommonName) {
		result = ResultStruct{Result: Error}
	}
	for _, dns := range c.DNSNames {
		if underscoreInSLD(dns) {
			result = ResultStruct{Result: Error}
		}
	}
	return result, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_underscore_in_sld",
		Description:   "DNSName should not have underscore in SLD",
		Provenance:    "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameUnderscoreInSLD{},
	})
}
