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

package lints

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rsaAlgIDNullParams struct{}

// Initialize is a NOP.
func (l *rsaAlgIDNullParams) Initialize() error {
	return nil
}

// CheckApplies returns true for certificates with an RSA public key.
func (l *rsaAlgIDNullParams) CheckApplies(c *x509.Certificate) bool {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA
}

// Execute returns a fatal lint result if the provided certificate encodes the
// RSA algorithm identifier SEQUENCE without the required NULL parameters.
//
// The baseline requirements mandate that certificates obey RFC 5280, which in
// turn defers to RFC 3279 for the encoding of the RSA algorithm identifier. RFC
// 3279 Section 2.3.1 "RSA keys" says:
//
//   The rsaEncryption OID is intended to be used in the algorithm field
//   of a value of type AlgorithmIdentifier.  The parameters field MUST
//   have ASN.1 type NULL for this algorithm identifier.
//
// This lint verifies that the NULL parameters MUST of this section is met.
func (l *rsaAlgIDNullParams) Execute(c *x509.Certificate) *LintResult {
	raw := c.SubjectAndKey().RawSubjectPublicKeyInfo

	var pkixPublicKey struct {
		Algo      pkix.AlgorithmIdentifier
		BitString asn1.BitString
	}

	if rest, err := asn1.Unmarshal(raw, &pkixPublicKey); err != nil {
		return &LintResult{
			Status:  Error,
			Details: fmt.Sprintf("Error unmarshaling raw pkixPublicKey: %v", err),
		}
	} else if len(rest) != 0 {
		return &LintResult{
			Status:  Error,
			Details: fmt.Sprintf("Trailing data after pkixPublicKey: %v", err),
		}
	}

	params := pkixPublicKey.Algo.Parameters

	if params.Tag != asn1.TagNull {
		return &LintResult{
			Status:  Fatal,
			Details: "certificate contains RSA public key algorithm identifier missing required NULL parameters",
		}
	}

	if !bytes.Equal(params.Bytes, []uint8{}) {
		return &LintResult{
			Status:  Fatal,
			Details: "certificate contains RSA public key algorithm identifier with non-NULL parameter bytes",
		}
	}

	return &LintResult{
		Status: Pass,
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "rsa_algid_null_params",
		Description:   "RSA: Encoded algorithm identifier MUST have NULL parameters",
		Citation:      "RFC 3279, Section 2.3.1",
		Source:        CABFBaselineRequirements, // BRs -> RFC 5280 -> RFC 3279
		EffectiveDate: util.RFC5280Date,
		Lint:          &rsaAlgIDNullParams{},
	})
}
