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

func TestOcspRsaSignatureAidEncoding(t *testing.T) {
	testCases := []struct {
		file string
		want lint.LintStatus
	}{
		{"ocspRsaSHA256", lint.Pass},
		{"ocspRsaSHA384", lint.Pass},
		{"ocspRsaSHA512", lint.Pass},
		{"ocspRsaPSSSHA256", lint.Pass},
		{"ocspRsaPSSSHA256Salt20", lint.Error},
		{"ocspRsaSHA224", lint.Error},
		// sha256WithRSAEncryption with the NULL parameters omitted (300b06092a864886f70d01010b).
		{"ocspRsaSHA256MissingNullParams", lint.Error},
		{"ocspRsaSHA1", lint.Error},
		{"ocspRsaSHA1ProducedBefore20220601", lint.Pass},
		{"ocspRsaSHA1ProducedOn20220601", lint.Error},
		{"ocspEcdsaP256SHA256", lint.NA},
	}
	for _, tc := range testCases {
		t.Run(tc.file, func(t *testing.T) {
			got := test.TestOCSPResponseLint(t, "e_ocsp_rsa_signature_encoding_correct", tc.file)
			if got.Status != tc.want {
				t.Errorf("expected %s, got %s: %s", tc.want, got.Status, got.Details)
			}
		})
	}
}
