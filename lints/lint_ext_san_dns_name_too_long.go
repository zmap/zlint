// lint_ext_san_dns_name_too_long.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANDNSTooLong struct {
	// Internal data here
}

func (l *SANDNSTooLong) Initialize() error {
	return nil
}

func (l *SANDNSTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SANOID) && len(c.DNSNames) > 0
}

func (l *SANDNSTooLong) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		if len(dns) > 253 {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_dns_name_underscore",
		Description:   "DNSName must be less than 253 bytes",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SANDNSTooLong{}})
}
