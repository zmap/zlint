// lint_ext_ian_dns_not_ia5_string.go
/********************************************************************
RFC 5280: 4.2.1.7
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
	"encoding/asn1"

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
	return util.IsExtInCert(c, util.IssuerANOID)
}

func (l *IANDNSNotIA5String) RunTest(c *x509.Certificate) (ResultStruct, error) {
	value := util.GetExtFromCert(c, util.IssuerANOID).Value
	var seq asn1.RawValue
	var err error
	if _, err = asn1.Unmarshal(value, &seq); err != nil {
		return ResultStruct{Result: NA}, err
	}
	if !seq.IsCompound || seq.Tag != 16 || seq.Class != 0 {
		err = asn1.StructuralError{Msg: "bad IAN sequence"}
		return ResultStruct{Result: Fatal}, err
	}

	rest := seq.Bytes
	for len(rest) > 0 {
		var v asn1.RawValue
		rest, err = asn1.Unmarshal(rest, &v)
		if err != nil {
			return ResultStruct{Result: NA}, err
		}
		if v.Tag == 2 {
			for _, bytes := range v.Bytes {
				if bytes > 127 {
					return ResultStruct{Result: Error}, nil
				}
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_ian_dns_not_ia5_string",
		Description:   "DNSNames MUST be IA5 strings",
		Providence:    "RFC 5280: 4.2.1.7",
		EffectiveDate: util.RFC2459Date,
		Test:          &IANDNSNotIA5String{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.EExtIanDnsNotIa5String = result },
	})
}
