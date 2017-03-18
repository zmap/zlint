// lint_ext_san_missing.go
/************************************************
CAB: 7.1.4.2.1
Subject Alternative Name Extension
Certificate Field: extensions:subjectAltName
Required/Optional: Required
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANMissing struct {
	// Internal data here
}

func (l *SANMissing) Initialize() error {
	return nil
}

func (l *SANMissing) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SANMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(c, util.SanOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_missing",
		Description:   "Certificates must contain the Subject Alternate Name extension.",
		Providence:    "CAB: 7.1.4.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &SANMissing{}})
}
