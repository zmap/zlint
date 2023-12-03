package cabf_smime_br

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

func TestLegacyMultipurposeEKUCheck(t *testing.T) {
	testCases := []struct {
		Name          string
		InputFilename string

		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "pass - mailbox validated, legacy with commonName",
			InputFilename:  "smime/mailboxValidatedLegacyWithCommonName.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "na - certificate without mailbox validated policy",
			InputFilename:  "smime/domainValidatedWithEmailCommonName.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "ne - certificate with NotBefore before effective date of lint",
			InputFilename:  "smime/mailboxValidatedLegacyWithCommonNameMay2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:            "error - certificate without emailProtection EKU",
			InputFilename:   "smime/mailboxValidatedLegacyWithoutEmailProtectionEKU.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "id-kp-emailProtection SHALL be present",
		},
		{
			Name:            "error - certificate containing serverAuthEKU",
			InputFilename:   "smime/organizationValidatedMultipurposeWithServerAuthEKU.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "id-kp-serverAuth, id-kp-codeSigning, id-kp-timeStamping, and anyExtendedKeyUsage values SHALL NOT be present",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_smime_legacy_multipurpose_eku_check", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}

			if tc.ExpectedDetails != "" && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %s, was %s", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
