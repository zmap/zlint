/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

package util

import (
	"encoding/asn1"
	"fmt"
	"reflect"
	"regexp"

	"github.com/zmap/zcrypto/x509"
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

type cabfOrgIdExt struct {
	Rsi             string `asn1:"printable"`
	Country         string `asn1:"printable"`
	StateOrProvince string `asn1:"printable,optional,tag:0"`
	RegRef          string `asn1:"utf8"`
}

// Returns an error string if the components of the parsed cabfOrganizationIdentifier
// do not comply with the naming schema allowed by CAB Forum.
func CheckParsedEvOrgId(parsedEvOrgId ParsedEvOrgId) string {
	if len(parsedEvOrgId.Rsi) != 3 {
		return "registrationSchemeIdentifier has not length 3"
	}
	if len(parsedEvOrgId.Country) != 2 {
		return "registrationCountry has not length 2"
	}
	if len(parsedEvOrgId.StateOrProvince) > 128 {
		return "registrationStateOrProvince has length more than 128"
	}

	regExpCorrectRsi := regexp.MustCompile(`^(NTR|VAT|PSD)`)

	if !regExpCorrectRsi.MatchString(parsedEvOrgId.Rsi) {
		return fmt.Sprintf("registrationSchemeIdentifier has an invalid value: %v", parsedEvOrgId.Rsi)
	}

	regExpCorrectCountry := regexp.MustCompile(`^[A-Z]{2}`)

	if !regExpCorrectCountry.MatchString(parsedEvOrgId.Country) {
		return fmt.Sprintf("registrationCountry has an invalid value: %v", parsedEvOrgId.Country)
	}

	return ""
}

// Returns the parsed organization identifier with its components (registrationSchemeIdentifier,
// registrationCountryPrintableString, registrationStateOrProvince, registrationReferenceUTF8String)
// from the cabfOrganizationIdentifier extension or an error if the extension could not be parsed.
func ParseCabfOrgIdExt(c *x509.Certificate) (string, ParsedEvOrgId) {
	var result ParsedEvOrgId

	ext := GetExtFromCert(c, CabfExtensionOrganizationIdentifier)
	var parsedExt cabfOrgIdExt
	// check that we can parse the extension:
	rest, err := asn1.Unmarshal(ext.Value, &parsedExt)
	if len(rest) != 0 {
		return "trailing bytes after extension", result
	}
	if err != nil {
		return "could not parse extension value:" + err.Error(), result
	}
	errStr := CheckAsn1Reencoding(reflect.ValueOf(parsedExt).Interface(), ext.Value, "invalid string type in extension")
	if errStr != "" {
		return errStr, result
	}

	result.Country = parsedExt.Country
	result.RegRef = parsedExt.RegRef
	result.Rsi = parsedExt.Rsi
	result.StateOrProvince = parsedExt.StateOrProvince

	return "", result
}

// Returns the parsed organization identifier with its components (registrationSchemeIdentifier,
//registrationCountryPrintableString, registrationStateOrProvince, registrationReferenceUTF8String)
// or an error if the value of the organization identifier does not have he following form:
//
//- 3 character Registration Scheme identifier;
//- 2 character ISO 3166 country code for the nation in which the Registration Scheme is operated, or if the scheme is operated globally ISO 3166 code "XG" shall be used;
//- For the NTR Registration Scheme identifier, if required under Section 9.2.4, a 2 character ISO 3166-2 identifier for the subdivision (state or province) of the nation in which the Registration Scheme is operated, preceded by plus "+" (0x2B (ASCII), U+002B (UTF-8));
//- a hyphen-minus "-" (0x2D (ASCII), U+002D (UTF-8));
//- Registration Reference allocated in accordance with the identified Registration Scheme
//
// ETSI specification allows LEI. This is also considered.
//
func ParseOrganizationIdentifier(oi string, isEtsi bool) (string, ParsedEvOrgId) {
	var result ParsedEvOrgId
	re_ntr := regexp.MustCompile(`^(NTR)([A-Z]{2})([+]([A-Z]{2}))?-(.+)$`)
	re_vat_psd := regexp.MustCompile(`^(VAT|PSD)([A-Z]{2})(())-(.+)$`)
	re_lei := regexp.MustCompile(`^(LEI)(XG)(())-(.+)$`)
	var sm []string
	if re_ntr.MatchString(oi) {
		sm = re_ntr.FindStringSubmatch(oi)
	} else if re_vat_psd.MatchString(oi) {
		sm = re_vat_psd.FindStringSubmatch(oi)
	} else if re_lei.MatchString(oi) {
		if isEtsi {
			sm = re_lei.FindStringSubmatch(oi)
		} else {
			return "CAB/F subject:organizationIdentifier does not allow LEI", result
		}
	} else {
		return "CAB/F subject:organizationIdentifier has an invalid format", result
	}
	result.Rsi = sm[1]
	result.Country = sm[2]
	result.StateOrProvince = sm[3]
	result.RegRef = sm[5]
	return "", result
}

func GetSubjectOrgId(rawSubject []byte) parsedSubjectElement {
	return GetSubjectElement(rawSubject, OrganizationIdentifierOID)
}

func GetSubjectElement(rawSubject []byte, soughtOid asn1.ObjectIdentifier) parsedSubjectElement {
	result := parsedSubjectElement{IsPresent: false, Value: "", ErrorString: ""}
	var nl RDNSequence

	rest, err := asn1.Unmarshal(rawSubject, &nl) // parse the sequence of sets, i.e. each list element in nl will be a set
	if err != nil {
		return parsedSubjectElement{IsPresent: false, Value: "", ErrorString: "error parsing outer SEQ of subject DN"}
	}
	if len(rest) != 0 {
		return parsedSubjectElement{IsPresent: false, ErrorString: "rest len of outer seq != 0 in subject DN", Value: ""}
	}
	for _, item := range nl {
		for _, typeAndValue := range item {
			if typeAndValue.Type.Equal(soughtOid) {
				if result.IsPresent {
					AppendToStringSemicolonDelim(&result.ErrorString, "double AVA found in subject:... encountered, this is not expected")
					return result
				}
				result.IsPresent = true
				var parsedString string
				_, _ = asn1.Unmarshal(typeAndValue.Value.FullBytes, &parsedString)
				result.Value = parsedString
				result.Asn1RawValue = typeAndValue.Value
			}
		}
	}
	return result
}

type ParsedOrgId struct {
	Rsi, Country, SubDiv, RegRef string
}
