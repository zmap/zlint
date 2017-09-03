package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertCountryNameMustAppear struct{}

func (l *subCertCountryNameMustAppear) Initialize() error {
	return nil
}

func (l *subCertCountryNameMustAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertCountryNameMustAppear) Execute(c *x509.Certificate) *LintResult {
	if len(c.Subject.Organization) > 0 || len(c.Subject.GivenName) > 0 || len(c.Subject.Surname) > 0 {
		if len(c.Subject.Country) == 0 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_country_name_must_appear",
		Description:   "Subscriber Certificate: subject:countryName MUST appear if the subject:organizationName field, subject:givenName field, or subject:surname fields are present.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABGivenNameDate,
		Lint:          &subCertCountryNameMustAppear{},
	})
}
