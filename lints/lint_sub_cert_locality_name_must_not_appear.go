package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertLocalityNameMustNotAppear struct{}

func (l *subCertLocalityNameMustNotAppear) Initialize() error {
	return nil
}

func (l *subCertLocalityNameMustNotAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertLocalityNameMustNotAppear) Execute(c *x509.Certificate) *LintResult {
	if len(c.Subject.Organization) == 0 && len(c.Subject.GivenName) == 0 && len(c.Subject.Surname) == 0 {
		if len(c.Subject.Locality) > 0 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_locality_name_must_not_appear",
		Description:   "Subscriber Certificate: subject:localityName MUST NOT appear if subject:organizationName, subject:givenName, and subject:surname fields are absent.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABGivenNameDate,
		Lint:          &subCertLocalityNameMustNotAppear{},
	})
}
