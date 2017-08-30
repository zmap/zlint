// lint_san_dns_name_includes_null_char.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANDNSNull struct {
	// Internal data here
}

func (l *SANDNSNull) Initialize() error {
	return nil
}

func (l *SANDNSNull) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANDNSNull) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
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
		Name:          "e_san_dns_name_includes_null_char",
		Description:   "DNSName MUST NOT include a null character",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SANDNSNull{},
	})
}
