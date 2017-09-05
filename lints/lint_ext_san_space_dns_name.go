// lint_ext_san_space_dns_name.go
/************************************************************************
RFC 5280: 4.2.1.6
When the subjectAltName extension contains a domain name system
   label, the domain name MUST be stored in the dNSName (an IA5String).
   The name MUST be in the "preferred name syntax", as specified by
   Section 3.5 of [RFC1034] and as modified by Section 2.1 of
   [RFC1123].  Note that while uppercase and lowercase letters are
   allowed in domain names, no significance is attached to the case.  In
   addition, while the string " " is a legal domain name, subjectAltName
   extensions with a dNSName of " " MUST NOT be used.  Finally, the use
   of the DNS representation for Internet mail addresses
   (subscriber.example.com instead of subscriber@example.com) MUST NOT
   be used; such identities are to be encoded as rfc822Name.  Rules for
   encoding internationalized domain names are specified in Section 7.2.
************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANIsSpaceDNS struct{}

func (l *SANIsSpaceDNS) Initialize() error {
	return nil
}

func (l *SANIsSpaceDNS) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANIsSpaceDNS) Execute(c *x509.Certificate) *LintResult {
	for _, dns := range c.DNSNames {
		if dns == " " {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_space_dns_name",
		Description:   "The dNSName ` ` MUST NOT be used",
		Source:        "RFC 5280: 4.2.1.6",
		Type:          RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &SANIsSpaceDNS{},
	})
}
