package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertStreetAddressShouldNotExist struct{}

func (l *subCertStreetAddressShouldNotExist) Initialize() error {
	return nil
}

func (l *subCertStreetAddressShouldNotExist) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertStreetAddressShouldNotExist) Execute(c *x509.Certificate) *LintResult {
	//If all fields are absent
	if len(c.Subject.Organization) == 0 && len(c.Subject.GivenName) == 0 && len(c.Subject.Surname) == 0 {
		if len(c.Subject.StreetAddress) > 0 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_street_address_should_not_exist",
		Description:   "Subscriber Certificate: subject:streetAddress MUST NOT appear if subject:organizationName, subject:givenName, and subject:surname fields are absent.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABGivenNameDate,
		Lint:          &subCertStreetAddressShouldNotExist{},
	})
}
