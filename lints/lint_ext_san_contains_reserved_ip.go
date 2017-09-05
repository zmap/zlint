// lint_ext_san_contains_reserved_ip.go
/************************************************
BRs: 7.1.4.2.1
Also as of the Effective Date, the CA SHALL NOT
issue a certificate with an Expiry Date later than
1 November 2015 with a subjectAlternativeName extension
or Subject commonName field containing a Reserved IP
Address or Internal Name.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANReservedIP struct{}

func (l *SANReservedIP) Initialize() error {
	return nil
}

func (l *SANReservedIP) CheckApplies(c *x509.Certificate) bool {
	return c.NotAfter.After(util.NoReservedIP)
}

func (l *SANReservedIP) Execute(c *x509.Certificate) *LintResult {
	for _, ip := range c.IPAddresses {
		if util.ValidIP(ip) && util.IsReservedIP(ip) {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_contains_reserved_ip",
		Description:   "Effective October 1, 2016, CAs must revoke all unexpired certificates that contains a reserved IP or internal name.",
		Source:        "BRs: 7.1.4.2.1",
		Type:          BRs,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &SANReservedIP{},
	})
}
