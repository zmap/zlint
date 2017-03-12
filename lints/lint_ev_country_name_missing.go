// lint_ev_country_name_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evCountryMissing struct {
	// Internal data here
}

func (l *evCountryMissing) Initialize() error {
	return nil
}

func (l *evCountryMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsEv(c.PolicyIdentifiers)
}

func (l *evCountryMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.TypeInName(&c.Subject, util.CountryNameOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_country_name_missing",
		Description:   "EV certificates must include countryName in subject",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &evCountryMissing{}})
}
