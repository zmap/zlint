package mozilla

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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
	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/test"
	"testing"
)

func TestPssAidEncoding(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Standard PSS with SHA256",
			InputFilename:  "pssWithSHA256.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard PSS with SHA256 but the hash parameters are empty instead of NULL",
			InputFilename:  "pssWithSHA256EmptyHashParams.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Standard PSS with SHA384",
			InputFilename:  "pssWithSHA384.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard PSS with SHA384 but the hash parameters are empty instead of NULL",
			InputFilename:  "pssWithSHA384EmptyHashParams.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Standard PSS with SHA512",
			InputFilename:  "pssWithSHA512.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard PSS with SHA512 but the hash parameters are empty instead of NULL",
			InputFilename:  "pssWithSHA512EmptyHashParams.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mp_pss_parameters_encoding_correct", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
