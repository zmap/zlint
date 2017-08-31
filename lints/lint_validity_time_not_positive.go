// lint_validity_time_not_positive.go
/************************************************
Change this to match provenance TEXT
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type validityNegative struct {
	// Internal data here
}

func (l *validityNegative) Initialize() error {
	return nil
}

func (l *validityNegative) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *validityNegative) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.NotBefore.After(c.NotAfter) {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_validity_time_not_positive",
		Description:   "Certificates MUST have a positive time for which they are valid",
		Source:        "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &validityNegative{},
	})
}
