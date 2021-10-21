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

package rfc

import (
	"crypto/x509/pkix"
	"fmt"
	"strings"
	"unsafe"

	"github.com/zmap/zcrypto/x509"
	asn1parser "github.com/zmap/zlint/v3/ext/asn1c"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_asn1_syntax",
		Description:   "Certificate must be structured according to the ASN.1 modules",
		Citation:      "RFC5280 and other RFCs",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          func() lint.LintInterface { return &asn1Syntax{} },
	})
}

var ext2pdu = map[string]string{
	"2.5.29.9":                 "SubjectDirectoryAttributes",
	"2.5.29.14":                "SubjectKeyIdentifier",
	"2.5.29.15":                "KeyUsage",
	"2.5.29.16":                "PrivateKeyUsagePeriod",
	"2.5.29.17":                "SubjectAltName",
	"2.5.29.18":                "IssuerAltName",
	"2.5.29.19":                "BasicConstraints",
	"2.5.29.30":                "NameConstraints",
	"2.5.29.31":                "CRLDistributionPoints",
	"2.5.29.32":                "CertificatePolicies",
	"2.5.29.33":                "PolicyMappings",
	"2.5.29.35":                "AuthorityKeyIdentifier",
	"2.5.29.37":                "ExtKeyUsageSyntax",
	"2.5.29.46":                "FreshestCRL",
	"2.5.29.54":                "InhibitAnyPolicy",
	"2.23.140.1.31":            "TorServiceDescriptorSyntax",
	"2.23.140.3.1":             "CABFOrganizationIdentifier",
	"1.2.250.1.86.1.1.1":       "OcspNoCheckValue",
	"1.2.840.113549.1.9.15":    "SMIMECapabilities",
	"1.3.6.1.4.1.11129.2.1.22": "CanSignHttpExchanges",
	"1.3.6.1.4.1.11129.2.4.2":  "SignedCertificateTimestampList",
	"1.3.6.1.4.1.11129.2.4.3":  "CertificateTransparencyPoison",
	"1.3.6.1.5.5.7.1.1":        "AuthorityInfoAccessSyntax",
	"1.3.6.1.5.5.7.1.3":        "QCStatements",
	"1.3.6.1.5.5.7.1.11":       "SubjectInfoAccessSyntax",
	"1.3.6.1.5.5.7.1.12":       "LogotypeExtn",
	"1.3.6.1.5.5.7.1.24":       "Features",
	"1.3.6.1.5.5.7.48.1.5":     "OcspNoCheckValue",
}

var extIgnored = map[string]bool{
	"1.2.250.1.86.1.1.1":           true, // Certinomis
	"1.2.840.113533.7.65.0":        true, // Entrust EntrustVersionInfo
	"1.2.840.113583.1.1.9.1":       true, // Adobe revocationInfoArchival
	"1.2.840.113583.1.1.9.2":       true, // Adobe revocationInfoArchival
	"1.2.840.113635.100.6.27.7.1":  true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.27.7.2":  true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.27.11.1": true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.27.11.2": true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.27.15.1": true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.27.15.2": true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.27.21.2": true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.27.22.2": true, // Apple Certificate Extensions
	"1.2.840.113635.100.6.48.1":    true, // Apple Certificate Extensions
	"1.3.6.1.4.1.9.9.86.1.1.1.1.2": true, // Cisco cdspCardState
	"1.3.6.1.4.1.311.2.1.27":       true, // Microsoft SPC_FINANCIAL_CRITERIA_OBJID
	"1.3.6.1.4.1.311.20.2":         true, // Microsoft szOID_ENROLL_CERTTYPE_EXTENSION
	"1.3.6.1.4.1.311.21.1":         true, // Microsoft szOID_CERTSRV_CA_VERSION
	"1.3.6.1.4.1.311.21.2":         true, // Microsoft szOID_CERTSRV_PREVIOUS_CERT_HASH
	"1.3.6.1.4.1.311.21.7":         true, // Microsoft szOID_CERTIFICATE_TEMPLATE
	"1.3.6.1.4.1.311.21.10":        true, // Microsoft szOID_APPLICATION_CERT_POLICIES
	"1.3.101.77":                   true, // IETF draft - reductedSubjectAltName
	"2.16.840.1.113730.1.1":        true, // Netscape netscapeCertType
	"2.16.840.1.113730.1.4":        true, // Netscape netscapeCaRevocationUrl
}

