// lint_issuer_dn_leading_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IssuerDNLeadingSpace struct{}

func (l *IssuerDNLeadingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNLeadingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNLeadingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	leading, _, err := util.CheckRDNSequenceWhiteSpace(c.RawIssuer)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	if leading {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_leading_whitespace",
		Description:   "AttributeValue in issuer RelativeDistinguishedName sequence SHOULD NOT have leading whitespace",
		Source:        "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerDNLeadingSpace{},
	})
}
