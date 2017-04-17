package util

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
