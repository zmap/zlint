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

func TestOcspEcdsaSignatureAidEncoding(t *testing.T) {
	testCases := []struct {
		file string
		want lint.LintStatus
	}{
		// Signed directly by the CA; the curve is inferred from the signature length.
		{"ocspEcdsaP256SHA256", lint.Pass},
		{"ocspEcdsaP384SHA384", lint.Pass},
		{"ocspEcdsaP521SHA512", lint.Pass},
		{"ocspEcdsaP256SHA384", lint.Error},
		{"ocspEcdsaP384SHA256", lint.Error},
		{"ocspEcdsaP521SHA256", lint.Error},
		{"ocspEcdsaP384SHA384NullParams", lint.Error},
		{"ocspEcdsaP256SHA224", lint.Error},
		{"ocspEcdsaP256SHA256Sig72", lint.Pass},
		{"ocspEcdsaP256SHA256Sig73", lint.Error},
		{"ocspEcdsaSig141", lint.Error},
		// Signed by a delegated responder; the curve is taken from the embedded certificate.
		{"ocspEcdsaResponderP384SHA384", lint.Pass},
		{"ocspEcdsaResponderP384SHA256", lint.Error},
		{"ocspEcdsaResponderP256SHA256", lint.Pass},
		{"ocspEcdsaResponderP256SHA384", lint.Error},
		{"ocspRsaSHA256", lint.NA},
	}
	for _, tc := range testCases {
		t.Run(tc.file, func(t *testing.T) {
			got := test.TestOCSPResponseLint(t, "e_ocsp_ecdsa_signature_encoding_correct", tc.file)
			if got.Status != tc.want {
				t.Errorf("expected %s, got %s: %s", tc.want, got.Status, got.Details)
			}
		})
	}
}
