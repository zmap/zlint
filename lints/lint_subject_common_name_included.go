// lint_subject_common_name_included.go
/***************************************************************
CAB: 7.1.4.2.2
Required/Optional: Deprecated (Discouraged, but not prohibited)
***************************************************************/
package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type commonNames struct {
	// Internal data here
}

func (l *commonNames) Initialize() error {
	return nil
}

func (l *commonNames) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *commonNames) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName == "" {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "subject_common_name_included",
		Description:   "Use of the common name field is discouraged.",
		Providence:    "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &commonNames{}})
}
