package lints

/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

/*******************************************************************************************************
"RFC5280: RFC 4055, Section 1.2"
RSA: Encoded algorithm identifier MUST have NULL parameters.
*******************************************************************************************************/

import (
	"bytes"
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/crypto/cryptobyte"
	asn1_cryptobyte "golang.org/x/crypto/cryptobyte/asn1"
)

type rsaEncryptionParamNotNULL struct{}

// byte representation of pkix.Algorithm with OID rsaEncryption and Parameters asn1.NULL
var expectedSPKIAlgoBytes = []byte{0x6, 0x9, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0xd, 0x1, 0x1, 0x1, 0x5, 0x0}

func (l *rsaEncryptionParamNotNULL) Initialize() error {
	return nil
}

func (l *rsaEncryptionParamNotNULL) CheckApplies(c *x509.Certificate) bool {
	// explicitly check for util.OidRSAEncryption, as RSA-PSS or RSA-OAEP certificates might be classified with c.PublicKeyAlgorithm = RSA
	return c.PublicKeyAlgorithmOID.Equal(util.OidRSAEncryption)
}

func (l *rsaEncryptionParamNotNULL) Execute(c *x509.Certificate) *LintResult {
	input := cryptobyte.String(c.RawSubjectPublicKeyInfo)

	var publicKeyInfo cryptobyte.String
	if !input.ReadASN1(&publicKeyInfo, asn1_cryptobyte.SEQUENCE) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey"}
	}

	var algorithm cryptobyte.String
	if !publicKeyInfo.ReadASN1(&algorithm, asn1_cryptobyte.SEQUENCE) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey algorithm"}
	}

	// byte comparison of algorithm sequence and checking no trailing data is present
	var algorithmBytes []byte
	if algorithm.ReadBytes(&algorithmBytes, len(expectedSPKIAlgoBytes)) {
		if bytes.Compare(algorithmBytes, expectedSPKIAlgoBytes) == 0 && algorithm.Empty() {
			return &LintResult{Status: Pass}
		}
	}

	// re-parse to get an error message detailing what did not match in the byte comparison
	input = cryptobyte.String(c.RawSubjectPublicKeyInfo)

	if !input.ReadASN1(&publicKeyInfo, asn1_cryptobyte.SEQUENCE) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey"}
	}

	if !publicKeyInfo.ReadASN1(&algorithm, asn1_cryptobyte.SEQUENCE) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey algorithm"}
	}

	encryptionOID := asn1.ObjectIdentifier{}
	if !algorithm.ReadASN1ObjectIdentifier(&encryptionOID) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey algorithm OID"}
	}

	if !encryptionOID.Equal(util.OidRSAEncryption) {
		return &LintResult{Status: Fatal, Details: "certificate pkixPublicKey algorithm OID is not rsaEncryption"}
	}

	if algorithm.Empty() {
		return &LintResult{Status: Error, Details: "certificate contains RSA public key algorithm identifier missing required NULL parameter"}
	}

	var nullValue cryptobyte.String
	if !algorithm.ReadASN1(&nullValue, asn1_cryptobyte.NULL) {
		return &LintResult{Status: Error, Details: "certificate contains RSA public key algorithm identifier with non-NULL parameter"}
	}

	if len(nullValue) != 0 {
		return &LintResult{Status: Error, Details: "certificate contains RSA public key algorithm identifier with NULL parameter containing trailing data"}
	}

	// ensure algorithm is empty and no trailing data is present
	if !algorithm.Empty() {
		return &LintResult{Status: Error, Details: "certificate contains RSA public key algorithm identifier with trailing data"}
	}

	return &LintResult{Status: Fatal, Details: "certificate rsa algorithm appears correct, but didn't match byte-wise comparison"}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_rsa_encryption_parameter_not_null",
		Description:   "RSA: Encoded algorithm identifier MUST have NULL parameters",
		Citation:      "RFC 4055, Section 1.2",
		Source:        RFC5280, // RFC4055 is referenced in RFC5280, Section 1
		EffectiveDate: util.RFC5280Date,
		Lint:          &rsaEncryptionParamNotNULL{},
	})
}