var attr2pdu = map[string]string{
	"2.5.4.3":                    "X520CommonName",
	"2.5.4.4":                    "X520name",
	"2.5.4.5":                    "X520SerialNumber",
	"2.5.4.6":                    "X520countryName",
	"2.5.4.7":                    "X520LocalityName",
	"2.5.4.8":                    "X520StateOrProvinceName",
	"2.5.4.9":                    "StreetAddress",
	"2.5.4.10":                   "X520OrganizationName",
	"2.5.4.11":                   "X520OrganizationalUnitName",
	"2.5.4.12":                   "X520Title",
	"2.5.4.13":                   "Description",
	"2.5.4.15":                   "BusinessCategory",
	"2.5.4.16":                   "PostalAddress",
	"2.5.4.17":                   "PostalCode",
	"2.5.4.18":                   "PostOfficeBox",
	"2.5.4.41":                   "X520name",
	"2.5.4.42":                   "X520name",
	"2.5.4.43":                   "X520name",
	"2.5.4.44":                   "X520name",
	"2.5.4.46":                   "X520dnQualifier",
	"2.5.4.65":                   "X520Pseudonym",
	"2.5.4.97":                   "OrganizationIdentifier",
	"0.9.2342.19200300.100.1.1":  "UserId",
	"0.9.2342.19200300.100.1.25": "DomainComponent",
	"1.2.840.113549.1.9.1":       "EmailAddress",
	"1.2.840.113549.1.9.2":       "UnstructuredName",
	"1.3.6.1.4.1.311.60.2.1.1":   "JurisdictionLocalityName",
	"1.3.6.1.4.1.311.60.2.1.2":   "JurisdictionStateOrProvinceName",
	"1.3.6.1.4.1.311.60.2.1.3":   "JurisdictionCountryName",
}

var attrIgnored = map[string]bool{
	"1.3.6.1.4.1.17326.30.3": true, // Camerfirma company ID number
	"1.3.6.1.4.1.18838.1.1":  true, // ACCV unknown attribute
	"1.3.6.1.4.1.519.1":      true, // DUNS Number
	"1.3.6.1.4.1.52266.1":    true, // LEI
	"1.3.6.1.4.1.53087.1.2":  true, // BIMI Group
	"1.3.6.1.4.1.53087.1.3":  true, // BIMI Group
	"1.3.6.1.4.1.53087.1.4":  true, // BIMI Group
}

var algo2param = map[string]string{
	// RSA
	"1.2.840.113549.1.1.1":  "NullParameters",
	"1.2.840.113549.1.1.7":  "RSAES-OAEP-params",
	"1.2.840.113549.1.1.8":  "HashAlgorithm",
	"1.2.840.113549.1.1.9":  "OctetStringParameters",
	"1.2.840.113549.1.1.10": "RSASSA-PSS-params",
	"1.2.840.113549.1.1.11": "NullParameters",
	"1.2.840.113549.1.1.12": "NullParameters",
	"1.2.840.113549.1.1.13": "NullParameters",
	"1.2.840.113549.1.1.14": "NullParameters",
	// DSA
	"1.2.840.10040.4": "DSS-Parms",
	// DH
	"1.2.840.10046.2.1": "DomainParameters",
	// ECDSA
	"1.2.840.10045.2.1": "ECParameters",
	// ECDH
	"1.3.132.1.12": "ECParameters",
	// ECMQV
	"1.3.132.1.13": "ECParameters",
}

