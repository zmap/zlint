package util

import (
	"encoding/asn1"
	"regexp"
	"strings"
	"unicode"

	"github.com/zmap/zcrypto/x509/pkix"
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

// IsIA5String returns true if raw is an IA5String, and returns false otherwise.
func IsIA5String(raw []byte) bool {
	for _, b := range raw {
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

// AllAlternateNameWithTagAreIA5 returns true if all sequence members with the
// given tag are encoded as IA5 strings, and false otherwise. If it encounteres
// errors parsing asn1, err will be non-nil.
func AllAlternateNameWithTagAreIA5(ext *pkix.Extension, tag int) (bool, error) {
	var seq asn1.RawValue
	var err error
	// Unmarshal the extension as a sequence
	if _, err = asn1.Unmarshal(ext.Value, &seq); err != nil {
		return false, err
	}
	// Ensure the sequence matches what we expect for SAN/IAN
	if !seq.IsCompound || seq.Tag != asn1.TagSequence || seq.Class != asn1.ClassUniversal {
		err = asn1.StructuralError{Msg: "bad alternate name sequence"}
		return false, err
	}

	// Iterate over the sequence and look for items tagged with tag
	rest := seq.Bytes
	for len(rest) > 0 {
		var v asn1.RawValue
		rest, err = asn1.Unmarshal(rest, &v)
		if err != nil {
			return false, err
		}
		if v.Tag == tag {
			if !IsIA5String(v.Bytes) {
				return false, nil
			}
		}
	}

	return true, nil
}
