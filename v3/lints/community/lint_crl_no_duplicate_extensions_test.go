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

package community

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCRLNoDuplicateExtensions(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		path string
		want lint.LintStatus
	}{
		{
			name: "CRL with no duplicate extensions",
			path: "crl_no_duplicate_ext.pem",
			want: lint.Pass,
		},
		{
			name: "CRL with duplicate extensions",
			path: "crl_with_duplicate_ext.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := test.TestRevocationListLint(t, "e_crl_no_duplicate_extensions", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
