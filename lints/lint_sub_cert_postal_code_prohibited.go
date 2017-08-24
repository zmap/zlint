package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertPostalCodeMustNotAppear struct {
	// Internal data here
}

func (l *subCertPostalCodeMustNotAppear) Initialize() error {
	return nil
}

func (l *subCertPostalCodeMustNotAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertPostalCodeMustNotAppear) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if len(c.Subject.Organization) == 0 || len(c.Subject.GivenName) == 0 || len(c.Subject.Surname) == 0 {
		if len(c.Subject.PostalCode) > 0 {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_postal_code_must_not_appear",
		Description:   "Subscriber Certificate: subject:postalCode MUST NOT appear if the subject:organizationName field, subject:givenName field, or subject:surname fields are absent.",
		Provenance:    "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertPostalCodeMustNotAppear{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESubCertPostalCodeMustNotAppear = result },
	})
}
