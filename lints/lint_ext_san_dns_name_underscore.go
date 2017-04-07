// lint_ext_san_dns_name_underscore.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type SANDNSContainsUnderscore struct {
	// Internal data here
}

func (l *SANDNSContainsUnderscore) Initialize() error {
	return nil
}

func (l *SANDNSContainsUnderscore) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SANOID)
}

func (l *SANDNSContainsUnderscore) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		if strings.Contains(dns, "_") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_san_dns_name_underscore",
		Description:   "DNS names SHOULD NOT have an underscore",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SANDNSContainsUnderscore{}})
}
