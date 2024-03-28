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

package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestPolicyQualifiersOtherThanCpsNotPermitted(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Valid",
			InputFilename:  "policyQualifiersOtherThanCpsNotPermittedValid.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Error",
			InputFilename:  "policyQualifiersOtherThanCpsNotPermittedError.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Not Applicable",
			InputFilename:  "policyQualifiersOtherThanCpsNotPermittedNotApplicable.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_policy_qualifiers_other_than_cps_not_permitted", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
