// lint_subject_locality_name_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-locality-name INTEGER ::= 128
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectLocalityNameMaxLength struct {
	// Internal data here
}

func (l *subjectLocalityNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectLocalityNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *subjectLocalityNameMaxLength) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, j := range c.Subject.Locality {
		if len(j) > 128 {
			return ResultStruct{Result: Error}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_locality_name_max_length",
		Description:   "The 'Locality Name' field of the subject MUST be less than 128 characters",
		Source:        "RFC 5280: A.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &subjectLocalityNameMaxLength{},
	})
}
