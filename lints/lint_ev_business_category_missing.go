// lint_ev_business_category_missing.go

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type evNoBiz struct {
	// Internal data here
}

func (l *evNoBiz) Initialize() error {
	return nil
}

func (l *evNoBiz) CheckApplies(c *x509.Certificate) bool {
	return util.IsEv(c.PolicyIdentifiers)
}

func (l *evNoBiz) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.TypeInName(&c.Subject, util.BusinessOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "ev_business_category_missing",
		Description:   "EV certificates must include businessCategory in subject",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &evNoBiz{}})
}
