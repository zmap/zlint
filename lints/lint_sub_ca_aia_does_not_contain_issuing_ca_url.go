// lint_sub_ca_aia_does_not_contain_issuing_ca_url.go
/***********************************************
CAB 7.1.2.2c
With the exception of stapling, which is noted below, this extension MUST be present. It MUST NOT be
marked critical, and it MUST contain the HTTP URL of the Issuing CA’s OCSP responder (accessMethod
= 1.3.6.1.5.5.7.48.1). It SHOULD also contain the HTTP URL of the Issuing CA’s certificate
(accessMethod = 1.3.6.1.5.5.7.48.2).
************************************************/

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type subCaIssuerUrl struct {
	// Internal data here
}

func (l *subCaIssuerUrl) Initialize() error {
	return nil
}

func (l *subCaIssuerUrl) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsCaCert(c) && !util.IsRootCA(c)
}

func (l *subCaIssuerUrl) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.IssuingCertificateURL != nil {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "sub_ca_aia_does_not_contain_issuing_ca_url",
		Description:   "Subordinate CA certificates authorityInformationAccess extension should contain the HTTP URL of the Issuing CA’s certificate",
		Providence:    "CAB: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCaIssuerUrl{}})
}
