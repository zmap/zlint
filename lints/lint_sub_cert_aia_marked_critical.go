package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertAiaMarkedCritical struct {
	// Internal data here
}

func (l *subCertAiaMarkedCritical) Initialize() error {
	return nil
}

func (l *subCertAiaMarkedCritical) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubscriberCert(c) && util.IsExtInCert(c, util.AiaOID)
}

func (l *subCertAiaMarkedCritical) RunTest(c *x509.Certificate) (ResultStruct, error) {
	e := util.GetExtFromCert(c, util.AiaOID)
	if e.Critical {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_aia_marked_critical",
		Description:   "Subscriber Certificate: authorityInformationAccess MUST NOT be marked critical",
		Source:        "CAB: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertAiaMarkedCritical{},
	})
}
