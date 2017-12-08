package util

import "encoding/asn1"

type AttributeTypeAndRawValue struct {
	Type  asn1.ObjectIdentifier
	Value asn1.RawValue
}

type AttributeTypeAndRawValueSET []AttributeTypeAndRawValue

type RawRDNSequence []AttributeTypeAndRawValueSET
