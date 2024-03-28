package cabf_smime_br

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

func TestMailboxValidatedEnforceSubjectFieldRestrictions(t *testing.T) {
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
			Name:           "pass - mailbox validated, multipurpose with commonName",
			InputFilename:  "smime/mailboxValidatedMultipurposeWithCommonName.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - mailbox validated, strict with commonName",
			InputFilename:  "smime/mailboxValidatedStrictWithCommonName.pem",
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
			Name:            "error - certificate with countryName",
			InputFilename:   "smime/mailboxValidatedLegacyWithCountryName.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "subject DN contains forbidden field: subject:countryName (2.5.4.6)",
		},
		{
			Name:            "error - certificate containing nonsense subject field (1.2.3.4.5.6.7.8.9.0)",
			InputFilename:   "smime/mailboxValidatedMultipurposeWithNonsenseSubjectField.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "subject DN contains forbidden field: 1.2.3.4.5.6.7.8.9.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mailbox_validated_enforce_subject_field_restrictions", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}

			if tc.ExpectedDetails != "" && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %s, was %s", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
