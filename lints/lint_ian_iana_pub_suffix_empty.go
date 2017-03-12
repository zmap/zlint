// lint_ian_iana_pub_suffix_empty.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type ianPubSuffix struct {
	// Internal data here
}

func (l *ianPubSuffix) Initialize() error {
	return nil
}

func (l *ianPubSuffix) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *ianPubSuffix) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.IANDNSNames {
		if len(strings.Split(dns, ".")) < 3 {
			return ResultStruct{Result: Warn}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ian_iana_pub_suffix_empty",
		Description:   "Domain SHOULD NOT have bare public suffix",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &ianPubSuffix{}})
}
