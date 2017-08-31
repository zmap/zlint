// lint_subject_contains_reserved_ip.go
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
	"net"
)

type subjectReservedIP struct{}

func (l *subjectReservedIP) Initialize() error {
	return nil
}

func (l *subjectReservedIP) CheckApplies(c *x509.Certificate) bool {
	return c.NotAfter.After(util.NoReservedIP)
}

func (l *subjectReservedIP) Execute(c *x509.Certificate) ResultStruct {
	if ip := net.ParseIP(c.Subject.CommonName); ip != nil {
		if !util.IsReservedIP(ip) {
			return ResultStruct{Result: Pass}
		} else {
			return ResultStruct{Result: Error}
		}
	} else {
		return ResultStruct{Result: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_contains_reserved_ip",
		Description:   "Certificates expiring later than 11 Jan 2015 MUST NOT contain a reserved IP address in the common name field",
		Source:        "BRs: 7.1.4.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subjectReservedIP{},
	})
}
