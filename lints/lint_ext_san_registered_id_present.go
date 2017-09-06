// lint_ext_san_registered_id_present.go
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

type SANRegId struct{}

func (l *SANRegId) Initialize() error {
	return nil
}

func (l *SANRegId) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANRegId) Execute(c *x509.Certificate) *LintResult {
	if c.RegisteredIDs != nil {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_registered_id_present",
		Description:   "The Subject Alternate Name extension MUST contain only 'dnsName' and 'ipaddress' name types.",
		Citation:      "BRs: 7.1.4.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &SANRegId{},
	})
}
