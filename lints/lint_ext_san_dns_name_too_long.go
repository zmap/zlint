package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANDNSTooLong struct{}

func (l *SANDNSTooLong) Initialize() error {
	return nil
}

func (l *SANDNSTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID) && len(c.DNSNames) > 0
}

func (l *SANDNSTooLong) Execute(c *x509.Certificate) *LintResult {
	for _, dns := range c.DNSNames {
		if len(dns) > 253 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_dns_name_too_long",
		Description:   "DNSName must be less than or equal to 253 bytes",
		Citation:      "RFC 5280",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &SANDNSTooLong{},
	})
}
