// lint_sub_cert_aia_missing.go
/**************************************************************************************************
BRs: 7.1.2.3
authorityInformationAccess
With the exception of stapling, which is noted below, this extension MUST be present. It MUST NOT be
marked critical, and it MUST contain the HTTP URL of the Issuing CA’s OCSP responder (accessMethod
= 1.3.6.1.5.5.7.48.1). It SHOULD also contain the HTTP URL of the Issuing CA’s certificate
(accessMethod = 1.3.6.1.5.5.7.48.2). See Section 13.2.1 for details.
***************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertAiaMissing struct {
	// Internal data here
}

func (l *subCertAiaMissing) Initialize() error {
	return nil
}

func (l *subCertAiaMissing) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return !util.IsCACert(c)
}

func (l *subCertAiaMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(c, util.AiaOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_aia_missing",
		Description:   "Subscriber Certiifcate: authorityInformationAccess MUST be present, with the exception of stapling.",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertAiaMissing{},
	})
}
