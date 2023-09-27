package mozilla

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestECDSASignatureAidEncoding(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Standard ECC certificate with a P-256 key signed by a P-256 key using SHA256withECDSA",
			InputFilename:  "eccP256.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard ECC certificate with a P-384 key signed by a P-384 key using SHA384withECDSA",
			InputFilename:  "eccP384.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard ECC certificate signed by a P-384 key using SHA256withECDSA",
			InputFilename:  "eccSignedWithP384ButSHA256Signature.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Certificate signed with RSA",
			InputFilename:  "evAllGood.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Standard ECC certificate with a P-256 key signed by a P-256 key using SHA512withECDSA",
			InputFilename:  "eccSignedWithSHA512Signature.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Standard ECC certificate with a secp521r1 key signed by a secp521r1 key using SHA512withECDSA",
			InputFilename:  "eccWithSecp521r1KeySignedWithSHA512Signature.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mp_ecdsa_signature_encoding_correct", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
