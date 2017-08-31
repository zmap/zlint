// lint_ext_subject_key_identifier_critical.go
/**********************************************************
RFC 5280: 4.2.1.2
 Conforming CAs MUST mark this extension as non-critical.
**********************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectKeyIdCritical struct {
	// Internal data here
}

func (l *subjectKeyIdCritical) Initialize() error {
	return nil
}

func (l *subjectKeyIdCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectKeyIdentityOID)
}

func (l *subjectKeyIdCritical) RunTest(c *x509.Certificate) (ResultStruct, error) {
	ski := util.GetExtFromCert(c, util.SubjectKeyIdentityOID) //pointer to the extension
	if ski.Critical {
		return ResultStruct{Result: Error}, nil
	} else { //implies !ski.Critical
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_subject_key_identifier_critical",
		Description:   "The subject key identifier extension MUST be non-critical",
		Source:        "RFC 5280: 4.2.1.2",
		EffectiveDate: util.RFC2459Date,
		Test:          &subjectKeyIdCritical{},
	})
}
