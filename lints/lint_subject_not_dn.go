// lint_subject_not_dn.go
/*************************************************************************
 RFC 5280: 4.1.2.6
 Where it is non-empty, the subject field MUST contain an X.500
   distinguished name (DN). The DN MUST be unique for each subject
   entity certified by the one CA as defined by the issuer name field. A
   CA may issue more than one certificate with the same DN to the same
   subject entity.
*************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/util"
	"reflect"
)

type subjectDN struct{}

func (l *subjectDN) Initialize() error {
	return nil
}

func (l *subjectDN) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectDN) Execute(c *x509.Certificate) *LintResult {
	if reflect.TypeOf(c.Subject) != reflect.TypeOf(*(new(pkix.Name))) {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_not_dn",
		Description:   "When not empty, the subject field MUST be a distinguished name",
		Citation:      "RFC 5280: 4.1.2.6",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectDN{},
	})
}
