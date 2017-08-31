// lint_ev_business_category_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evNoBiz struct {
	// Internal data here
}

func (l *evNoBiz) Initialize() error {
	return nil
}

func (l *evNoBiz) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
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
		Name:          "e_ev_business_category_missing",
		Description:   "EV certificates must include businessCategory in subject",
		Source:        "CAB 7.1.6.1",
		EffectiveDate: util.ZeroDate,
		Test:          &evNoBiz{},
	})
}
