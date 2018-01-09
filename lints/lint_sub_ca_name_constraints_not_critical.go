/************************************************
Change this to match source TEXT
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubCANameConstraintsNotCritical struct{}

func (l *SubCANameConstraintsNotCritical) Initialize() error {
	return nil
}

func (l *SubCANameConstraintsNotCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsSubCA(cert) && util.IsExtInCert(cert, util.NameConstOID)
}

func (l *SubCANameConstraintsNotCritical) Execute(cert *x509.Certificate) *LintResult {
	if ski := util.GetExtFromCert(cert, util.NameConstOID); ski.Critical {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Warn}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_name_constraints_not_critical",
		Description:   "Subordinate CA Certificate: NameConstraints if present, SHOULD be marked critical.",
		Citation:      "BRs: 7.1.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABV102Date,
		Lint:          &SubCANameConstraintsNotCritical{},
	})
}
