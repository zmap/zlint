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

func TestQcNaturalPersonKUMandatory(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Pass - certificate is issued to a natural person and has keu usage extension",
			InputFilename:  "etsi/qcNaturalWithKU.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - certificate is issued to a legal person and has keu usage extension",
			InputFilename:  "etsi/qcLegalWithKU.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:            "Error - certificate is issued to a natural person and does not have the keu usage extension",
			InputFilename:   "etsi/qcNaturalWithoutKU.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "ETSI natural person certificates does not have the key usage extension",
		},
		{
			Name:           "NA - certificate is issued to a legal person because subjectDN parts for natural persons are not present",
			InputFilename:  "etsi/qcNaturalOrLegalWithKU.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_etsi_natural_person_key_usage_mandatory", tc.InputFilename)

			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %v was %v", tc.ExpectedDetails, result.Details)
			}

			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
