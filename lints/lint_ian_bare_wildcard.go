// lint_br_IAN_bare_wildcard.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type brIanBareWildcard struct {
	// Internal data here
}

func (l *brIanBareWildcard) Initialize() error {
	return nil
}

func (l *brIanBareWildcard) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *brIanBareWildcard) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.IANDNSNames {
		if strings.HasSuffix(dns, "*") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_IAN_bare_wildcard",
		Description:   "Wildcard MUST be accompanied by other data to it's right (Only checks DNSName)",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &brIanBareWildcard{}})
}
