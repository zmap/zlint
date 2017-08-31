// lint_invalid_certificate_version.go
/************************************************
Certificates MUST be of type X.509 v3.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type InvalidCertificateVersion struct {
	// Internal data here
}

func (l *InvalidCertificateVersion) Initialize() error {
	return nil
}

func (l *InvalidCertificateVersion) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *InvalidCertificateVersion) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if cert.Version != 3 {
		return ResultStruct{Result: Error}, nil
	}
	//else
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_invalid_certificate_version",
		Description:   "Certificates MUST be of type X.590 v3",
		Source:        "BRs: 7.1.1",
		EffectiveDate: util.CABV130Date,
		Test:          &InvalidCertificateVersion{},
	})
}
