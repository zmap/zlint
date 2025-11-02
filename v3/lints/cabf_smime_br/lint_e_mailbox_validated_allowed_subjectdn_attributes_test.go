/*
 * ZLint Copyright 2025 Regents of the University of Michigan
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

package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestMBVSubjectAttributes(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Error - certificate is mailbox-validated and has organization in subject",
			InputFilename:  "smime/mailboxValidatedWithOrganizationInSubject.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Pass - certificate is mailbox-validated and has commonName in subject",
			InputFilename:  "smime/mailboxValidatedWithCommonNameInSubject.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Pass - certificate is mailbox-validated and has emailAddress in subject",
			InputFilename:  "smime/mailboxValidatedWithEmailAddressInSubject.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Pass - certificate is mailbox-validated and has serialNumber in subject",
			InputFilename:  "smime/mailboxValidatedWithSerialNumberInSubject.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - certificate is organization-validated",
			InputFilename:  "smime/organization_validated_with_matching_country.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mailbox_validated_allowed_subjectdn_attributes", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
