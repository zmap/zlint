package util

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
)

func DNSHasNonStringIA5(c *x509.Certificate, isIssuer bool) (bool, error) {
	oid := SANOID
	if isIssuer {
		oid = IssuerANOID
	}
	value := GetExtFromCert(c, oid).Value
	var seq asn1.RawValue
	var err error
	if _, err = asn1.Unmarshal(value, &seq); err != nil {
		return false, err
	}
	if !seq.IsCompound || seq.Tag != asn1.TagSequence || seq.Class != asn1.ClassUniversal {
		err = asn1.StructuralError{Msg: "bad IAN sequence"}
		return false, err
	}

	rest := seq.Bytes
	const dNSNameTag = 2
	for len(rest) > 0 {
		var v asn1.RawValue
		rest, err = asn1.Unmarshal(rest, &v)
		if err != nil {
			return false, err
		}
		if v.Tag == dNSNameTag {
			for _, bytes := range v.Bytes {
				if bytes > 127 {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
