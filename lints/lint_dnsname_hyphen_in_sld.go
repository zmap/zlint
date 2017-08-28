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
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func hyphenInSLD(domain string) (bool, ResultStruct) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, ResultStruct{Result: NA}
	}
	if strings.HasPrefix(domainName.SLD, "-") || strings.HasSuffix(domainName.SLD, "-") {
		return true, ResultStruct{Result: Error}
	} else {
		return false, ResultStruct{Result: Pass}
	}
}

func (l *DNSNameHyphenInSLD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if hyphenFound, result := hyphenInSLD(c.Subject.CommonName); hyphenFound {
		return result, nil
	}
	for _, dns := range c.DNSNames {
		if hyphenFound, result := hyphenInSLD(dns); hyphenFound {
			return result, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
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
