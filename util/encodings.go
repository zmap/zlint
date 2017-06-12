package util

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"regexp"
	"strings"
	"unicode"
)

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

// IsIA5String returns true if raw is an IA5String, and returns false otherwise,
// including if raw fails to unmarshal.
func IsIA5String(raw []byte) bool {
	var v asn1.RawValue
	_, err := asn1.Unmarshal(raw, &v)
	if err != nil {
		return false
	}
	for _, b := range v.Bytes {
		i := int(b)
		if i > 127 || i < 0 {
			return false
		}
	}
	return true
}

func IsInPrefSyn(name string) bool {
	// If the DNS name is just a space, it is valid
	if name == " " {
		return true
	}
	// This is the expression that matches the ABNF syntax from RFC 1034: Sec 3.5, specifically for subdomain since the " " case for domain is covered above
	prefsyn := regexp.MustCompile(`^([[:alpha:]]{1}(([[:alnum:]]|[-])*[[:alnum:]]{1})*){1}([.][[:alpha:]]{1}(([[:alnum:]]|[-])*[[:alnum:]]{1})*)*$`)
	return prefsyn.MatchString(name)
}
