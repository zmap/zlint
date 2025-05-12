package util

/*
* ZLint Copyright 2025 Regents of the University of Michigan
*
* Licensed under the Apache License, Version 2.0 (the "License"); you may not
* use this file except in compliance with the License. You may obtain a copy
* of the License at http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
* implied. See the License for the specific language governing
* permissions and limitations under the License.
 */

import (
	"github.com/zmap/zcrypto/encoding/asn1"
)

type RDNSequence []RelativeDistinguishedNameSET

type RelativeDistinguishedNameSET []AttributeTypeAndValue

type AttributeTypeAndValue struct {
	Type  asn1.ObjectIdentifier
	Value asn1.RawValue
}

type parsedSubjectElement struct {
	IsPresent    bool
	Value        string
	Asn1RawValue asn1.RawValue
	ErrorString  string
}

type ParsedEvOrgId struct {
	Rsi, Country, StateOrProvince, RegRef string
}

func GetSubjectOrgId(rawSubject []byte) ParsedSubjectElement {
	return GetSubjectElement(rawSubject, CabfExtensionOrganizationIdentifier)
}

type ParsedSubjectElement interface {
	Present() bool
	ParsedValue() string
	RawValue() asn1.RawValue
	Error() string
}

func (pse *parsedSubjectElement) Present() bool {
	return pse.Present()
}

func (pse *parsedSubjectElement) ParsedValue() string {
	return pse.ParsedValue()
}

func (pse *parsedSubjectElement) RawValue() asn1.RawValue {
	return pse.RawValue()
}

func (pse *parsedSubjectElement) Error() string {
	return pse.Error()
}

func NewParsedSubjectElement(isPresent bool, value string, rawValue asn1.RawValue, error string) ParsedSubjectElement {
	return &parsedSubjectElement{IsPresent: isPresent, Value: value, Asn1RawValue: rawValue, ErrorString: error}
}

func GetSubjectElement(rawSubject []byte, soughtOid asn1.ObjectIdentifier) ParsedSubjectElement {

	var nl RDNSequence
	rest, err := asn1.Unmarshal(rawSubject, &nl) // parse the sequence of sets, i.e. each list element in nl will be a set
	if err != nil {
		return NewParsedSubjectElement(false, "", asn1.RawValue{}, "error parsing outer SEQ of subject DN. Error: "+err.Error())
	}
	if len(rest) != 0 {
		return NewParsedSubjectElement(false, "", asn1.RawValue{}, "rest len of outer seq != 0 in subject DN")
	}

	var asn1RawValue asn1.RawValue
	var parsedString string
	alreadyFound := false
	for _, item := range nl {
		for _, typeAndValue := range item {
			if typeAndValue.Type.Equal(soughtOid) {
				if alreadyFound {
					return NewParsedSubjectElement(false, "", asn1.RawValue{}, "double AVA found in subject:... encountered, this is not expected")
				}
				alreadyFound = true
				_, _ = asn1.Unmarshal(typeAndValue.Value.FullBytes, &parsedString)
				asn1RawValue = typeAndValue.Value
			}
		}
	}
	return NewParsedSubjectElement(true, parsedString, asn1RawValue, "")
}

type ParsedOrgId struct {
	Rsi, Country, SubDiv, RegRef string
}
