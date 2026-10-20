package etsi

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

func TestRsaKeySizeProhibitedAfterSunset(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{"Error - 2048-bit RSA QC cert issued after sunset", "etsi/rsa2048QcAfterSunset.pem", lint.Error, "Certificate was issued with a 2048-bit RSA public key, which is less than the 3000-bit minimum required for certificates issued after 2026-12-31"},
		{"Pass - 3072-bit RSA QC cert issued after sunset", "etsi/rsa3072QcAfterSunset.pem", lint.Pass, ""},
		{"NA - no qcStatements or ETSI policy", "etsi/rsa2048NoQcAfterSunset.pem", lint.NA, ""},
		{"NA - only a non-QcCompliance qcStatement (QcType)", "etsi/rsa2048QcTypeOnlyAfterSunset.pem", lint.NA, ""},
		{"Error - EN 319 411-1 OVCP policy, no qcStatements", "etsi/rsa2048OvcpPolicyAfterSunset.pem", lint.Error, "Certificate was issued with a 2048-bit RSA public key, which is less than the 3000-bit minimum required for certificates issued after 2026-12-31"},
		{"NA - ECDSA key", "etsi/ecQcAfterSunset.pem", lint.NA, ""},
		{"NE - issued before sunset", "etsi/rsa2048QcBeforeSunsetLongValidity.pem", lint.NE, ""},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_etsi_rsa_key_size_prohibited_after_sunset_date", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %v was %v", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
