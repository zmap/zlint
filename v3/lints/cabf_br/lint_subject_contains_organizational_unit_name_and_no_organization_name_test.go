package cabf_br

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

func TestSubjectContainsOrganizationalUnitNameButNoOrganizationName(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Subject does not contain organizational unit name",
			InputFilename:  "subjectDnWithoutOuEntry.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:            "Subject contains organizational unit name but no organization name",
			InputFilename:   "subjectDnWithOuEntryButWithoutOEntry.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "subject:organizationalUnitName is prohibited if subject:organizationName is absent",
		},
		{
			Name:           "Subject contains organizational unit and organization name but is issued before the effective date",
			InputFilename:  "subjectWithOandOUBeforeEffectiveDate.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Subject contains organizational unit and organization name and is issued after the effective date",
			InputFilename:  "subjectWithOandOUAfterEffectiveDate.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_subject_contains_organizational_unit_name_and_no_organization_name", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
