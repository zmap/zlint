// lint_ev_locality_name_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evLocalityMissing struct{}

func (l *evLocalityMissing) Initialize() error {
	return nil
}

func (l *evLocalityMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
}

func (l *evLocalityMissing) Execute(c *x509.Certificate) ResultStruct {
	if util.TypeInName(&c.Subject, util.LocalityNameOID) {
		return ResultStruct{Result: Pass}
	} else {
		return ResultStruct{Result: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_locality_name_missing",
		Description:   "EV certificates must include localityName in subject",
		Source:        "CAB 7.1.6.1",
		EffectiveDate: util.ZeroDate,
		Lint:          &evLocalityMissing{},
	})
}
