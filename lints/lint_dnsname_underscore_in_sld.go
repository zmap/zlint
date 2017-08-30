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
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func underscoreInSLD(domain string) (bool, error) {
	domainName, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, err
	}
	if strings.Contains(domainName.SLD, "_") {
		return true, nil
	} else {
		return false, nil
	}
}

func (l *DNSNameUnderscoreInSLD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName != "" {
		underscoreFound, err := underscoreInSLD(c.Subject.CommonName)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
		if underscoreFound {
			return ResultStruct{Result: Error}, nil
		}
	}
	for _, dns := range c.DNSNames {
		underscoreFound, err := underscoreInSLD(dns)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
		if underscoreFound {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_underscore_in_sld",
		Description:   "DNSName should not have underscore in SLD",
		Source:        "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameUnderscoreInSLD{},
	})
}
