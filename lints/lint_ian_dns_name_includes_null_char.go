// lint_IAN_dns_name_includes_null_char.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IANDnsNull struct {
	// Internal data here
}

func (l *IANDnsNull) Initialize() error {
	return nil
}

func (l *IANDnsNull) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *IANDnsNull) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.IANDNSNames {
		for i := 0; i < len(dns); i++ {
			if dns[i] == 0 {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_IAN_dns_name_includes_null_char",
		Description:   "DNSNames MUST NOT include a null character ",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &IANDnsNull{}})
}
