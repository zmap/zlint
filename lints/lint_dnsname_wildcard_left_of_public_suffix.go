package lints

import (
	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameWildcardLeftofPublicSuffix struct {
	// Internal data here
}

func (l *DNSNameWildcardLeftofPublicSuffix) Initialize() error {
	return nil
}

func (l *DNSNameWildcardLeftofPublicSuffix) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func wildcardLeftOfPublicSuffix(domain string) (bool, error) {
	parsedDomain, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, err
	}
	if parsedDomain.SLD == "*" {
		return true, nil
	}
	return false, nil
}

func (l *DNSNameWildcardLeftofPublicSuffix) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName != "" {
		wildcardFound, err := wildcardLeftOfPublicSuffix(c.Subject.CommonName)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
		if wildcardFound {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, dns := range c.DNSNames {
		wildcardFound, err := wildcardLeftOfPublicSuffix(dns)
		if err != nil {
			return ResultStruct{Result: Fatal}, nil
		}
		if wildcardFound {
			return ResultStruct{Result: Warn}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_dnsname_wildcard_left_of_public_suffix",
		Description:   "the CA MUST establish and follow a documented procedure[^pubsuffix] that determines if the wildcard character occurs in the first label position to the left of a “registry‐controlled” label or “public suffix”",
		Source:        "BRs: 3.2.2.6",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &DNSNameWildcardLeftofPublicSuffix{},
	})
}
