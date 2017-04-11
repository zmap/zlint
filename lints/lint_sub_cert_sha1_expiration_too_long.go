// lint_sub_cert_sha1_expiration_too_long.go
/***************************************************************************************************************
Effective 16 January 2015, CAs SHOULD NOT issue Subscriber Certificates utilizing the SHA‐1 algorithm with
an Expiry Date greater than 1 January 2017 because Application Software Providers are in the process of
deprecating and/or removing the SHA‐1 algorithm from their software, and they have communicated that
CAs and Subscribers using such certificates do so at their own risk.
****************************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"time"
)

type SHA1CertifcateExpirationTooLong struct {
	// Internal data here
}

func (l *SHA1CertifcateExpirationTooLong) Initialize() error {
	return nil
}

func (l *SHA1CertifcateExpirationTooLong) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCaCert(c) && c.SignatureAlgorithm == x509.SHA1WithRSA
}

func (l *SHA1CertifcateExpirationTooLong) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.NotAfter.After(time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)) {
		return ResultStruct{Result: Warn}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_cert_sha1_expiration_too_long",
		Description:   "Subscriber certificates using the SHA1 algorithm should not have an expiration date greater than 1 Jan 2017",
		Providence:    "CAB: 7.1.3",
		EffectiveDate: time.Date(2015, time.January, 16, 0, 0, 0, 0, time.UTC),
		Test:          &SHA1CertifcateExpirationTooLong{}})
}
