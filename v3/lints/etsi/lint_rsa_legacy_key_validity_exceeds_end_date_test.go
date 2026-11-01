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

func TestRsaLegacyKeyValidityExceedsEndDate(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{"Error - 2048-bit RSA issued before sunset, validity beyond 2028-12-31", "etsi/rsa2048QcBeforeSunsetLongValidity.pem", lint.Error, "Certificate uses a 2048-bit RSA public key and was issued on or before 2026-12-31, so its validity period is required to end no later than 2028-12-31"},
		{"Pass - 2048-bit RSA issued before sunset, validity ends 2028-12-31", "etsi/rsa2048QcBeforeSunsetShortValidity.pem", lint.Pass, ""},
		{"NA - 1024-bit RSA below legacy band", "etsi/rsa1024QcBeforeSunsetLongValidity.pem", lint.NA, ""},
		{"NA - 3072-bit RSA", "etsi/rsa3072QcBeforeSunsetLongValidity.pem", lint.NA, ""},
		{"NA - only a non-QcCompliance qcStatement (QcType)", "etsi/rsa2048QcTypeOnlyBeforeSunsetLongValidity.pem", lint.NA, ""},
		{"Error - EN 319 411-1 OVCP policy, no qcStatements", "etsi/rsa2048OvcpPolicyBeforeSunsetLongValidity.pem", lint.Error, "Certificate uses a 2048-bit RSA public key and was issued on or before 2026-12-31, so its validity period is required to end no later than 2028-12-31"},
		{"NE - issued after sunset", "etsi/rsa2048QcAfterSunset.pem", lint.NE, ""},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_etsi_rsa_legacy_key_validity_period_exceeds_end_date", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %v was %v", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
