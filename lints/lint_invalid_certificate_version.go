// lint_invalid_certificate_version.go
/************************************************
Certificates MUST be of type X.509 v3.
************************************************/

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
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
		return ResultStruct{Result: Error, Details: "version " + string(cert.Version)}, nil
	}
	//else
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "invalid_certificate_version",
		Description:   "Certificate must be version 3 (encoded as 2)",
		Providence:    "CAB: 7.1.1",
		EffectiveDate: util.CABV130Date,
		Test:          &InvalidCertificateVersion{}})
}
