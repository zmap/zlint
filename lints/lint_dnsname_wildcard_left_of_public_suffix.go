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

func wildcardLeftOfPublicSuffix(domain string) (bool, ResultStruct) {
	parsedDomain, err := publicsuffix.Parse(domain)
	if err != nil {
		return true, ResultStruct{Result: Fatal}
	}
	if parsedDomain.TRD == "" {
		if parsedDomain.SLD == "*" {
			return true, ResultStruct{Result: Warn}
		}
	}
	return false, ResultStruct{Result: Pass}
}

func (l *DNSNameWildcardLeftofPublicSuffix) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if wildcardFound, result := wildcardLeftOfPublicSuffix(c.Subject.CommonName); wildcardFound {
		return result, nil
	}
	for _, dns := range c.DNSNames {
		if wildcardFound, result := wildcardLeftOfPublicSuffix(dns); wildcardFound {
			return result, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_dnsname_wildcard_left_of_public_suffix",
		Description:   "the CA MUST establish and follow a documented procedure[^pubsuffix] that determines if the wildcard character occurs in the first label position to the left of a “registry‐controlled” label or “public suffix”",
		Provenance:    "BRs: 3.2.2.6",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &DNSNameWildcardLeftofPublicSuffix{},
	})
}
