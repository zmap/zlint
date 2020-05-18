/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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
	"bytes"
	"encoding/asn1"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/zmap/zcrypto/x509"
)

var EtsiQcStmtOidList = [...]*asn1.ObjectIdentifier{
	&IdEtsiQcsQcCompliance,
	&IdEtsiQcsQcLimitValue,
	&IdEtsiQcsQcRetentionPeriod,
	&IdEtsiQcsQcSSCD,
	&IdEtsiQcsQcEuPDS,
	&IdEtsiQcsQcType,
	&IdEtsiPsd2Statem,
}

type anyContent struct {
	Raw asn1.RawContent
}

type qcStatementWithInfoField struct {
	Oid asn1.ObjectIdentifier
	Any asn1.RawValue
}

type qcStatementWithoutInfoField struct {
	Oid asn1.ObjectIdentifier
}

// === etsi base ==>
type etsiBase struct {
	errorInfo string
	isPresent bool
}

func (this etsiBase) GetErrorInfo() string {
	return this.errorInfo
}

func (this etsiBase) IsPresent() bool {
	return this.isPresent
}

// <== etsi base ===

type EtsiQcStmtIf interface {
	GetErrorInfo() string
	IsPresent() bool
}

type Etsi421QualEuCert struct {
	etsiBase
}

type Etsi423QcType struct {
	etsiBase
	TypeOids []asn1.ObjectIdentifier
}

type EtsiQcSscd struct {
	etsiBase
}

type EtsiMonetaryValueAlph struct {
	Iso4217CurrencyCodeAlph string `asn1:"printable"`
	Amount                  int
	Exponent                int
}
type EtsiMonetaryValueNum struct {
	Iso4217CurrencyCodeNum int
	Amount                 int
	Exponent               int
}

type EtsiQcLimitValue struct {
	etsiBase
	Amount       int
	Exponent     int
	IsNum        bool
	CurrencyAlph string
	CurrencyNum  int
}

type EtsiQcRetentionPeriod struct {
	etsiBase
	Period int
}
type PdsLocation struct {
	Url      string `asn1:"ia5"`
	Language string `asn1:"printable"`
}
type EtsiQcPds struct {
	etsiBase
	PdsLocations []PdsLocation
}

// ==== QcStatement 2 (RFC3739)types ===>

type DecodedQcS2 struct {
	etsiBase
	Decoded QcStatemt2
}
type QcStatemt2 struct {
	SemanticsId        asn1.ObjectIdentifier       `asn1:"optional"`
	NameRegAuthorities NameRegistrationAuthorities `asn1:"optional"`
}

type NameRegistrationAuthorities []asn1.RawValue

// <=== QcStatement 2 (RFC3739)types ====

// ==== PSD2 QcStatement types ===>
type Psd2RoleOfPspType int

const (
	RoleAs Psd2RoleOfPspType = 1
	RolePi Psd2RoleOfPspType = 2
	RoleAi Psd2RoleOfPspType = 3
	RoleIc Psd2RoleOfPspType = 4
)

// 	=== ASN.1 Types ==>
type Psd2RoleOfPsp struct {
	RoleType      asn1.ObjectIdentifier
	RoleOfPspName string `asn1:"utf8"`
}

type EtsiPsd2QcStatem struct {
	Roles           []Psd2RoleOfPsp
	NCAName         string `asn1:"utf8"`
	CountryAndNCAId string `asn1:"utf8"`
}

// 	<== ASN.1 Types ===

type EtsiPsd2 struct {
	etsiBase
	DecodedPsd2Statm EtsiPsd2QcStatem
}

func (this EtsiPsd2) getCountryAndNcaId() (string, string) {
	runes := []rune(this.DecodedPsd2Statm.CountryAndNCAId)
	if len(this.DecodedPsd2Statm.CountryAndNCAId) < 4 || !unicode.IsUpper(runes[0]) || !unicode.IsUpper(runes[1]) || runes[2] != '-' {
		return "", ""
	}
	return string(runes[0:2]), string(runes[3:])
}

