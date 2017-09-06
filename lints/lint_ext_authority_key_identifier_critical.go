// lint_ext_authority_key_identifier_critical.go
/*********************************************************
RFC 5280: 4.2.1.1
Conforming CAs MUST mark this extension as non-critical.
**********************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type authorityKeyIdCritical struct{}

func (l *authorityKeyIdCritical) Initialize() error {
	return nil
}

func (l *authorityKeyIdCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.AuthkeyOID)
}

func (l *authorityKeyIdCritical) Execute(c *x509.Certificate) *LintResult {
	aki := util.GetExtFromCert(c, util.AuthkeyOID) //pointer to the extension
	if aki.Critical {
		return &LintResult{Status: Error}
	} else { //implies !aki.Critical
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_ext_authority_key_identifier_critical",
		Description:    "The authority key identifier extension must be non-critical",
		ReadableSource: "RFC 5280: 4.2.1.1",
		Source:         RFC5280,
		EffectiveDate:  util.RFC2459Date,
		Lint:           &authorityKeyIdCritical{},
	})
}
