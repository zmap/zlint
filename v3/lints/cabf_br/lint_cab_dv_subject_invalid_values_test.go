package cabf_br

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

func TestNewDvSubjectInvalidValues(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "ne - DV with valid values in subjectDN, before SC62",
			InputFilename:  "domainValGoodSubject.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:            "error - DV with organization in subjectDN, on SC62",
			InputFilename:   "dvWithOrganization.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "DV certificate contains the invalid attribute type 2.5.4.10",
		},
		{
			Name:            "error - DV with serialNumber in subjectDN, on SC62",
			InputFilename:   "dvWithSerialNumber.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "DV certificate contains the invalid attribute type 2.5.4.5",
		},
		{
			Name:            "warn - DV with valid values in subjectDN, with CN, on SC62",
			InputFilename:   "dvWithCNAndCountry.pem",
			ExpectedResult:  lint.Warn,
			ExpectedDetails: "DV certificate contains a subject common name, this is not recommended",
		},
		{
			Name:           "pass - DV with valid values in subjectDN, country only, on SC62",
			InputFilename:  "dvCountry.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - DV with empty subjectDN, on SC62",
			InputFilename:  "dvEmptySubject.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "na - EV certificate",
			InputFilename:  "evAllGood.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cab_dv_subject_invalid_values", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
			if tc.ExpectedResult == lint.Error && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %q, was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
