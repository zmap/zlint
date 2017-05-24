// lint_ext_san_dns_not_ia5_string.go
/********************************************************************
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
********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANDNSNotIA5String struct {
	// Internal data here
}

func (l *SANDNSNotIA5String) Initialize() error {
	return nil
}

func (l *SANDNSNotIA5String) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SANOID)
}

func (l *SANDNSNotIA5String) RunTest(c *x509.Certificate) (ResultStruct, error) {
	notIA5, err := util.DNSHasNonStringIA5(c, false)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	if notIA5 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}
func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_dns_not_ia5_string",
		Description:   "dNSNames are IA5 strings",
		Providence:    "RFC 5280: 4.2.1.6",
		EffectiveDate: util.RFC2459Date,
		Test:          &SANDNSNotIA5String{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.EExtSanDnsNotIa5String = result },
	})
}
