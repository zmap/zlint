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
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type sanDnsNotIa5 struct {
	// Internal data here
}

func (l *sanDnsNotIa5) Initialize() error {
	return nil
}

func (l *sanDnsNotIa5) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SanOID)
}

func (l *sanDnsNotIa5) RunTest(c *x509.Certificate) (ResultStruct, error) {
	value := util.GetExtFromCert(c, util.SanOID).Value
	var seq asn1.RawValue
	var err error
	if _, err = asn1.Unmarshal(value, &seq); err != nil {
		return ResultStruct{Result: NA}, err
	}
	if !seq.IsCompound || seq.Tag != 16 || seq.Class != 0 {
		err = asn1.StructuralError{Msg: "bad SAN sequence"}
		return ResultStruct{Result: NA}, err
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
		Name:          "ext_san_dns_not_ia5_string",
		Description:   "dNSNames are IA5 strings",
		Providence:    "RFC 5280: 4.2.1.6",
		EffectiveDate: util.RFC2459Date,
		Test:          &sanDnsNotIa5{}})
}
