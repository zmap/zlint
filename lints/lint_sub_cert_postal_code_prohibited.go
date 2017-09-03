package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertPostalCodeMustNotAppear struct{}

func (l *subCertPostalCodeMustNotAppear) Initialize() error {
	return nil
}

func (l *subCertPostalCodeMustNotAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertPostalCodeMustNotAppear) Execute(c *x509.Certificate) *LintResult {
	// BR 7.1.4.2.2 uses "or" and "and" interchangeably when they mean "and".
	if len(c.Subject.Organization) == 0 && len(c.Subject.GivenName) == 0 && len(c.Subject.Surname) == 0 {
		if len(c.Subject.PostalCode) > 0 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_postal_code_must_not_appear",
		Description:   "Subscriber Certificate: subject:postalCode MUST NOT appear if the subject:organizationName field, subject:givenName field, or subject:surname fields are absent.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABGivenNameDate,
		Lint:          &subCertPostalCodeMustNotAppear{},
	})
}
