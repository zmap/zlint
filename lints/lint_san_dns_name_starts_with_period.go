// lint_SAN_dns_name_starts_with_period.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type SANDnsPeriod struct {
	// Internal data here
}

func (l *SANDnsPeriod) Initialize() error {
	return nil
}

func (l *SANDnsPeriod) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SanOID)
}

func (l *SANDnsPeriod) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		if strings.HasPrefix(dns, ".") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_SAN_dns_name_starts_with_period",
		Description:   "DNSName MUST NOT start with a period",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &SANDnsPeriod{}})
}
