package cabf_br

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

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type AlgorithmObjectIdentifierEncoding struct {
	// Contains all specified byte sequences depending on the public key algorithm/type
	expectedRawAlgorithmIdentifiers map[string][]byte
}

/************************************************
This lint refers to CAB Baseline Requirements (Version 1.7.4) chapter 7.1.3.1, which defines the
required encodings of AlgorithmObjectIdentifiers inside a SubjectPublicKeyInfo field.

Section 7.1.3.1.2: When encoded, the AlgorithmIdentifier for RSA keys MUST be byte‐for‐byte
identical with the following hex‐encoded bytes: 300d06092a864886f70d0101010500

Section 7.1.3.1.2: When encoded, the AlgorithmIdentifier for ECDSA keys MUST be
byte‐for‐byte identical with the following hex‐encoded bytes:
For P‐256 keys: 301306072a8648ce3d020106082a8648ce3d030107
For P‐384 keys: 301006072a8648ce3d020106052b81040022
For P‐521 keys: 301006072a8648ce3d020106052b81040023
************************************************/
func init() {
	lint.RegisterLint(&lint.Lint{
		Name: "e_algorithm_identifier_improper_encoding",
		Description: "Encoded AlgorithmObjectIdentifier objects inside a SubjectPublicKeyInfo field " +
			"MUST comply with specified byte sequences.",
		Citation:      "BRs: 7.1.3.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABFBRs_1_7_1_Date,
		Lint:          NewAlgorithmObjectIdentifierEncoding,
	})
}

func NewAlgorithmObjectIdentifierEncoding() lint.LintInterface {
	expectedRawAlgorithmIdentifiers := make(map[string][]byte)
	expectedRawAlgorithmIdentifiers["RSA"] = []byte{0x30, 0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86,
		0xf7, 0x0d, 0x01, 0x01, 0x01, 0x05, 0x00}
	expectedRawAlgorithmIdentifiers["P-256"] = []byte{0x30, 0x13, 0x06, 0x07, 0x2a, 0x86, 0x48, 0xce,
		0x3d, 0x02, 0x01, 0x06, 0x08, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x03, 0x01, 0x07}
	expectedRawAlgorithmIdentifiers["P-384"] = []byte{0x30, 0x10, 0x06, 0x07, 0x2a, 0x86, 0x48, 0xce,
		0x3d, 0x02, 0x01, 0x06, 0x05, 0x2b, 0x81, 0x04, 0x00, 0x22}
	expectedRawAlgorithmIdentifiers["P-521"] = []byte{0x30, 0x10, 0x06, 0x07, 0x2a, 0x86, 0x48, 0xce,
		0x3d, 0x02, 0x01, 0x06, 0x05, 0x2b, 0x81, 0x04, 0x00, 0x23}

	return &AlgorithmObjectIdentifierEncoding{
		expectedRawAlgorithmIdentifiers: expectedRawAlgorithmIdentifiers,
	}
}

func (l *AlgorithmObjectIdentifierEncoding) CheckApplies(c *x509.Certificate) bool {
	// the requirement is only specified for RSA and ECDSA (P-256, P-384, P-521)
	if c.PublicKeyAlgorithm == x509.RSA {
		return true
	}
	if c.PublicKeyAlgorithm == x509.ECDSA {
		curveName := l.determinePublicKeyType(c)
		return curveName == "P-256" || curveName == "P-384" || curveName == "P-521"
	}
	return false
}

func (l *AlgorithmObjectIdentifierEncoding) Execute(c *x509.Certificate) *lint.LintResult {

	resolvedPublicKeyAlgorithm := l.determinePublicKeyType(c)

	rawAlgorithmIdentifier, err := util.GetPublicKeyAidEncoded(c)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal, Details: "error parsing SubjectPublicKeyInfo"}
	}

	expectedRawAlgorithmIdentifier, found := l.expectedRawAlgorithmIdentifiers[resolvedPublicKeyAlgorithm]
	if !found {
		return &lint.LintResult{Status: lint.Fatal, Details: "unexpected public key type"}
	}

	if bytes.Equal(rawAlgorithmIdentifier, expectedRawAlgorithmIdentifier) {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{
			Status: lint.Error,
			Details: fmt.Sprintf(
				"The encoded AlgorithmObjectIdentifier for %s inside the the SubjectPublicKeyInfo field is %q but the expected one is %q.",
				resolvedPublicKeyAlgorithm,
				hex.EncodeToString(rawAlgorithmIdentifier), hex.EncodeToString(expectedRawAlgorithmIdentifier)),
		}
	}
}

func (l *AlgorithmObjectIdentifierEncoding) determinePublicKeyType(c *x509.Certificate) string {

	switch c.PublicKeyAlgorithm {
	case x509.RSA:
		return "RSA"
	case x509.ECDSA:
		return l.determineCurveName(c.PublicKey)
	default:
		return ""
	}
}

func (l *AlgorithmObjectIdentifierEncoding) determineCurveName(pubKey interface{}) string {

	var ecPubKey *ecdsa.PublicKey
	switch keyType := pubKey.(type) {
	case *x509.AugmentedECDSA:
		ecPubKey = keyType.Pub
	case *ecdsa.PublicKey:
		ecPubKey = keyType
	}

	return ecPubKey.Curve.Params().Name
}
