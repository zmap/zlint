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

package community

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSuspiciousCharsInAddress(t *testing.T) {
	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "certificate without localityName or stateOrProvinceName",
			path: "suspicious_locnull_stnull.pem",
			want: lint.NA,
		},
		{
			desc: "certificate without localityName, with clean stateOrProvinceName",
			path: "suspicious_locnull_stgood.pem",
			want: lint.Pass,
		},
		{
			desc: "certificate without localityName, with suspicious stateOrProvinceName",
			path: "suspicious_locnull_stbad.pem",
			want: lint.Notice,
		},
		{
			desc: "certificate with clean localityName, without stateOrProvinceName",
			path: "suspicious_locgood_stnull.pem",
			want: lint.Pass,
		},
		{
			desc: "certificate with suspicious localityName, without stateOrProvinceName",
			path: "suspicious_locbad_stnull.pem",
			want: lint.Notice,
		},
		{
			desc: "certificate with clean localityName and stateOrProvinceName",
			path: "suspicious_locgood_stgood.pem",
			want: lint.Pass,
		},
		{
			desc: "certificate with clean localityName and suspicious stateOrProvinceName",
			path: "suspicious_locgood_stbad.pem",
			want: lint.Notice,
		},
		{
			desc: "certificate with suspicious localityName and stateOrProvinceName",
			path: "suspicious_locbad_stbad.pem",
			want: lint.Notice,
		},
		{
			desc: "certificate with two localityName values, only the second suspicious",
			path: "suspicious_locmulti_secondbad.pem",
			want: lint.Notice,
		},
		{
			desc: "certificate with two stateOrProvinceName values, only the second suspicious",
			path: "suspicious_stmulti_secondbad.pem",
			want: lint.Notice,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("n_suspicious_chars_in_address", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
