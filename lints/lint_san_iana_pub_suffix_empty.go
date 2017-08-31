// lint_san_iana_pub_suffix_empty.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/net/publicsuffix"
)

type pubSuffix struct {
	// Internal data here
}

func (l *pubSuffix) Initialize() error {
	return nil
}

func (l *pubSuffix) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *pubSuffix) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		suffix, _ := publicsuffix.PublicSuffix(dns)
		if suffix == dns {
			return ResultStruct{Result: Warn}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_san_iana_pub_suffix_empty",
		Description:   "The domain SHOULD NOT have a bare public suffix",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &pubSuffix{},
	})
}
