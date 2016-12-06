// lint_ian_dns_name_starts_with_period.go

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
	"strings"
)

type ianDnsPeriod struct {
	// Internal data here
}

func (l *ianDnsPeriod) Initialize() error {
	return nil
}

func (l *ianDnsPeriod) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *ianDnsPeriod) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.IANDNSNames {
		if strings.HasPrefix(dns, ".") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ian_dns_name_starts_with_period",
		Description:   "DNSName MUST NOT start with a period",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &ianDnsPeriod{}})
}