func (this EtsiPsd2) GetNcaCountry() string {
	co, _ := this.getCountryAndNcaId()
	return co
}
func (this EtsiPsd2) GetNcaId() string {
	_, ncaId := this.getCountryAndNcaId()
	return ncaId
}

// <=== PSD2 QcStatement types ====

func CheckAsn1Reencoding(i interface{}, originalEncoding []byte, appendIfComparisonFails string) string {
	return CheckAsn1ReencodingWithParams(i, originalEncoding, appendIfComparisonFails, "")
}

func CheckAsn1ReencodingWithParams(i interface{}, originalEncoding []byte, appendIfComparisonFails string, params string) string {
	result := ""
	reencoded, marshErr := asn1.MarshalWithParams(i, params)
	if marshErr != nil {
		AppendToStringSemicolonDelim(&result, fmt.Sprintf("error reencoding ASN1 value of statementInfo field: %s",
			marshErr))
	}
	if !bytes.Equal(reencoded, originalEncoding) {
		AppendToStringSemicolonDelim(&result, appendIfComparisonFails)
	}
	return result
}

func CertHasSubjectOrgIdWithPrefix(c *x509.Certificate, prefix string) bool {

	if !IsExtInCert(c, QcStateOid) {
		return false
	}
	if !ParseQcStatem(GetExtFromCert(c, QcStateOid).Value, IdEtsiPsd2Statem).IsPresent() {
		return false
	}

	orgId := GetSubjectOrgId(c.RawSubject)
	if len(orgId.ErrorString) != 0 || !orgId.IsPresent {
		return false
	}
	runes := []rune(orgId.Value)
	prefixLen := len(prefix)
	if len(runes) < prefixLen || string(runes[0:prefixLen]) != prefix {
		return false
	}
	return true
}

type EtsiPsd2OrgId struct {
	Rsi, Country, NcaId, PspId string
}

func ParseEtsiPsd2OrgId(oi *string) (string, EtsiPsd2OrgId) {
	var result EtsiPsd2OrgId
	re_psd := regexp.MustCompile(`^(PSD)([A-Z]{2})-([A-Z]{2,8})-(.+)$`)
	re_generic := regexp.MustCompile(`^(.{3})([A-Z]{2})()-(.+)$`)
	var sm []string
	if re_psd.MatchString(*oi) {
		sm = re_psd.FindStringSubmatch(*oi)
	} else if !strings.HasPrefix(*oi, "PSD") && re_generic.MatchString(*oi) {
		sm = re_generic.FindStringSubmatch(*oi)
	} else {
		return "invalid format of PSD2 organizationIdentifier", result
	}
	result.Rsi = sm[1]
	result.Country = sm[2]
	result.NcaId = sm[3]
	result.PspId = sm[4]
	return "", result
}

func CheckEtsiPsd2OrgIdPsd(oi *string) string {
	errStr, x := ParseEtsiPsd2OrgId(oi)
	if len(errStr) != 0 {
		return errStr
	}
	if x.Rsi != "PSD" {
		return "ETSI PSD2 OrganizationIdentifier does not start with 'PSD'"
	}
	return ""
}

func GetEtsiQcTypes(c *x509.Certificate) []asn1.ObjectIdentifier {
	var result []asn1.ObjectIdentifier
	ext := GetExtFromCert(c, QcStateOid)
	if ext == nil {
		return nil
	}
	s := ParseQcStatem(ext.Value, IdEtsiQcsQcType)
	if len(s.GetErrorInfo()) != 0 {
		return nil
	}
	if !s.IsPresent() {
		return result
	}
	qcType := s.(Etsi423QcType)
	result = append(result, qcType.TypeOids...)
	return result
}

