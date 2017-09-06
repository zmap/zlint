// lint_invalid_certificate_version.go
/************************************************
Certificates MUST be of type X.509 v3.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type InvalidCertificateVersion struct{}

func (l *InvalidCertificateVersion) Initialize() error {
	return nil
}

func (l *InvalidCertificateVersion) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *InvalidCertificateVersion) Execute(cert *x509.Certificate) *LintResult {
	if cert.Version != 3 {
		return &LintResult{Status: Error}
	}
	//else
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_invalid_certificate_version",
		Description:    "Certificates MUST be of type X.590 v3",
		ReadableSource: "BRs: 7.1.1",
		Source:         CABFBaselineRequirements,
		EffectiveDate:  util.CABV130Date,
		Lint:           &InvalidCertificateVersion{},
	})
}
