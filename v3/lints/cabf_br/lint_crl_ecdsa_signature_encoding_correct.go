package cabf_br

/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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
	"encoding/hex"
	"fmt"

	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type crlEcdsaSignatureAidEncoding struct{}

/*
*************************************************************************
BRs: 7.1.3.2 Signature AlgorithmIdentifier
All objects signed by a CA Private Key MUST conform to these requirements on the
use of the AlgorithmIdentifier or AlgorithmIdentifier-derived type in the
context of signatures.

BRs: 7.1.3.2.2 ECDSA
The CA SHALL use the appropriate signature algorithm and encoding based upon the
signing key used.
**************************************************************************
*/
func init() {
	lint.RegisterRevocationListLint(&lint.RevocationListLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_crl_ecdsa_signature_encoding_correct",
			Description:   "The CA SHALL use the appropriate signature algorithm and encoding based upon the signing key used.",
			Citation:      "BRs: 7.1.3.2.2",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABFBRs_1_7_1_Date,
		},
		Lint: NewCrlEcdsaSignatureAidEncoding,
	})
}

func NewCrlEcdsaSignatureAidEncoding() lint.RevocationListLintInterface {
	return &crlEcdsaSignatureAidEncoding{}
}

func (l *crlEcdsaSignatureAidEncoding) CheckApplies(r *x509.RevocationList) bool {
	return r.SignatureAlgorithm == x509.ECDSAWithSHA1 ||
		r.SignatureAlgorithm == x509.ECDSAWithSHA256 ||
		r.SignatureAlgorithm == x509.ECDSAWithSHA384 ||
		r.SignatureAlgorithm == x509.ECDSAWithSHA512 ||
		isSHA224WithECDSA(r)
}
func isSHA224WithECDSA(r *x509.RevocationList) bool {
	encoded, err := util.GetSignatureAlgorithmInTBSCertListEncoded(r)
	if err != nil {
		return false
	}
	var aid struct {
		Algorithm  asn1.ObjectIdentifier
		Parameters asn1.RawValue `asn1:"optional"`
	}
	if _, err := asn1.Unmarshal(encoded, &aid); err != nil {
		return false
	}
	return aid.Algorithm.Equal(util.OidSignatureSHA224withECDSA)
}

//nolint:nestif
func (l *crlEcdsaSignatureAidEncoding) Execute(r *x509.RevocationList) *lint.LintResult {
	signature := r.Signature
	signatureSize := len(signature)

	encoded, err := util.GetSignatureAlgorithmInTBSCertListEncoded(r)
	if err != nil {
		return &lint.LintResult{Status: lint.Error, Details: err.Error()}
	}

	const maxP256SigByteLen = 72
	const maxP384SigByteLen = 104
	const maxP521SigByteLen = 140

	if signatureSize <= maxP256SigByteLen {
		expectedEncoding := []byte{0x30, 0x0a, 0x06, 0x08, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x04, 0x03, 0x02}
		if bytes.Equal(encoded, expectedEncoding) {
			return &lint.LintResult{Status: lint.Pass}
		}
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "Encoding of signature algorithm does not match signing key on P-256 curve. Got the unsupported " + hex.EncodeToString(encoded),
		}
	} else if signatureSize <= maxP384SigByteLen {
		expectedEncoding := []byte{0x30, 0x0a, 0x06, 0x08, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x04, 0x03, 0x03}
		if bytes.Equal(encoded, expectedEncoding) {
			return &lint.LintResult{Status: lint.Pass}
		}
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "Encoding of signature algorithm does not match signing key on P-384 curve. Got the unsupported " + hex.EncodeToString(encoded),
		}
	} else if signatureSize <= maxP521SigByteLen {
		expectedEncoding := []byte{0x30, 0x0a, 0x06, 0x08, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x04, 0x03, 0x04}
		if bytes.Equal(encoded, expectedEncoding) {
			return &lint.LintResult{Status: lint.Pass}
		}
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "Encoding of signature algorithm does not match signing key on P-521 curve. Got the unsupported " + hex.EncodeToString(encoded),
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: fmt.Sprintf("Encoding of signature algorithm does not match signing key. Got signature length %v", signatureSize),
	}
}
