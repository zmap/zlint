package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameValidTLD struct {
	// Internal data here
}

func (l *DNSNameValidTLD) Initialize() error {
	return nil
}

func (l *DNSNameValidTLD) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *DNSNameValidTLD) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName != "" {
		if !util.HasValidTLD(c.Subject.CommonName) {
			return ResultStruct{Result: Error}, nil
		}
	}
	for _, dns := range c.DNSNames {
		if !util.HasValidTLD(dns) {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_not_valid_tld",
		Description:   "DNSNames must have a valid TLD.",
		Source:        "RFC 5280",
		EffectiveDate: util.RFC5280Date,
		Test:          &DNSNameValidTLD{},
	})
}