var algo2key = map[string]string{
	// RSA
	"1.2.840.113549.1.1.1":  "RSAPublicKey",
	"1.2.840.113549.1.1.7":  "RSAPublicKey",
	"1.2.840.113549.1.1.8":  "RSAPublicKey",
	"1.2.840.113549.1.1.9":  "RSAPublicKey",
	"1.2.840.113549.1.1.10": "RSAPublicKey",
	"1.2.840.113549.1.1.11": "RSAPublicKey",
	"1.2.840.113549.1.1.12": "RSAPublicKey",
	"1.2.840.113549.1.1.13": "RSAPublicKey",
	"1.2.840.113549.1.1.14": "RSAPublicKey",
	// DSA
	"1.2.840.10040.4": "DSAPublicKey",
	// DH
	"1.2.840.10046.2.1": "DHPublicKey",
}

var algo2keybs = map[string]bool{
	// ECDSA
	"1.2.840.10045.2.1": true,
	// ECDH
	"1.3.132.1.12": true,
	// ECMQV
	"1.3.132.1.13": true,
}

type asn1Syntax struct{}

func NewAsn1Syntax() lint.LintInterface {
	return &asn1Syntax{}
}

func (l *asn1Syntax) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *asn1Syntax) Execute(c *x509.Certificate) *lint.LintResult {
	pdu, err := asn1parser.DecodePdu("Certificate", c.Raw)
	if err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "incorrectly encoded certificate found",
		}
	}
	defer asn1parser.FreePdu("Certificate", pdu)

	if err = asn1parser.CheckConstraints("Certificate", pdu); err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "incorrectly encoded certificate found",
		}
	}

	var errs, warns []string

	invExt, unkExt, err := lintExtensions(pdu)
	if err != nil {
		errs = append(errs, "cannot decode extensions")
	}
	if len(invExt) > 0 {
		errs = append(errs, fmt.Sprintf("invalid extension(s) %s found", strings.Join(invExt, ",")))
	}
	if len(unkExt) > 0 {
		warns = append(warns, fmt.Sprintf("unknown extension(s) %s found", strings.Join(unkExt, ",")))
	}

	invSubj, unkSubj, err := lintSubject(pdu)
	if err != nil {
		errs = append(errs, "cannot decode subject")
	}
	if len(invSubj) > 0 {
		errs = append(errs, fmt.Sprintf("invalid attribute(s) %s found in subject", strings.Join(invSubj, ",")))
	}
	if len(unkSubj) > 0 {
		warns = append(warns, fmt.Sprintf("unknown attribute(s) %s found in subject", strings.Join(unkSubj, ",")))
	}

	invIssuer, unkIssuer, err := lintIssuer(pdu)
	if err != nil {
		errs = append(errs, "cannot decode issuer")
	}
	if len(invIssuer) > 0 {
		errs = append(errs, fmt.Sprintf("invalid attribute(s) %s found in issuer", strings.Join(invIssuer, ",")))
	}
	if len(unkIssuer) > 0 {
		warns = append(warns, fmt.Sprintf("unknown attribute(s) %s found in issuer", strings.Join(unkIssuer, ",")))
	}

	invSPKI, unkSPKI, err := lintSubjectPublicKeyIdentifier(pdu)
	if err != nil {
		errs = append(errs, "cannot decode SPKI")
	}
	if len(invSPKI) > 0 {
		errs = append(errs, strings.Join(invSPKI, ","))
	}
	if len(unkSPKI) > 0 {
		warns = append(warns, strings.Join(unkSPKI, ","))
	}

	if len(errs) > 0 {
		fmt.Printf("Error at certificate with serial %x\n", c.SerialNumber)
		return &lint.LintResult{
			Status:  lint.Error,
			Details: strings.Join(errs, ","),
		}
	}

	if len(warns) > 0 {
		fmt.Printf("Warning at certificate with serial %x\n", c.SerialNumber)
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: strings.Join(warns, ","),
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}

