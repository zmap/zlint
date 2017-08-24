package lints


import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertLocalityNameMustNotAppear struct {
	// Internal data here
}

func (l *subCertLocalityNameMustNotAppear) Initialize() error {
	return nil
}

func (l *subCertLocalityNameMustNotAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertLocalityNameMustNotAppear) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.Subject.GivenName != "" && len(c.Subject.Organization) > 0 && c.Subject.Surname != "" {
		if len(c.Subject.Locality) > 0 {
			return ResultStruct{Result: Error}, nil
		} else {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: NA}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_locality_name_must_not_appear",
		Description:   "Subscriber Certificate: subject:localityName MUST NOT appear is subject:organizationName, subject:givenName, and subject:surname fields are present." ,
		Provenance:    "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertLocalityNameMustNotAppear{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESubCertLocalityNameMustNotAppear = result },
	})
}

