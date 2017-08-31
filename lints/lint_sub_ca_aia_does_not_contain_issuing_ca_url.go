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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type subCaIssuerUrl struct {
	// Internal data here
}

func (l *subCaIssuerUrl) Initialize() error {
	return nil
}

func (l *subCaIssuerUrl) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsCACert(c) && !util.IsRootCA(c)
}

func (l *subCaIssuerUrl) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, url := range c.IssuingCertificateURL {
		if strings.HasPrefix(url, "http://") {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Warn}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_aia_does_not_contain_issuing_ca_url",
		Description:   "Subordinate CA Certificate: authorityInformationAccess SHOULD also contain the HTTP URL of the Issuing CA's certificate.",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCaIssuerUrl{},
	})
}
