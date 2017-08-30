// lint_br_san_bare_wildcard.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type brSANBareWildcard struct {
	// Internal data here
}

func (l *brSANBareWildcard) Initialize() error {
	return nil
}

func (l *brSANBareWildcard) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *brSANBareWildcard) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		if strings.HasSuffix(dns, "*") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_san_bare_wildcard",
		Description:   "A wildcard MUST be accompanied by other data to its right (Only checks DNSName)",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &brSANBareWildcard{},
	})
}
