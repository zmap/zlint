// lint_ca_country_name_missing.go
/************************************************
BRs: 7.1.2.1e
The	Certificate	Subject	MUST contain the following:
‐	countryName	(OID 2.5.4.6).	This field MUST	contain	the	two‐letter	ISO	3166‐1 country code	for	the
country	in which the CA’s place	of business	is located.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type caCountryNameMissing struct {
	// Internal data here
}

func (l *caCountryNameMissing) Initialize() error {
	return nil
}

func (l *caCountryNameMissing) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.IsCA
}

func (l *caCountryNameMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.Country != nil && c.Subject.Country[0] != "" {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_country_name_missing",
		Description:   "Root and Subordinate CA certificates MUST have a countryName present in subject information",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &caCountryNameMissing{},
	})
}
