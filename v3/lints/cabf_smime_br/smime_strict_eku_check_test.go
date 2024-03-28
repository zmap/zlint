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

func TestStrictEKUCheck(t *testing.T) {
	testCases := []struct {
		Name          string
		InputFilename string

		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - mailbox validated, strict with EmailProtectionEKU",
			InputFilename:  "smime/mailboxValidatedStrictWithCommonName.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "na - certificate without mailbox validated policy",
			InputFilename:  "smime/domainValidatedWithEmailCommonName.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "na - mailbox validated legacy certificate",
			InputFilename:  "smime/mailboxValidatedLegacyWithCommonName.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "ne - certificate with NotBefore before effective date of lint",
			InputFilename:  "smime/mailboxValidatedStrictMay2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "error - certificate with extra EKU",
			InputFilename:  "smime/individualValidatedStrictWithServerAuthEKU.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_smime_strict_eku_check", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
