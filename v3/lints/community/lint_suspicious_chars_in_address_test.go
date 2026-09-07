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

func TestSuspiciousCharsInAddress(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Certificate without L or ST",
			path: "suspicious_locnull_stnull.pem",
			want: lint.NA,
		},
		{
			desc: "Certificate without L, with good ST",
			path: "suspicious_locnull_stgood.pem",
			want: lint.Pass,
		},
		{
			desc: "Certificate without L, with bad ST",
			path: "suspicious_locnull_stbad.pem",
			want: lint.Warn,
		},
		{
			desc: "Certificate with good L, without ST",
			path: "suspicious_locgood_stnull.pem",
			want: lint.Pass,
		},
		{
			desc: "Certificate with bad L, without ST",
			path: "suspicious_locbad_stnull.pem",
			want: lint.Warn,
		},
		{
			desc: "Certificate with good L and ST",
			path: "suspicious_locgood_stgood.pem",
			want: lint.Pass,
		},
		{
			desc: "Certificate with good L and bad ST",
			path: "suspicious_locgood_stbad.pem",
			want: lint.Warn,
		},
		{
			desc: "Certificate with bad L and bad ST",
			path: "suspicious_locbad_stbad.pem",
			want: lint.Warn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("w_suspicious_chars_in_address", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}

}