func HasCertAnyEtsiQcpPolicy(c *x509.Certificate) bool {
	for _, p := range c.PolicyIdentifiers {
		if p.Equal(IdEtsiPolicyQcpNatural) || p.Equal(IdEtsiPolicyQcpLegal) || p.Equal(IdEtsiPolicyQcpNaturalQscd) || p.Equal(IdEtsiPolicyQcpLegalQscd) || p.Equal(IdEtsiPolicyQcpWeb) {
			return true
		}
	}
	return false

}

func HasCertPolicy(c *x509.Certificate, soughtPolicyOid asn1.ObjectIdentifier) bool {

	for _, policyOid := range c.PolicyIdentifiers {
		if policyOid.Equal(soughtPolicyOid) {
			return true
		}
	}
	return false
}

func HasCertEtsiQcType(c *x509.Certificate, soughtTypeOid asn1.ObjectIdentifier) bool {
	typeList := GetEtsiQcTypes(c)
	if typeList == nil {
		return false
	}
	for _, typeOid := range typeList {
		if typeOid.Equal(soughtTypeOid) {
			return true
		}
	}
	return false
}

func HasCertAnyEtsiQcStatement(c *x509.Certificate) bool {
	ext := GetExtFromCert(c, QcStateOid)
	if ext == nil {
		return false
	}
	return IsAnyEtsiQcStatementPresent(ext.Value)
}

func IsAnyEtsiQcStatementPresent(extVal []byte) bool {
	for _, oid := range EtsiQcStmtOidList {
		r := ParseQcStatem(extVal, *oid)
		if r.IsPresent() {
			return true
		}
	}
	return false
}

func IsQcStatemPresent(c *x509.Certificate, oid *asn1.ObjectIdentifier) (string, bool) {
	if !IsExtInCert(c, QcStateOid) {
		return "", false
	}
	qcs := ParseQcStatem(GetExtFromCert(c, QcStateOid).Value, *oid)
	if qcs.GetErrorInfo() != "" {
		return qcs.GetErrorInfo(), qcs.IsPresent()
	}
	return "", qcs.IsPresent()
}

func CheckNationalScheme(oi string) bool {
	if len(oi) < 6 {
		return false
	}
	re := regexp.MustCompile(`^.{2}:[A-Z]{2}-.+$`)
	return re.MatchString(oi)
}

func GetQcStatemExtValue(c *x509.Certificate) []byte {
	return GetExtFromCert(c, QcStateOid).Value
}

func ParseQcStatem(extVal []byte, sought asn1.ObjectIdentifier) EtsiQcStmtIf {
	sl := make([]anyContent, 0)
	rest, err := asn1.Unmarshal(extVal, &sl)
	if err != nil {
		return etsiBase{errorInfo: "error parsing outer SEQ", isPresent: true}
	}
	if len(rest) != 0 {
		return etsiBase{errorInfo: "rest len of outer seq != 0", isPresent: true}
	}

	for _, raw := range sl {
		parseErrorString := "format error in at least one QC statement within the QC statements extension." +
			" this message may appear multiple times for the same error cause."
		var statem qcStatementWithInfoField
		rest, err = asn1.Unmarshal(raw.Raw, &statem)
		if err != nil {
			var statemWithoutInfo qcStatementWithoutInfoField

			rest, err = asn1.Unmarshal(raw.Raw, &statemWithoutInfo)
			if err != nil || len(rest) != 0 {
				return etsiBase{errorInfo: parseErrorString, isPresent: false}
			}
			copy(statem.Oid, statemWithoutInfo.Oid)
			if len(statem.Any.FullBytes) != 0 {
				return etsiBase{errorInfo: "internal error, default optional content len is not zero"}
			}
		} else if 0 != len(rest) {
			return etsiBase{errorInfo: parseErrorString, isPresent: false}
		}

		if !statem.Oid.Equal(sought) {
			continue
		}
		if statem.Oid.Equal(IdEtsiQcsQcCompliance) {
			return handleIdEtsiQcsQcCompliance(statem, raw)
		} else if statem.Oid.Equal(IdEtsiQcsQcLimitValue) {
			return handleIdEtsiQcsQcLimitValue(statem)
		} else if statem.Oid.Equal(IdEtsiQcsQcRetentionPeriod) {
			return handleIdEtsiQcsQcRetentionPeriod(statem)
		} else if statem.Oid.Equal(IdEtsiQcsQcSSCD) {
			return handleIdEtsiQcsQcSSCD(statem, raw)
		} else if statem.Oid.Equal(IdEtsiQcsQcEuPDS) {
			return handleIdEtsiQcsQcEuPDS(statem)
		} else if statem.Oid.Equal(IdEtsiQcsQcType) {
			return handleIdEtsiQcsQcType(statem)
		} else if statem.Oid.Equal(IdEtsiPsd2Statem) {
			return handleIdEtsiPsd2Statem(statem)
		} else if statem.Oid.Equal(IdQcsPkixQCSyntaxV2) {
			return handleIdQcsPkixQCSyntaxV2(statem)
		} else {
			return etsiBase{errorInfo: "", isPresent: true}
		}
	}

	return etsiBase{errorInfo: "", isPresent: false}

}

