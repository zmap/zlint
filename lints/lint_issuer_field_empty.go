// lint_issuer_field_empty.go
/************************************************
RFC 5280: 4.1.2.4
The issuer field identifies the entity that has signed and issued the
   certificate.  The issuer field MUST contain a non-empty distinguished
   name (DN).  The issuer field is defined as the X.501 type Name
   [X.501].
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type issuerFieldEmpty struct{}

func (l *issuerFieldEmpty) Initialize() error {
	return nil
}

func (l *issuerFieldEmpty) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *issuerFieldEmpty) Execute(c *x509.Certificate) *LintResult {
	if &c.Issuer != nil && util.NotAllNameFieldsAreEmpty(&c.Issuer) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_issuer_field_empty",
		Description:   "Certificate issuer field MUST NOT be empty and must have a non-empty distingushed name",
		Source:        "RFC 5280: 4.1.2.4",
		EffectiveDate: util.RFC2459Date,
		Lint:          &issuerFieldEmpty{},
	})
}
