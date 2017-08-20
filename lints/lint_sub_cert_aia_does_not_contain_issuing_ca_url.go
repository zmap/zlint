// lint_sub_cert_aia_does_not_contain_issuing_ca_url.go
/************************************************************************
CAB: 7.1.2.3
cRLDistributionPoints
This extension MAY be present. If present, it MUST NOT be marked critical, and it MUST contain the
HTTP URL of the CA’s CRL service. See Section 13.2.1 for details.
*************************************************************************/

package lints

import (
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
	return !util.IsCACert(c)
}

func (l *subCertIssuerUrl) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.IssuingCertificateURL != nil {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_aia_does_not_contain_issuing_ca_url",
		Description:   "Subscriber certificates authorityInformationAccess extension should contain the HTTP URL of the issuing CA’s certificate",
		Providence:    "CAB: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertIssuerUrl{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESubCertAiaDoesNotContainIssuingCaUrl = result },
	})
}
