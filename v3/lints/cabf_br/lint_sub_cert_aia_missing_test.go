package cabf_br

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

func TestSubCertAiaMissing(t *testing.T) {

	var tests = []struct {
		name      string
		inputPath string

		expected lint.LintStatus
	}{
		{
			name:      "pass - cert with an AIA extension",
			inputPath: "subCertWBothURL.pem",

			expected: lint.Pass,
		},
		{
			name:      "error - cert dated before BRs 2.3.1 missing AIA",
			inputPath: "subCertWNoURL.pem",

			expected: lint.Error,
		},
		{
			name:      "not effective - test case for w_sub_cert_aia_missing",
			inputPath: "subCertWNoURLPostCABFBR231.pem",

			expected: lint.NE,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			out := test.TestLint("e_sub_cert_aia_missing", testCase.inputPath)

			if out.Status != testCase.expected {
				t.Errorf("%s: expected %s, got %s", testCase.inputPath, testCase.expected, out.Status)
			}
		})
	}
}
