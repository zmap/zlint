// lint_IAN_dns_name_starts_with_period.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type IANDnsPeriod struct {
	// Internal data here
}

func (l *IANDnsPeriod) Initialize() error {
	return nil
}

func (l *IANDnsPeriod) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *IANDnsPeriod) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.IANDNSNames {
		if strings.HasPrefix(dns, ".") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_IAN_dns_name_starts_with_period",
		Description:   "DNSName MUST NOT start with a period",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &IANDnsPeriod{}})
}
