// lint_br_san_wildcard_not_first.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANWildCardFirst struct {
	// Internal data here
}

func (l *SANWildCardFirst) Initialize() error {
	return nil
}

func (l *SANWildCardFirst) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANWildCardFirst) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		for i := 1; i < len(dns); i++ {
			if dns[i] == '*' {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_san_wildcard_not_first",
		Description:   "A wildcard MUST be in the first label of FQDN (ie not: www.*.com) (Only checks DNSName)",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SANWildCardFirst{},
	})
}
