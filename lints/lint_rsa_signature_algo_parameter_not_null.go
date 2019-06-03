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

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/crypto/cryptobyte"
	asn1_cryptobyte "golang.org/x/crypto/cryptobyte/asn1"
)

type rsaEncryptionParamNotNULL struct{}

// byte representation of pkix.Algorithm with OID rsaEncryption and Parameters asn1.NULL, includes BITSTRING tag for public key
var expectedSPKIAlgoBytes = []byte{0x30, 0xd, 0x6, 0x9, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0xd, 0x1, 0x1, 0x1, 0x5, 0x0, 0x03}

func (l *rsaEncryptionParamNotNULL) Initialize() error {
	return nil
}

func (l *rsaEncryptionParamNotNULL) CheckApplies(c *x509.Certificate) bool {
	// explicitly check for util.OidRSAEncryption, as RSA-PSS or RSA-OAEP certificates might be classified with c.PublicKeyAlgorithm = RSA
	return c.PublicKeyAlgorithmOID.Equal(util.OidRSAEncryption)
}

func (l *rsaEncryptionParamNotNULL) Execute(c *x509.Certificate) *LintResult {
	// Try byte comparison for Algorithm
	// Determine offset (SEQUENCE + length) first. Notably 1024 encodes length in 2 bytes, 2048 and up in 3 bytes
	offset := getOffset(c.RawSubjectPublicKeyInfo)
	if bytes.Compare(c.RawSubjectPublicKeyInfo[offset:len(expectedSPKIAlgoBytes)+offset], expectedSPKIAlgoBytes) == 0 {
		return &LintResult{Status: Pass}
	}

	input := cryptobyte.String(c.RawSubjectPublicKeyInfo)

	var publicKeyInfo cryptobyte.String
	if !input.ReadASN1(&publicKeyInfo, asn1_cryptobyte.SEQUENCE) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey"}
	}

	var algorithm cryptobyte.String
	if !publicKeyInfo.ReadASN1(&algorithm, asn1_cryptobyte.SEQUENCE) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey algorithm"}
	}

	if !algorithm.SkipASN1(asn1_cryptobyte.OBJECT_IDENTIFIER) {
		return &LintResult{Status: Fatal, Details: "error reading pkixPublicKey algorithm OID"}
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

	return &LintResult{Status: Fatal, Details: "certificate rsa algorithm identifier appears correct, but didn't match byte-wise comparison"}
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

func getOffset(asn1Sequence []byte) int {
	if len(asn1Sequence) < 2 {
		return 0
	}

	if asn1Sequence[1]&0x80 == 0 {
		// short form length encoding in 1 octet
		return 2 // 1 tag octet +1 length encoding octet
	}

	// long form length encoding
	length := int(asn1Sequence[1] & 0x7f)
	return length + 2 // +1 for the tag and +1 for the first length encoding octet
}
