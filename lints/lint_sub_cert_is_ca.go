package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertNotCA struct {
	// Internal data here
}

func (l *subCertNotCA) Initialize() error {
	return nil
}

func (l *subCertNotCA) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.KeyUsageOID) && c.KeyUsage&x509.KeyUsageCertSign == 0
}

func (l *subCertNotCA) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.IsCA == true {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_not_is_ca",
		Description:   "Subscriber Certificate: basicContrainsts cA field MUST NOT be true.",
		Provenance:    "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertNotCA{},
	})
}
