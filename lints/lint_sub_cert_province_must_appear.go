package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertProvinceMustAppear struct{}

func (l *subCertProvinceMustAppear) Initialize() error {
	return nil
}

func (l *subCertProvinceMustAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertProvinceMustAppear) Execute(c *x509.Certificate) *LintResult {
	if len(c.Subject.Organization) > 0 || len(c.Subject.GivenName) > 0 || len(c.Subject.Surname) > 0 {
		if len(c.Subject.Locality) == 0 {
			if len(c.Subject.Province) == 0 {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_province_must_appear",
		Description:   "Subscriber Certificate: subject:stateOrProvinceName MUST appear if the subject:organizationName, subject:givenName, or subject:surname fields are present and subject:localityName is absent.",
		Citation:      "BRs: 7.1.4.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABGivenNameDate,
		Lint:          &subCertProvinceMustAppear{},
	})
}
