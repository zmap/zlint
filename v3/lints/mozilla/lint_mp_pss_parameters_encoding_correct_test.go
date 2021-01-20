package mozilla

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

func TestPssAidEncoding(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Standard RSASSA-PSS with SHA256",
			InputFilename:  "rsassapssWithSHA256.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard RSASSA-PSS with SHA256 but the hash parameters are empty instead of NULL",
			InputFilename:  "rsassapssWithSHA256EmptyHashParams.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Standard RSASSA-PSS with SHA384",
			InputFilename:  "rsassapssWithSHA384.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard RSASSA-PSS with SHA384 but the hash parameters are empty instead of NULL",
			InputFilename:  "rsassapssWithSHA384EmptyHashParams.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Standard RSASSA-PSS with SHA512",
			InputFilename:  "rsassapssWithSHA512.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard RSASSA-PSS with SHA512 but the hash parameters are empty instead of NULL",
			InputFilename:  "rsassapssWithSHA512EmptyHashParams.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Standard RSASSA-PSS with SHA256 but the salt length is 17 instead of 32",
			InputFilename:  "rsassapssWithSHA256ButIrregularSaltLength.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mp_rsassa-pss_parameters_encoding_in_signature_algorithm_correct", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
