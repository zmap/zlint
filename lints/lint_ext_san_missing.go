// lint_ext_san_missing.go
/************************************************
BRs: 7.1.4.2.1
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
	return !util.IsCACert(c)
}

func (l *SANMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(c, util.SubjectAlternateNameOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_missing",
		Description:   "Subscriber certificates MUST contain the Subject Alternate Name extension",
		Source:        "BRs: 7.1.4.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &SANMissing{},
	})
}
