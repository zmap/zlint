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

func TestSubCertReservedPolicyMissing(t *testing.T) {

	var tests = []struct {
		name      string
		inputPath string
		details   string
		expected  lint.LintStatus
	}{
		{
			name:      "not applicable - sub CA certificate",
			inputPath: "subCAAIAValid.pem",
			details:   "",
			expected:  lint.NA,
		},
		{
			name:      "not effective - certificate issued before CAB/F 2.0.0",
			inputPath: "sha1ExpirePrior2017.pem",
			details:   "",
			expected:  lint.NE,
		},
		{
			name:      "pass - certificate contains a reserved certificate policy identifier",
			inputPath: "domainValWithProvincePost200.pem",
			details:   "",
			expected:  lint.Pass,
		},
		{
			name:      "error - TLS server certificate without policies",
			inputPath: "basicConstraintsCriticalSC62.pem",
			details:   "Missing a CA/B Forum TLS reserved certificate policy identifier",
			expected:  lint.Error,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			out := test.TestLint("e_sub_cert_reserved_policy_missing", testCase.inputPath)

			if out.Status != testCase.expected {
				t.Errorf("%s: expected %s, got %s", testCase.inputPath, testCase.expected, out.Status)
			}
			if out.Details != "" && out.Details != testCase.details {
				t.Errorf("%s: expected details %q, got %q", testCase.inputPath, testCase.details, out.Details)
			}
		})
	}
}
