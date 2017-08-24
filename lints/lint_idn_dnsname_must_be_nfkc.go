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
	"golang.org/x/net/idna"
	"golang.org/x/text/unicode/norm"
	"strings"
)

type IDNNotNFKC struct {
	// Internal data here
}

func (l *IDNNotNFKC) Initialize() error {
	return nil
}

func (l *IDNNotNFKC) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *IDNNotNFKC) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		splitLabels := strings.Split(dns, ".")
		for _, label := range splitLabels {
			if strings.HasPrefix(label, "xn--") {
				//IDN domain name
				unicodeLabel, err := idna.ToUnicode(label)
				if err != nil {
					return ResultStruct{Result: Fatal}, nil
				}
				if !norm.NFKC.IsNormalString(unicodeLabel) {
					return ResultStruct{Result: Error}, nil
				}
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_international_dns_name_not_nfkc",
		Description:   "Internationalized DNSNames must be normalized by unicode normalization form KC",
		Provenance:    "RFC 3490",
		EffectiveDate: util.RFC3490Date,
		Test:          &IDNNotNFKC{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.EInternationalDnsNameNotNfkc = result },
	})
}
