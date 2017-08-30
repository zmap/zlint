// lint_subject_organizational_unit_name_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-organizational-unit-name INTEGER ::= 64
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectOrganizationalUnitNameMaxLength struct {
	// Internal data here
}

func (l *subjectOrganizationalUnitNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectOrganizationalUnitNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *subjectOrganizationalUnitNameMaxLength) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, j := range c.Subject.OrganizationalUnit {
		if len(j) > 64 {
			return ResultStruct{Result: Error}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_organizational_unit_name_max_length",
		Description:   "The 'Organizational Unit Name' field of the subject MUST be less than 64 characters",
		Source:        "RFC 5280: A.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &subjectOrganizationalUnitNameMaxLength{},
	})
}
