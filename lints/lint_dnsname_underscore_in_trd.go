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

func underscoreInTRD(domain string) (bool, error) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, err
	}
	if strings.Contains(domainName.TRD, "_") {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *DNSNameUnderscoreInTRD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName != "" {
		underscoreFound, err := underscoreInTRD(c.Subject.CommonName)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
		if underscoreFound {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, dns := range c.DNSNames {
		underscoreFound, err := underscoreInTRD(dns)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
		if underscoreFound {
			return ResultStruct{Result: Warn}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_dnsname_underscore_in_trd",
		Description:   "DNSName should not have an underscore in labels left of the ETLD+1",
		Source:        "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameUnderscoreInTRD{},
	})
}
