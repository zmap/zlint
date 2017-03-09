// lint_ext_san_dnsname_not_fqdn.go
/**************************************************************************************************************************
7.1.4.2.1. Subject Alternative Name Extension
Certificate Field: extensions:subjectAltName
Required/Optional:  Required
Contents:  This extension MUST contain at least one entry.  Each entry MUST be either a dNSName containing
the Fully‐Qualified Domain Name or an iPAddress containing the IP address of a server.  The CA MUST
confirm that the Applicant controls the Fully‐Qualified Domain Name or IP address or has been granted the
right to use it by the Domain Name Registrant or IP address assignee, as appropriate.
Wildcard FQDNs are permitted.
**************************************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dnsFqdn struct {
	// Internal data here
}

func (l *dnsFqdn) Initialize() error {
	return nil
}

func (l *dnsFqdn) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SanOID)
}

func (l *dnsFqdn) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		if !util.IsFqdn(dns) {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_dnsname_not_fqdn",
		Description:   "SAN dnsnames must be a fully qualified domain name.",
		Providence:    "CAB: 7.1.4.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &dnsFqdn{}})
}
