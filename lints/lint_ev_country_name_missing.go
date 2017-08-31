// lint_ev_country_name_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evCountryMissing struct{}

func (l *evCountryMissing) Initialize() error {
	return nil
}

func (l *evCountryMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
}

func (l *evCountryMissing) Execute(c *x509.Certificate) ResultStruct {
	if util.TypeInName(&c.Subject, util.CountryNameOID) {
		return ResultStruct{Result: Pass}
	} else {
		return ResultStruct{Result: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_country_name_missing",
		Description:   "EV certificates must include countryName in subject",
		Source:        "CAB 7.1.6.1",
		EffectiveDate: util.ZeroDate,
		Lint:          &evCountryMissing{},
	})
}
