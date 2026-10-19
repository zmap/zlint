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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestMldsaKeyUsages(t *testing.T) {
	testCases := []struct {
		name            string
		filename        string
		expectedStatus  lint.LintStatus
		expectedDetails string
	}{
		{
			name:           "non-mldsa cert",
			filename:       "ecdsaP256.pem",
			expectedStatus: lint.NA,
		},
		{
			name:           "mldsa cert, key usage is absent",
			filename:       "mldsa44AbsentKU.pem",
			expectedStatus: lint.NA,
		},
		{
			name:           "mldsa cert, notBefore before RFC 9881",
			filename:       "mldsa44NotBeforeBeforeRFC9881.pem",
			expectedStatus: lint.NE,
		},
		{
			name:           "mldsa-44 cert, key usage is digitalSignature",
			filename:       "mldsa44.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:           "mldsa-65 CA cert, key usage is keyCertSign and cRLSign",
			filename:       "mldsa65CA.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:           "mldsa-87 cert, key usage is digitalSignature and nonRepudiation",
			filename:       "mldsa87ValidKUs.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:            "mldsa cert, key usage has keyEncipherment",
			filename:        "mldsa44KUIsKeyEnciphermentInvalidKU.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "Certificate contains invalid key usage(s): KeyUsageKeyEncipherment",
		},
		{
			name:            "mldsa cert, key usage is keyAgreement only",
			filename:        "mldsa44KUIsKeyAgreementOnlyInvalidKU.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "Certificate contains invalid key usage(s): KeyUsageKeyAgreement",
		},
		{
			name:            "mldsa cert, key usage has no bits set",
			filename:        "mldsa44KUEmptyInvalidKU.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "Certificate does not contain any of the key usages: KeyUsageDigitalSignature, KeyUsageContentCommitment, KeyUsageCertSign, KeyUsageCRLSign",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := test.TestLint("e_mldsa_key_usages", tc.filename)
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v - details: %s", tc.expectedStatus, result.Status, result.Details)
			}
			if result.Details != tc.expectedDetails {
				t.Errorf("expected details %q was %q", tc.expectedDetails, result.Details)
			}
		})
	}
}
