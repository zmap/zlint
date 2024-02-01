package rfc

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

func TestECDSAAllowedKU(t *testing.T) {
	testCases := []struct {
		name            string
		filename        string
		expectedStatus  lint.LintStatus
		expectedDetails string
	}{
		{
			name:           "non-ecdsa ee cert",
			filename:       "rsaKeyWithParameters.pem",
			expectedStatus: lint.NA,
		},
		{
			name:           "ecdsa ee cert, valid key usage, notBefore before RFC",
			filename:       "ecdsaP256ValidKUs.pem",
			expectedStatus: lint.NE,
		},
		{
			name:           "ecdsa ee cert, key usage is absent",
			filename:       "ecdsaP256AbsentKU.pem",
			expectedStatus: lint.NA,
		},
		{
			name:           "ecdsa ee cert, valid key usage",
			filename:       "ecdsaP256KUIsDigitalSignatureValidKU.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:            "ecdsa ee cert, invalid key usage",
			filename:        "ecdsaP256KUIsDataEnciphermentInvalidKU.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "Certificate contains invalid key usage(s): KeyUsageDataEncipherment",
		},
		{
			name:            "ecdsa ee cert, invalid key usage",
			filename:        "ecdsaP256KUIsKeyEnciphermentInvalidKU.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "Certificate contains invalid key usage(s): KeyUsageKeyEncipherment",
		},
		{
			name:            "ecdsa ee cert, invalid key usage",
			filename:        "ecdsaP256KUIsKeyEnciphermentAndDataEnciphermentInvalidKU.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "Certificate contains invalid key usage(s): KeyUsageDataEncipherment, KeyUsageKeyEncipherment",
		},
	}

	for _, tc := range testCases {
		result := test.TestLint("e_ecdsa_allowed_ku", tc.filename)
		if result.Status != tc.expectedStatus {
			t.Errorf("expected result %v. actual result was %v",
				tc.expectedStatus, result.Status)
		}
		if result.Details != tc.expectedDetails {
			t.Errorf("expected details %q. actual result was %q",
				tc.expectedDetails, result.Details)
		}
	}
}
