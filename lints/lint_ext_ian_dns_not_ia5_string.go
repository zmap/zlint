// lint_ext_ian_dns_not_ia5_string.go
/********************************************************************
RFC 5280: 4.2.1.7
When the subjectAltName extension contains a domain name system
label, the domain name MUST be stored in the DNSName (an IA5String).
The name MUST be in the "preferred name syntax", as specified by
Section 3.5 of [RFC1034] and as modified by Section 2.1 of
[RFC1123].  Note that while uppercase and lowercase letters are
allowed in domain names, no significance is attached to the case.  In
addition, while the string " " is a legal domain name, subjectAltName
extensions with a DNSName of " " MUST NOT be used.  Finally, the use
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

type IANDNSNotIA5String struct {
	// Internal data here
}

func (l *IANDNSNotIA5String) Initialize() error {
	return nil
}

func (l *IANDNSNotIA5String) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANDNSNotIA5String) RunTest(c *x509.Certificate) (ResultStruct, error) {
	ext := util.GetExtFromCert(c, util.IssuerAlternateNameOID)
	if ext == nil {
		return ResultStruct{Result: Fatal}, nil
	}
	ok, err := util.AllAlternateNameWithTagAreIA5(ext, util.DNSNameTag)
	if err != nil {
		return ResultStruct{Result: Fatal}, nil
	}
	if ok {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_ian_dns_not_ia5_string",
		Description:   "DNSNames MUST be IA5 strings",
		Source:        "RFC 5280: 4.2.1.7",
		EffectiveDate: util.RFC2459Date,
		Test:          &IANDNSNotIA5String{},
	})
}
