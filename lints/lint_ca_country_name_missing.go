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

type caCountryNameMissing struct{}

func (l *caCountryNameMissing) Initialize() error {
	return nil
}

func (l *caCountryNameMissing) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

func (l *caCountryNameMissing) Execute(c *x509.Certificate) LintResult {
	if c.Subject.Country != nil && c.Subject.Country[0] != "" {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_country_name_missing",
		Description:   "Root and Subordinate CA certificates MUST have a countryName present in subject information",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caCountryNameMissing{},
	})
}
