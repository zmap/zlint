// lint_ian_dns_name_starts_with_period.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type IANDNSPeriod struct {
	// Internal data here
}

func (l *IANDNSPeriod) Initialize() error {
	return nil
}

func (l *IANDNSPeriod) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANDNSPeriod) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.IANDNSNames {
		if strings.HasPrefix(dns, ".") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ian_dns_name_starts_with_period",
		Description:   "DNSName MUST NOT start with a period",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IANDNSPeriod{},
	})
}
