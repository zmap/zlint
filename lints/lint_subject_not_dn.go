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

type subjectDN struct {
	// Internal data here
}

func (l *subjectDN) Initialize() error {
	return nil
}

func (l *subjectDN) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectDN) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if reflect.TypeOf(c.Subject) != reflect.TypeOf(*(new(pkix.Name))) {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_not_dn",
		Description:   "When not empty, the subject field MUST be a distinguished name",
		Source:        "RFC 5280: 4.1.2.6",
		EffectiveDate: util.RFC2459Date,
		Test:          &subjectDN{},
	})
}
