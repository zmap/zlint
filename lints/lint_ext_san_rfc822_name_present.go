// lint_ext_san_rfc822_name_present.go
/************************************************************************************************************
7.1.4.2.1. Subject Alternative Name Extension
Certificate Field: extensions:subjectAltName
Required/Optional:  Required
Contents:  This extension MUST contain at least one entry.  Each entry MUST be either a dNSName containing
the Fully‐Qualified Domain Name or an iPAddress containing the IP address of a server.  The CA MUST
confirm that the Applicant controls the Fully‐Qualified Domain Name or IP address or has been granted the
right to use it by the Domain Name Registrant or IP address assignee, as appropriate.
Wildcard FQDNs are permitted.
*************************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANRfc822 struct {
	// Internal data here
}

func (l *SANRfc822) Initialize() error {
	return nil
}

func (l *SANRfc822) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANRfc822) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.EmailAddresses != nil {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_rfc822_name_present",
		Description:   "The Subject Alternate Name extension MUST contain only 'dnsName' and 'ipaddress' name types.",
		Source:        "BRs: 7.1.4.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &SANRfc822{},
	})
}
