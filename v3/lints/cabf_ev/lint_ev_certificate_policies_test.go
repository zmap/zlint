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

package cabf_ev

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestEvCertificatePolicies(t *testing.T) {
	tests := []struct {
		inputPath string
		expected  lint.LintStatus
	}{
		{
			inputPath: "evAllGood.pem",
			expected:  lint.Pass,
		},
		{
			inputPath: "evNoCPSURI.pem",
			expected:  lint.Error,
		},
		{
			inputPath: "dnsNameValidTLD.pem",
			expected:  lint.NA,
		},
	}
	for _, tt := range tests {
		t.Run(tt.inputPath, func(t *testing.T) {
			out := test.TestLint("e_ev_certificate_policies", tt.inputPath)
			if out.Status != tt.expected {
				t.Errorf("%s: expected %s, got %s", tt.inputPath, tt.expected, out.Status)
			}
		})
	}
}
