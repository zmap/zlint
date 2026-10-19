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

func TestMldsaSigAlgParamsNotAbsent(t *testing.T) {
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
			name:           "mldsa cert, notBefore before RFC 9881",
			filename:       "mldsa44NotBeforeBeforeRFC9881.pem",
			expectedStatus: lint.NE,
		},
		{
			name:           "mldsa-44 cert, parameters absent",
			filename:       "mldsa44.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:           "mldsa-65 CA cert, parameters absent",
			filename:       "mldsa65CA.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:           "mldsa-87 cert, parameters absent",
			filename:       "mldsa87ValidKUs.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:            "NULL parameters in tbsCertificate.signature and signatureAlgorithm",
			filename:        "mldsa44SigAlgNullParams.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "tbsCertificate.signature algorithm identifier has parameters",
		},
		{
			name:            "NULL parameters in tbsCertificate.signature only",
			filename:        "mldsa44TBSSigAlgNullParams.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "tbsCertificate.signature algorithm identifier has parameters",
		},
		{
			name:            "NULL parameters in signatureAlgorithm only",
			filename:        "mldsa44OuterSigAlgNullParams.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "signatureAlgorithm algorithm identifier has parameters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := test.TestLint("e_mldsa_sig_alg_params_not_absent", tc.filename)
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v - details: %s", tc.expectedStatus, result.Status, result.Details)
			}
			if result.Details != tc.expectedDetails {
				t.Errorf("expected details %q was %q", tc.expectedDetails, result.Details)
			}
		})
	}
}
