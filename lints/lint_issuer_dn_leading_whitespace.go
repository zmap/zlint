// lint_issuer_dn_leading_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IssuerDNLeadingSpace struct {
	// Internal data here
}

func (l *IssuerDNLeadingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNLeadingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNLeadingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	hasSpace, err := util.DNSAttributeHasSpace(c, true)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	if hasSpace&1 != 0 {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_leading_whitespace",
		Description:   "AttributeValue in issuer RelativeDistinguishedName sequence should not have leading whitespace",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerDNLeadingSpace{}})
}
