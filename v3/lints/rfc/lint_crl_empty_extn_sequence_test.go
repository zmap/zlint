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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCRLEmptyExtnSequence(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected lint.LintStatus
	}{
		{"all entries with non-empty or absent extensions", "crl_empty_extn_seq_ok.pem", lint.Pass},
		{"first entry without extensions", "crl_empty_extn_seq_first_entry_no_ext_ok.pem", lint.Pass},
		{"entry with empty extensions SEQUENCE", "crl_empty_extn_seq_ko.pem", lint.Error},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := test.TestRevocationListLint(t, "e_crl_empty_extn_sequence", tc.input)
			if out.Status != tc.expected {
				t.Errorf("%s: expected %s, got %s", tc.input, tc.expected, out.Status)
			}
		})
	}
}
