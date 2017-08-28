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
	return util.IsSubscriberCert(c)
}

func underscoreInSLD(domain string) (bool, ResultStruct) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, ResultStruct{Result: Fatal}
	}
	if strings.Contains(domainName.SLD, "_") {
		return true, ResultStruct{Result: Error}
	} else {
		return false, ResultStruct{Result: Pass}
	}
}

func (l *DNSNameUnderscoreInSLD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if underscoreFound, result := underscoreInSLD(c.Subject.CommonName); underscoreFound {
		return result, nil
	}
	for _, dns := range c.DNSNames {
		if underscoreFound, result := underscoreInSLD(dns); underscoreFound {
			return result, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
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
