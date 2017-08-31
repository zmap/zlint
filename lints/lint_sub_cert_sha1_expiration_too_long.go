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

type sha1ExpireLong struct{}

func (l *sha1ExpireLong) Initialize() error {
	return nil
}

func (l *sha1ExpireLong) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c) && (c.SignatureAlgorithm == x509.SHA1WithRSA ||
		c.SignatureAlgorithm == x509.DSAWithSHA1 ||
		c.SignatureAlgorithm == x509.ECDSAWithSHA1)
}

func (l *sha1ExpireLong) Execute(c *x509.Certificate) LintResult {
	if c.NotAfter.After(time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)) {
		return &LintResult{Status: Warn}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_cert_sha1_expiration_too_long",
		Description:   "Subscriber certificates using the SHA-1 algorithm SHOULD NOT have an expiration date later than 1 Jan 2017",
		Source:        "BRs: 7.1.3",
		EffectiveDate: time.Date(2015, time.January, 16, 0, 0, 0, 0, time.UTC),
		Lint:          &sha1ExpireLong{},
	})
}
