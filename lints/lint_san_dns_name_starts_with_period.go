// lint_san_dns_name_starts_with_period.go

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
	"strings"
)

type sanDnsPeriod struct {
	// Internal data here
}

func (l *sanDnsPeriod) Initialize() error {
	return nil
}

func (l *sanDnsPeriod) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SanOID)
}

func (l *sanDnsPeriod) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		if strings.HasPrefix(dns, ".") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "san_dns_name_starts_with_period",
		Description:   "DNSName MUST NOT start with a period",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &sanDnsPeriod{}})
}
