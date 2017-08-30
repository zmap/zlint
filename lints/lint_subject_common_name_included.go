// lint_subject_common_name_included.go
/***************************************************************
BRs: 7.1.4.2.2
Required/Optional: Deprecated (Discouraged, but not prohibited)
***************************************************************/
package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type commonNames struct {
	// Internal data here
}

func (l *commonNames) Initialize() error {
	return nil
}

func (l *commonNames) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return !util.IsCACert(c)
}

func (l *commonNames) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.CommonName == "" {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Notice}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_subject_common_name_included",
		Description:   "Subscriber Certificate: commonName is deprecated.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &commonNames{},
	})
}
