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

func TestCrlMldsaSigAlgParamsNotAbsent(t *testing.T) {
	testCases := []struct {
		name            string
		filename        string
		expectedStatus  lint.LintStatus
		expectedDetails string
	}{
		{
			name:           "non-mldsa CRL",
			filename:       "crlEcdsaP256SHA256.pem",
			expectedStatus: lint.NA,
		},
		{
			name:           "mldsa-65 CRL, parameters absent",
			filename:       "crlMLDSA65.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:            "mldsa-65 CRL, NULL parameters",
			filename:        "crlMLDSA65SigAlgNullParams.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "tbsCertList.signature algorithm identifier has parameters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := test.TestRevocationListLint(t, "e_crl_mldsa_sig_alg_params_not_absent", tc.filename)
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v - details: %s", tc.expectedStatus, result.Status, result.Details)
			}
			if result.Details != tc.expectedDetails {
				t.Errorf("expected details %q was %q", tc.expectedDetails, result.Details)
			}
		})
	}
}
