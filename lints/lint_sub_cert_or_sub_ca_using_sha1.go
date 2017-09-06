// lint_sub_cert_or_sub_ca_using_sha1.go
/**************************************************************************************************
BRs: 7.1.3
SHA‐1	MAY	be	used	with	RSA	keys	in	accordance	with	the	criteria	defined	in	Section	7.1.3.
**************************************************************************************************/
package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type sigAlgTestsSHA1 struct{}

func (l *sigAlgTestsSHA1) Initialize() error {
	return nil
}

func (l *sigAlgTestsSHA1) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *sigAlgTestsSHA1) Execute(c *x509.Certificate) *LintResult {
	if c.SignatureAlgorithm == x509.SHA1WithRSA || c.SignatureAlgorithm == x509.DSAWithSHA1 || c.SignatureAlgorithm == x509.ECDSAWithSHA1 {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_or_sub_ca_using_sha1",
		Description:   "CAs MUST NOT issue any new Subscriber certificates or Subordinate CA certificates using SHA-1 after 1 January 2016",
		Citation:      "BRs: 7.1.3",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.NO_SHA1,
		Lint:          &sigAlgTestsSHA1{},
	})
}