func handleIdQcsPkixQCSyntaxV2(statem qcStatementWithInfoField) EtsiQcStmtIf {
	var qcs2Statem DecodedQcS2
	qcs2Statem.isPresent = true
	if len(statem.Any.FullBytes) == 0 {
		return qcs2Statem
	}
	rest, err := asn1.Unmarshal(statem.Any.FullBytes, &qcs2Statem.Decoded)
	if err != nil {
		AppendToStringSemicolonDelim(&qcs2Statem.errorInfo, "error parsing statement: "+err.Error())
	}
	if len(rest) != 0 {
		AppendToStringSemicolonDelim(&qcs2Statem.errorInfo, "trailing bytes after QcStatement")
	}
	return qcs2Statem
}

func handleIdEtsiPsd2Statem(statem qcStatementWithInfoField) EtsiQcStmtIf {
	var psd2Statem EtsiPsd2
	psd2Statem.isPresent = true
	rest, err := asn1.Unmarshal(statem.Any.FullBytes, &psd2Statem.DecodedPsd2Statm)
	if len(rest) != 0 || err != nil {
		return etsiBase{errorInfo: "error parsing IdEtsiPsd2Statem extension statementInfo field", isPresent: true}
	}
	if psd2Statem.DecodedPsd2Statm.CountryAndNCAId == "" || psd2Statem.DecodedPsd2Statm.NCAName == "" {
		AppendToStringSemicolonDelim(&psd2Statem.errorInfo, "field has length 0")
	}
	for _, role := range psd2Statem.DecodedPsd2Statm.Roles {
		if role.RoleOfPspName == "" {
			AppendToStringSemicolonDelim(&psd2Statem.errorInfo, "field has length 0")
		}
	}
	AppendToStringSemicolonDelim(&psd2Statem.errorInfo,
		CheckAsn1Reencoding(reflect.ValueOf(psd2Statem.DecodedPsd2Statm).Interface(), statem.Any.FullBytes,
			"error with ASN.1 encoding, possibly a wrong ASN.1 string type was used"))
	return psd2Statem
}

func handleIdEtsiQcsQcType(statem qcStatementWithInfoField) EtsiQcStmtIf {
	var qcType Etsi423QcType
	qcType.isPresent = true
	rest, err := asn1.Unmarshal(statem.Any.FullBytes, &qcType.TypeOids)
	if len(rest) != 0 || err != nil {
		return etsiBase{errorInfo: "error parsing IdEtsiQcsQcType extension statementInfo field", isPresent: true}
	}
	return qcType
}

