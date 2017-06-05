// lint_ub_state_name_invalid.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-state-name INTEGER ::= 128
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ubStateNameInvalid struct {
	// Internal data here
}

func (l *ubStateNameInvalid) Initialize() error {
	return nil
}

func (l *ubStateNameInvalid) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *ubStateNameInvalid) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, j := range c.Subject.Province {
		if len(j) > 128 {
			return ResultStruct{Result: Error}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ub_state_name_invalid",
		Description:   "The 'State Name' field must be less than 128 integers long.",
		Providence:    "RFC 5280: A.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &ubStateNameInvalid{}})
}
