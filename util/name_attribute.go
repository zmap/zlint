package util

import (
	"encoding/asn1"
	"strings"
	"unicode"

	"github.com/zmap/zcrypto/x509/pkix"
)

type empty struct{}

var nameAttributePrefix = asn1.ObjectIdentifier{2, 5, 4}
var nameAttributeLeaves = map[int]empty{
	// Name attributes defined in RFC 5280 appendix A
	3:  empty{}, // id-at-commonName	AttributeType ::= { id-at 3 }
	4:  empty{}, // id-at-surname	AttributeType ::= { id-at  4 }
	5:  empty{}, // id-at-serialNumber	AttributeType ::= { id-at 5 }
	6:  empty{}, // id-at-countryName	AttributeType ::= { id-at 6 }
	7:  empty{}, // id-at-localityName	AttributeType ::= { id-at 7 }
	8:  empty{}, // id-at-stateOrProvinceName	AttributeType ::= { id-at 8 }
	10: empty{}, // id-at-organizationName	AttributeType ::= { id-at 10 }
	11: empty{}, // id-at-organizationalUnitName	AttributeType ::= { id-at 11 }
	12: empty{}, // id-at-title	AttributeType ::= { id-at 12 }
	41: empty{}, // id-at-name	AttributeType ::= { id-at 41 }
	42: empty{}, // id-at-givenName	AttributeType ::= { id-at 42 }
	43: empty{}, // id-at-initials	AttributeType ::= { id-at 43 }
	44: empty{}, // id-at-generationQualifier	AttributeType ::= { id-at 44 }
	46: empty{}, // id-at-dnQualifier	AttributeType ::= { id-at 46 }

	// Name attributes not present in RFC 5280, but appeared in golang crypto/x509/pkix.go
	9:  empty{}, // id-at-streetName	AttributeType ::= { id-at 9 }
	17: empty{}, // id-at-postalCodeName	AttributeType ::= { id-at 17 }
}

// IsNameAttribute returns true if the given ObjectIdentifier corresponds with
// the type of any name attribute for PKIX.
func IsNameAttribute(oid asn1.ObjectIdentifier) bool {
	if len(oid) != 4 {
		return false
	}
	if !nameAttributePrefix.Equal(oid[0:3]) {
		return false
	}
	_, ok := nameAttributeLeaves[oid[3]]
	return ok
}

// CheckRDNSequenceWhiteSpace returns true if there is leading or trailing
// whitespace in any name attribute in the sequence, respectively.
func CheckRDNSequenceWhiteSpace(raw []byte) (leading, trailing bool, err error) {
	var seq pkix.RDNSequence
	if _, err = asn1.Unmarshal(raw, &seq); err != nil {
		return
	}
	for _, rdn := range seq {
		for _, atv := range rdn {
			if !IsNameAttribute(atv.Type) {
				continue
			}
			value, ok := atv.Value.(string)
			if !ok {
				continue
			}
			if leftStrip := strings.TrimLeftFunc(value, unicode.IsSpace); leftStrip != value {
				leading = true
			}
			if rightStrip := strings.TrimRightFunc(value, unicode.IsSpace); rightStrip != value {
				trailing = true
			}
		}
	}
	return
}