func handleIdEtsiQcsQcEuPDS(statem qcStatementWithInfoField) EtsiQcStmtIf {
	etsiObj := EtsiQcPds{etsiBase: etsiBase{isPresent: true}}
	rest, err := asn1.Unmarshal(statem.Any.FullBytes, &etsiObj.PdsLocations)
	if len(rest) != 0 || err != nil {
		etsiObj.errorInfo = "error parsing the statementInfo field"
	} else {
		AppendToStringSemicolonDelim(&etsiObj.errorInfo,
			CheckAsn1Reencoding(reflect.ValueOf(etsiObj.PdsLocations).Interface(), statem.Any.FullBytes,
				"error with ASN.1 encoding, possibly a wrong ASN.1 string type was used"))
	}
	return etsiObj
}

func handleIdEtsiQcsQcSSCD(statem qcStatementWithInfoField, raw anyContent) EtsiQcStmtIf {
	etsiObj := EtsiQcSscd{etsiBase: etsiBase{isPresent: true}}
	statemWithoutInfo := qcStatementWithoutInfoField{Oid: statem.Oid}
	AppendToStringSemicolonDelim(&etsiObj.errorInfo, CheckAsn1Reencoding(reflect.ValueOf(statemWithoutInfo).Interface(), raw.Raw,
		"invalid format of ETSI SCSD statement"))
	return etsiObj
}

func handleIdEtsiQcsQcRetentionPeriod(statem qcStatementWithInfoField) EtsiQcStmtIf {
	etsiObj := EtsiQcRetentionPeriod{etsiBase: etsiBase{isPresent: true}}
	rest, err := asn1.Unmarshal(statem.Any.FullBytes, &etsiObj.Period)

	if len(rest) != 0 || err != nil {
		etsiObj.errorInfo = "error parsing the statementInfo field"
	}
	return etsiObj
}

func handleIdEtsiQcsQcLimitValue(statem qcStatementWithInfoField) EtsiQcStmtIf {
	etsiObj := EtsiQcLimitValue{etsiBase: etsiBase{isPresent: true}}
	numErr := false
	alphErr := false
	var numeric EtsiMonetaryValueNum
	var alphabetic EtsiMonetaryValueAlph
	restNum, errNum := asn1.Unmarshal(statem.Any.FullBytes, &numeric)
	if len(restNum) != 0 || errNum != nil {
		numErr = true
	} else {
		etsiObj.IsNum = true
		etsiObj.Amount = numeric.Amount
		etsiObj.Exponent = numeric.Exponent
		etsiObj.CurrencyNum = numeric.Iso4217CurrencyCodeNum

	}
	if numErr {
		restAlph, errAlph := asn1.Unmarshal(statem.Any.FullBytes, &alphabetic)
		if len(restAlph) != 0 || errAlph != nil {
			alphErr = true
		} else {
			etsiObj.IsNum = false
			etsiObj.Amount = alphabetic.Amount
			etsiObj.Exponent = alphabetic.Exponent
			etsiObj.CurrencyAlph = alphabetic.Iso4217CurrencyCodeAlph
			AppendToStringSemicolonDelim(&etsiObj.errorInfo,
				CheckAsn1Reencoding(reflect.ValueOf(alphabetic).Interface(),
					statem.Any.FullBytes, "error with ASN.1 encoding, possibly a wrong ASN.1 string type was used"))
		}
	}
	if numErr && alphErr {
		etsiObj.errorInfo = "error parsing the ETSI Qc Statement statementInfo field"
	}
	return etsiObj
}

func handleIdEtsiQcsQcCompliance(statem qcStatementWithInfoField, raw anyContent) EtsiQcStmtIf {
	etsiObj := Etsi421QualEuCert{etsiBase: etsiBase{isPresent: true}}
	statemWithoutInfo := qcStatementWithoutInfoField{Oid: statem.Oid}
	AppendToStringSemicolonDelim(&etsiObj.errorInfo, CheckAsn1Reencoding(reflect.ValueOf(statemWithoutInfo).Interface(), raw.Raw,
		"invalid format of ETSI Complicance statement"))
	return etsiObj
}
