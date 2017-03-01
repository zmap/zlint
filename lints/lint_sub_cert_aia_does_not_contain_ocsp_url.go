// lint_sub_cert_aia_does_not_contain_ocsp_url.go
/**************************************************************************************************
CAB: 7.1.2.3
authorityInformationAccess
With the exception of stapling, which is noted below, this extension MUST be present. It MUST NOT be
marked critical, and it MUST contain the HTTP URL of the Issuing CA’s OCSP responder (accessMethod
= 1.3.6.1.5.5.7.48.1). It SHOULD also contain the HTTP URL of the Issuing CA’s certificate
(accessMethod = 1.3.6.1.5.5.7.48.2). See Section 13.2.1 for details.
***************************************************************************************************/

package lints

import (
	"github.com/zmap/zgrab/ztools/x509"
	"github.com/zmap/zlint/util"
)

type subCertOcspUrl struct {
	// Internal data here
}

func (l *subCertOcspUrl) Initialize() error {
	return nil
}

func (l *subCertOcspUrl) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return !util.IsCaCert(c)
}

func (l *subCertOcspUrl) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.OCSPServer != nil {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "sub_cert_aia_does_not_contain_ocsp_url",
		Description:   "Subscriber certificates authorityInformationAccess extension must contain the HTTP URL of the Issuing CA’s OCSP responder",
		Providence:    "CAB: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertOcspUrl{}})
}
