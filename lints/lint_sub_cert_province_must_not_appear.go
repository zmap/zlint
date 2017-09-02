package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertProvinceMustNotAppear struct{}

func (l *subCertProvinceMustNotAppear) Initialize() error {
	return nil
}

func (l *subCertProvinceMustNotAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertProvinceMustNotAppear) Execute(c *x509.Certificate) *LintResult {
	if len(c.Subject.Organization) == 0 && len(c.Subject.GivenName) == 0 && len(c.Subject.Surname) == 0 {
		if len(c.Subject.Province) > 0 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_province_must_not_appear",
		Description:   "Subscriber Certificate: subject:stateOrProvinceName MUST NOT appeear if the subject:organizationName, subject:givenName, and subject:surname fields are absent.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABGivenNameDate,
		Lint:          &subCertProvinceMustNotAppear{},
	})
}
