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

func TestCRLNumberRange(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		path string
		want lint.LintStatus
	}{
		{
			name: "CRL with a valid number",
			path: "crl_number_valid.pem",
			want: lint.Pass,
		},
		{
			name: "CRL with a negative number",
			path: "crl_number_negative.pem",
			want: lint.Error,
		},
		{
			name: "CRL with a number >= 2^159",
			path: "crl_number_too_large.pem",
			want: lint.Error,
		},
		{
			name: "CRL without a CRL number extension",
			path: "crl_no_crl_number.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := test.TestRevocationListLint(t, "e_crl_number_out_of_range", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
