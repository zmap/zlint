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
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func underscoreInTRD(domain string) (bool, ResultStruct) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, ResultStruct{Result: Fatal}
	}
	if strings.Contains(domainName.TRD, "_") {
		return true, ResultStruct{Result: Warn}
	} else {
		return false, ResultStruct{Result: Pass}
	}
}

func (l *DNSNameUnderscoreInTRD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if underscoreFound, result := underscoreInTRD(c.Subject.CommonName); underscoreFound {
		return result, nil
	}
	for _, dns := range c.DNSNames {
		if underscoreFound, result := underscoreInTRD(dns); underscoreFound {
			return result, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
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
