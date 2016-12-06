// lint_subject_contains_noninformational_value.go
/**********************************************************************************************************************
CAB: 7.1.4.2.2
Other Subject Attributes
With the exception of the subject:organizationalUnitName (OU) attribute, optional attributes, when present within
the subject field, MUST contain information that has been verified by the CA. Metadata such as ‘.’, ‘-‘, and ‘ ‘ (i.e.
space) characters, and/or any other indication that the value is absent, incomplete, or not applicable, SHALL NOT
be used.
**********************************************************************************************************************/
package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
	"strings"
)

type illegalChar struct {
	// Internal data here
}

func (l *illegalChar) Initialize() error {
	return nil
}

func (l *illegalChar) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *illegalChar) RunTest(c *x509.Certificate) (ResultStruct, error) {
	domain := c.Subject.DomainComponent
	serial := c.Subject.SerialNumber
	names := c.Subject.Names
	for _, j := range names {
		tempStr, ok := j.Value.(string)
		if !ok {
			continue //TODO: change this?
		}
		if tempStr == "-" || tempStr == "." || tempStr == " " {
			return ResultStruct{Result: Error}, nil
		}
	}
	if serial == "-" || serial == "." || serial == " " {
		return ResultStruct{Result: Error}, nil
	}
	for _, j := range domain {
		if strings.Compare(j, "-") == 0 ||
			strings.Compare(j, ".") == 0 ||
			strings.Compare(j, " ") == 0 {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "subject_contains_noninformational_value",
		Description:   "Subject name fields must not contain '.','-',' ' or any other indication that the field has been ommited.",
		Providence:    "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &illegalChar{}})
}
