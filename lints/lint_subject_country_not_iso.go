// lint_subject_country_not_iso.go
/**************************************************************************************************************
BRs: 7.1.4.2.2
Certificate Field: issuer:countryName (OID 2.5.4.6)
Required/Optional: Required
Contents: This field MUST contain the two-letter ISO 3166-1 country code for the country in which the issuer’s
place of business is located.
**************************************************************************************************************/

package lints

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type countryNotIso struct {
	// Internal data here
}

func (l *countryNotIso) Initialize() error {
	return nil
}

func (l *countryNotIso) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *countryNotIso) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, j := range c.Subject.Country {
		if !util.IsISOCountryCode(strings.ToUpper(j)) {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_country_not_iso",
		Description:   "The country name field MUST contain the two-letter ISO code for the country or XX",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &countryNotIso{},
	})
}
