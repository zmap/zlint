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

func TestRsaAllowedKUCa(t *testing.T) {
	testCases := []struct {
		name            string
		filename        string
		expectedStatus  lint.LintStatus
		expectedDetails string
	}{
		{
			name:            "Certificate with EC key",
			filename:        "ecdsaP384.pem",
			expectedStatus:  lint.NA,
			expectedDetails: "",
		},
		{
			name:            "Subscriber certificate with RSA key and key usages digitalSignature and nonRepudiation",
			filename:        "eeWithRSAAllowedKeyUsage.pem",
			expectedStatus:  lint.NA,
			expectedDetails: "",
		},
		{
			name:            "CA certificate with RSA key and key usages digitalSignature and nonRepudiation older than 2002",
			filename:        "caWithRSAAllowedKeyUsageOld.pem",
			expectedStatus:  lint.NE,
			expectedDetails: "",
		},
		{
			name:            "CA certificate with RSA key and key usages digitalSignature, certificateSign, and crlSign",
			filename:        "caBasicConstCrit.pem",
			expectedStatus:  lint.Pass,
			expectedDetails: "",
		},
		{
			name:            "CA certificate with RSA key and key usages certificateSign and keyAgreement",
			filename:        "caWithRSADisallowedKeyUsage.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "CA certificate with an RSA key contains invalid key usage(s): KeyUsageKeyAgreement",
		},
		{
			name:            "CA certificate with RSA key and key usages certificateSign and keyEncipherment",
			filename:        "caWithRSAAndEnciphermentKeyUsage.pem",
			expectedStatus:  lint.Pass,
			expectedDetails: "",
		},
	}

	for _, tc := range testCases {
		result := test.TestLint("e_rsa_allowed_ku_ca", tc.filename)
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
