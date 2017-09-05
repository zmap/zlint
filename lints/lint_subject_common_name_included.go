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

type commonNames struct{}

func (l *commonNames) Initialize() error {
	return nil
}

func (l *commonNames) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c)
}

func (l *commonNames) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName == "" {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Notice}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_subject_common_name_included",
		Description:   "Subscriber Certificate: commonName is deprecated.",
		Source:        "BRs: 7.1.4.2.2",
		Type:          CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &commonNames{},
	})
}
