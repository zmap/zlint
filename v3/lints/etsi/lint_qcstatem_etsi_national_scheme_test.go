package etsi

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

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestQcStatemNationalScheme(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "NA - certificate has the natural person semantics identifier and no national scheme value",
			InputFilename:  "qcNaturalNoNationalScheme.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Pass - certificate has the natural person semantics identifier and a correct national scheme value",
			InputFilename:  "qcNaturalCorrectNationalScheme.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Error - certificate has the natural person semantics identifier and a wrong national scheme value",
			InputFilename:  "qcNaturalNotCorrectScheme.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "NA - certificate has the legal person semantics identifier and no national scheme value",
			InputFilename:  "qcLegalNoNationalScheme.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Pass - certificate has the legal person semantics identifier and a correct national scheme value",
			InputFilename:  "qcLegalCorrectNationalScheme.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Error - certificate has the legal person semantics identifier and a wrong national scheme value",
			InputFilename:  "qcLegalNotCorrectScheme.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_qcstatem_correct_national_scheme", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
