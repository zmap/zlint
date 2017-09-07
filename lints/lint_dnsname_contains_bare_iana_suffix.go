// lint_dnsname_contains_bare_iana_suffix.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dnsNameContainsBareIANASuffix struct{}

func (l *dnsNameContainsBareIANASuffix) Initialize() error {
	return nil
}

func (l *dnsNameContainsBareIANASuffix) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *dnsNameContainsBareIANASuffix) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		if util.IsInTLDMap(c.Subject.CommonName) {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		if util.IsInTLDMap(dns) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_contains_bare_iana_suffix",
		Description:   "DNSNames should not contain a bare IANA suffix.",
		Citation:      "BRs: 7.1.4.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &dnsNameContainsBareIANASuffix{},
	})
}
