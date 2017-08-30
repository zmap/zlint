// lint_sub_cert_aia_does_not_contain_issuing_ca_url.go
/************************************************************************
BRs: 7.1.2.3
cRLDistributionPoints
This extension MAY be present. If present, it MUST NOT be marked critical, and it MUST contain the
HTTP URL of the CA’s CRL service. See Section 13.2.1 for details.
*************************************************************************/

package lints

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertIssuerUrl struct {
	// Internal data here
}

func (l *subCertIssuerUrl) Initialize() error {
	return nil
}

func (l *subCertIssuerUrl) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubscriberCert(c)
}

func (l *subCertIssuerUrl) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, url := range c.IssuingCertificateURL {
		if strings.HasPrefix(url, "http://") {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Warn}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_cert_aia_does_not_contain_issuing_ca_url",
		Description:   "Subscriber certificates authorityInformationAccess extension should contain the HTTP URL of the issuing CA’s certificate",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertIssuerUrl{},
	})
}