// verifyPdu will decode and check constraints for a specified PDU type.
func verifyPdu(name string, data []byte) bool {
	pdu, err := asn1parser.DecodePdu(name, data)
	if err != nil {
		return false
	}
	defer asn1parser.FreePdu(name, pdu)

	if err = asn1parser.CheckConstraints(name, pdu); err != nil {
		return false
	}

	return true
}

// verifyATVs will lint all AttributeTypeAndValue pairs and return all invalid
// and all unknown pairs found.
func verifyATVs(atvs []pkix.AttributeTypeAndValue) ([]string, []string) {
	var invalidPair, unknownPair []string

	for _, atv := range atvs {
		ign, ok := attrIgnored[atv.Type.String()]
		if ok && ign {
			continue
		}

		pduname, ok := attr2pdu[atv.Type.String()]
		if !ok {
			unknownPair = append(unknownPair, atv.Type.String())
		}

		if verifyPdu(pduname, atv.Value.([]byte)) == false {
			invalidPair = append(invalidPair, atv.Type.String())
		}
	}

	return invalidPair, unknownPair
}

// lintExtensions will lint all extensions and return all invalid and all
// unknown extensions found.
func lintExtensions(pdu unsafe.Pointer) ([]string, []string, error) {
	var invalidExt, unknownExt []string

	exts, err := asn1parser.GetCertExtensions(pdu)
	if err != nil {
		return nil, nil, err
	}

	for _, ext := range exts {
		ign, ok := extIgnored[ext.Id.String()]
		if ok && ign {
			continue
		}

		pduname, ok := ext2pdu[ext.Id.String()]
		if !ok {
			unknownExt = append(unknownExt, ext.Id.String())
			continue
		}

		if verifyPdu(pduname, ext.Value) == false {
			invalidExt = append(invalidExt, ext.Id.String())
		}
	}

	return invalidExt, unknownExt, nil
}

// lintSubject will lint all AttributeTypeAndValue pairs in the subject and
// return all invalid and all unknown pairs found.
func lintSubject(pdu unsafe.Pointer) ([]string, []string, error) {
	atvs, err := asn1parser.GetSubjectATV(pdu)
	if err != nil {
		return nil, nil, err
	}

	inv, unk := verifyATVs(atvs)
	return inv, unk, nil
}

// lintIssuer will lint all AttributeTypeAndValue pairs in the issuer and
// return all invalid and all unknown pairs found.
func lintIssuer(pdu unsafe.Pointer) ([]string, []string, error) {
	atvs, err := asn1parser.GetIssuerATV(pdu)
	if err != nil {
		return nil, nil, err
	}

	inv, unk := verifyATVs(atvs)
	return inv, unk, nil
}

// lintSubjectPublicKeyIdentifier will lint the SubjectPublicKeyIdentifier and
// return invalid or unknown parameters or key.
func lintSubjectPublicKeyIdentifier(pdu unsafe.Pointer) ([]string, []string, error) {
	var invalidSPKI, unknownSPKI []string

	algoid, params, key, err := asn1parser.GetSubjectPublicKeyInfo(pdu)
	if err != nil {
		return nil, nil, err
	}

	pduname, ok := algo2param[algoid.String()]
	if !ok {
		unknownSPKI = append(unknownSPKI, "unknown parameters for SPKI algorithm")
	} else {
		if verifyPdu(pduname, params) == false {
			invalidSPKI = append(invalidSPKI, "invalid parameters for SPKI algorithm")
		}
	}

	isBs, ok := algo2keybs[algoid.String()]
	if ok && isBs {
		return invalidSPKI, unknownSPKI, nil
	}

	pduname, ok = algo2key[algoid.String()]
	if !ok {
		unknownSPKI = append(unknownSPKI, "unknown key for SPKI algorithm")
	} else {
		if verifyPdu(pduname, key) == false {
			invalidSPKI = append(invalidSPKI, "invalid key for SPKI algorithm")
		}
	}

	return invalidSPKI, unknownSPKI, nil
}
