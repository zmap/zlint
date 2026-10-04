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
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCrlEcdsaSignatureAidEncoding(t *testing.T) {
	testCases := []struct {
		file string
		want lint.LintStatus
	}{
		{"crlEcdsaP256SHA256.pem", lint.Pass},
		{"crlEcdsaP384SHA384.pem", lint.Pass},
		{"crlEcdsaP256SHA384.pem", lint.Error},
		{"crlEcdsaP384SHA256.pem", lint.Error},
		{"crlEcdsaP384SHA384NullParams.pem", lint.Error},
		{"crlEcdsaP256SHA224.pem", lint.Error},
		{"crlEcdsaP521SHA512.pem", lint.Pass},
		{"crlEcdsaP521SHA256.pem", lint.Error},
		// Signature length boundaries that select the expected curve.
		{"crlEcdsaP256SHA256Sig72.pem", lint.Pass},
		{"crlEcdsaP256SHA256Sig73.pem", lint.Error},
		{"crlEcdsaP384SHA384Sig105.pem", lint.Error},
		{"crlEcdsaP521SHA512Sig105.pem", lint.Pass},
		{"crlEcdsaSig141.pem", lint.Error},
		{"crlRsaSHA256.pem", lint.NA},
	}
	for _, tc := range testCases {
		t.Run(tc.file, func(t *testing.T) {
			got := test.TestRevocationListLint(t, "e_crl_ecdsa_signature_encoding_correct", tc.file)
			if got.Status != tc.want {
				t.Errorf("expected %s, got %s: %s", tc.want, got.Status, got.Details)
			}
		})
	}
}
