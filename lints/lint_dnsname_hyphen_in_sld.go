package lints

import (
	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type DNSNameHyphenInSLD struct {
	// Internal data here
}

func (l *DNSNameHyphenInSLD) Initialize() error {
	return nil
}

func (l *DNSNameHyphenInSLD) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func hyphenInSLD(domain string) bool {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return false
	}
	if strings.HasPrefix(domainName.SLD, "-") || strings.HasSuffix(domainName.SLD, "-") {
		return true
	} else {
		return false
	}
}

func (l *DNSNameHyphenInSLD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	result := ResultStruct{Result: Pass}
	if hyphenInSLD(c.Subject.CommonName) {
		result = ResultStruct{Result: Error}
	}
	for _, dns := range c.DNSNames {
		if hyphenInSLD(dns) {
			result = ResultStruct{Result: Error}
		}
	}
	return result, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_hyphen_in_sld",
		Description:   "DNSName should not have a hyphen beginning or ending the SLD",
		Provenance:    "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameHyphenInSLD{},
	})
}
