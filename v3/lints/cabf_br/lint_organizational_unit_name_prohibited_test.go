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

func TestOrganizationalUnitNameProhibited(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Certificate issued after rule that doesn't have an OU",
			InputFilename:  "ouAbsentAfterSep22.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Certificate issued before rule comes into effect",
			InputFilename:  "ouPresentBeforeSep22.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "CA Certificate issued after rule comes into effect",
			InputFilename:  "ouPresentCATrueAfterSep22.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Certificate issued after rule applies that contains an OU",
			InputFilename:  "ouPresentAfterSep22.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_organizational_unit_name_prohibited", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
