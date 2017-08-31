package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertLocalityNameMustAppear struct {
	// Internal data here
}

func (l *subCertLocalityNameMustAppear) Initialize() error {
	return nil
}

func (l *subCertLocalityNameMustAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertLocalityNameMustAppear) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if len(c.Subject.Organization) > 0 || len(c.Subject.GivenName) > 0 || len(c.Subject.Surname) > 0 {
		if len(c.Subject.Province) == 0 {
			if len(c.Subject.Locality) == 0 {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_locality_name_must_appear",
		Description:   "Subscriber Certificate: subject:localityName MUST appear if subject:organizationName, subject:givenName, or subject:surname fields are present but the subject:stateOrProvinceName field is absent.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertLocalityNameMustAppear{},
	})
}
