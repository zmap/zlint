// lint_ca_country_name_invalid.go
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

type caCountryNameInvalid struct {
	// Internal data here
}

func (l *caCountryNameInvalid) Initialize() error {
	return nil
}

func (l *caCountryNameInvalid) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.IsCA
}

func (l *caCountryNameInvalid) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.Country != nil {
		for _, j := range c.Subject.Country {
			if !util.IsISOCountryCode(j) {
				return ResultStruct{Result: Error}, nil
			}
		}
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: NA}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_country_name_invalid",
		Description:   "Root and Subordinate CA certificates MUST have a two-letter country code specified in ISO 3166-1",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &caCountryNameInvalid{},
	})
}
