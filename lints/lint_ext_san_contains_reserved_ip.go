// lint_ext_san_contains_reserved_ip.go
/************************************************
CAB: 7.1.4.2.1
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

type SANReservedIP struct {
	// Internal data here
}

func (l *SANReservedIP) Initialize() error {
	return nil
}

func (l *SANReservedIP) CheckApplies(c *x509.Certificate) bool {
	return c.NotAfter.After(util.NoReservedIP)
}

func (l *SANReservedIP) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, ip := range c.IPAddresses {
		if util.ValidIP(ip) && util.IsReservedIP(ip) {
			return ResultStruct{Result: Error}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_contains_reserved_ip",
		Description:   "Certs and expiring after 2015-11-01 must not contain a reserved ip address in the subjectAlternativeName extension.",
		Providence:    "CAB: 7.1.4.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &SANReservedIP{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.EExtSanContainsReservedIp = result },
	})
}
