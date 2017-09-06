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

func (l *IssuerDNLeadingSpace) Execute(c *x509.Certificate) *LintResult {
	leading, _, err := util.CheckRDNSequenceWhiteSpace(c.RawIssuer)
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	if leading {
		return &LintResult{Status: Warn}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "w_issuer_dn_leading_whitespace",
		Description:    "AttributeValue in issuer RelativeDistinguishedName sequence SHOULD NOT have leading whitespace",
		ReadableSource: "AWSLabs certlint",
		Source:         AWSLabs,
		EffectiveDate:  util.ZeroDate,
		Lint:           &IssuerDNLeadingSpace{},
	})
}
