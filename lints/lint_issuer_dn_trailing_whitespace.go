// lint_issuer_dn_trailing_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IssuerDNTrailingSpace struct{}

func (l *IssuerDNTrailingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNTrailingSpace) Execute(c *x509.Certificate) * LintResult{
	_, trailing, err := util.CheckRDNSequenceWhiteSpace(c.RawIssuer)
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	if trailing {
		return &LintResult{Status: Warn}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_trailing_whitespace",
		Description:   "AttributeValue in issuer RelativeDistinguishedName sequence SHOULD NOT have trailing whitespace",
		Source:        "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Lint:          &IssuerDNTrailingSpace{},
	})
}
