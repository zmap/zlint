// lint_sub_ca_name_constraints_not_critical.go
/************************************************
Change this to match provenance TEXT
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SubCANameConstraintsNotCritical struct {
	// Internal data here
}

func (l *SubCANameConstraintsNotCritical) Initialize() error {
	return nil
}

func (l *SubCANameConstraintsNotCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsSubCA(cert) && util.IsExtInCert(cert, util.NameConstOID)
}

func (l *SubCANameConstraintsNotCritical) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if ski := util.GetExtFromCert(cert, util.NameConstOID); ski.Critical {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_name_constraints_not_critical",
		Description:   "Subordinate CA Certificate: NameConstraints if present, SHOULD be marked critical.",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABV102Date,
		Test:          &SubCANameConstraintsNotCritical{},
	})
}
