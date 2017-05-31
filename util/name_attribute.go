package util

import (
	"encoding/asn1"
  "github.com/zmap/zcrypto/x509"
  "github.com/zmap/zcrypto/x509/pkix"
	"strings"
)

var attributes = map[int]bool{
	// Name attributes defined in RFC 5280 appendix A
	3:  true, // id-at-commonName	AttributeType ::= { id-at 3 }
	4:  true, // id-at-surname	AttributeType ::= { id-at  4 }
	5:  true, // id-at-serialNumber	AttributeType ::= { id-at 5 }
	6:  true, // id-at-countryName	AttributeType ::= { id-at 6 }
	7:  true, // id-at-localityName	AttributeType ::= { id-at 7 }
	8:  true, // id-at-stateOrProvinceName	AttributeType ::= { id-at 8 }
	10: true, // id-at-organizationName	AttributeType ::= { id-at 10 }
	11: true, // id-at-organizationalUnitName	AttributeType ::= { id-at 11 }
	12: true, // id-at-title	AttributeType ::= { id-at 12 }
	41: true, // id-at-name	AttributeType ::= { id-at 41 }
	42: true, // id-at-givenName	AttributeType ::= { id-at 42 }
	43: true, // id-at-initials	AttributeType ::= { id-at 43 }
	44: true, // id-at-generationQualifier	AttributeType ::= { id-at 44 }
	46: true, // id-at-dnQualifier	AttributeType ::= { id-at 46 }
	// Name attributes not present in RFC 5280, but appeared in golang crypto/x509/pkix.go
	9:  true, // id-at-streetName	AttributeType ::= { id-at 9 }
	17: true, // id-at-postalCodeName	AttributeType ::= { id-at 17 }
}

func IsAttributeInList(in int) bool {
	return attributes[in]
}

// returns 0 when there is no space; bit 0 indicates if there is prefix space; bit 1 suffix space
func DNSAttributeHasSpace(c *x509.Certificate, isIssuer bool) (int, error) {
  var name pkix.RDNSequence
	raw := c.RawSubject
	ret := 0
	if isIssuer {
		raw = c.RawIssuer
	}
  if _, err := asn1.Unmarshal(raw, &name); err != nil {
    return ret, err
  }
  for _, rdn := range name {
    if len(rdn) == 0 {
      continue
    }
    atv := rdn[0]
    value, ok := atv.Value.(string)
    if !ok {
      continue
    }

    t := atv.Type
    if len(t) == 4 && t[0] == 2 && t[1] == 5 && t[2] == 4 && IsAttributeInList(t[3]) {
      if strings.HasPrefix(value, " ") {
        ret |= 1
      }
			if strings.HasSuffix(value, " ") {
        ret |= 1 << 1
      }
    }
  }
  return ret, nil
}
