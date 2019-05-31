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
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/crypto/cryptobyte"
	asn1_cryptobyte "golang.org/x/crypto/cryptobyte/asn1"
)

var rsaEncryptionOID = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}

type rsaEncryptionParamNotNULL struct{}

func (l *rsaEncryptionParamNotNULL) Initialize() error {
	return nil
}

func (l *rsaEncryptionParamNotNULL) CheckApplies(c *x509.Certificate) bool {
	// explicitly check for rsaEncryptionOID, as RSA-PSS or RSA-OAEP certificates might be classified with c.PublicKeyAlgorithm = RSA
	return c.PublicKeyAlgorithmOID.Equal(rsaEncryptionOID)
}

func (l *rsaEncryptionParamNotNULL) Execute(c *x509.Certificate) *LintResult {
	var input = cryptobyte.String(c.RawSubjectPublicKeyInfo)

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

	if !algorithm.PeekASN1Tag(asn1_cryptobyte.NULL) {
		return &LintResult{Status: Error, Details: "certificate contains RSA public key algorithm identifier with non-NULL parameter"}
	}

	return &LintResult{Status: Pass}
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
