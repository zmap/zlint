package cabf_br

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

func TestCnNotExactlyFromSAN(t *testing.T) {
	var testCases = []struct {
		name      string
		inputFile string

		expectedOutput lint.LintStatus
	}{
		{
			name:           "Pass - commonName in SAN.DNSNames",
			inputFile:      "SANWithCNSeptember2021.pem",
			expectedOutput: lint.Pass,
		},
		{
			name:           "Error - common name not in SAN.DNSNames",
			inputFile:      "SANWithoutCNSeptember2021.pem",
			expectedOutput: lint.Error,
		},
		{
			name:           "Error - common name in SAN.DNSNames but case mismatch",
			inputFile:      "SANCaseNotMatchingCNSeptember2021.pem",
			expectedOutput: lint.Error,
		},
		{
			name:           "NE - certificate issued before 21 August 2021",
			inputFile:      "SANWithMissingCN.pem",
			expectedOutput: lint.NE,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := test.TestLint("e_subject_common_name_not_exactly_from_san", tc.inputFile)
			if out.Status != tc.expectedOutput {
				t.Errorf("%s: expected %s, got %s", tc.inputFile, tc.expectedOutput, out.Status)
			}
		})
	}
}
