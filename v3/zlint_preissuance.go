/*
 * ZLint Copyright 2022 Regents of the University of Michigan
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

// Used to check parsed info from certificate for compliance

package zlint

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

type tbsCertificatePartial struct {
	Version            int `asn1:"optional,explicit,default:0,tag:0"`
	SerialNumber       *big.Int
	SignatureAlgorithm pkix.AlgorithmIdentifier
}

type signed struct {
	ToBeSigned         asn1.RawValue
	SignatureAlgorithm pkix.AlgorithmIdentifier
	SignatureValue     asn1.BitString
}

// LintTBSCertificate runs all registered lints on rawTBSCertificate using default options,
// producing a ResultSet.
//
// Using LintTBSCertificate(rawTBSCertificate) is equivalent to calling LintCertificateEx(rawTBSCertificate, nil).
func LintTBSCertificate(rawTBSCertificate []byte) *ResultSet {
	// Run all lints from the global registry
	return LintTBSCertificateEx(rawTBSCertificate, nil)
}

// LintTBSCertificateEx runs lints from the provided registry on rawTBSCertificate producing
// a ResultSet. Providing an explicit registry allows the caller to filter the
// lints that will be run. (See lint.Registry.Filter())
//
// If registry is nil then the global registry of all lints is used and this
// function is equivalent to calling LintTBSCertificate(rawTBSCertificate).
func LintTBSCertificateEx(rawTBSCertificate []byte, registry lint.Registry) *ResultSet {
	// Decode enough of the TBSCertificate to discover the signature algorithm
	var tbs tbsCertificatePartial
	if _, err := asn1.Unmarshal(rawTBSCertificate, &tbs); err != nil {
		return nil
	}

	// Package the TBSCertificate in a dummy, yet syntactically valid, X.509 certificate that LintCertificateEx will be able to parse
	dummy := signed{
		ToBeSigned:         asn1.RawValue{FullBytes: rawTBSCertificate},
		SignatureAlgorithm: tbs.SignatureAlgorithm,
	}

	// For ECDSA signature algorithms, produce a dummy signature that will satisfy the e_mp_ecdsa_signature_encoding_correct lint
	if tbs.SignatureAlgorithm.Algorithm.Equal(asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 2}) { // ecdsa-with-SHA256
		dummy.SignatureValue.Bytes = []byte{
			0x30, 0x46, 0x02, 0x21, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67,
			0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x02, 0x21, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01,
			0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89,
			0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01,
		}
	} else if tbs.SignatureAlgorithm.Algorithm.Equal(asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 3}) { // ecdsa-with-SHA384
		dummy.SignatureValue.Bytes = []byte{
			0x30, 0x66, 0x02, 0x31, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67,
			0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67,
			0x89, 0xab, 0xcd, 0xef, 0x01, 0x02, 0x31, 0x01, 0x23, 0x45, 0x67, 0x89,
			0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01,
			0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89,
			0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01,
			0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01,
		}
	}

	// DER encode the dummy certificate, then decode it again into an object that can be passed to LintCertificateEx
	if certDER, err := asn1.Marshal(dummy); err != nil {
		return nil
	} else if c, err := x509.ParseCertificate(certDER); err != nil {
		return nil
	} else {
		return LintCertificateEx(c, registry)
	}
}
