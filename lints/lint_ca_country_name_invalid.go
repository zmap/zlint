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

type caCountryNameInvalid struct{}

func (l *caCountryNameInvalid) Initialize() error {
	return nil
}

func (l *caCountryNameInvalid) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

func (l *caCountryNameInvalid) Execute(c *x509.Certificate) * LintResult{
	if c.Subject.Country != nil {
		for _, j := range c.Subject.Country {
			if !util.IsISOCountryCode(j) {
				return &LintResult{Status: Error}
			}
		}
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: NA}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_country_name_invalid",
		Description:   "Root and Subordinate CA certificates MUST have a two-letter country code specified in ISO 3166-1",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caCountryNameInvalid{},
	})
}
