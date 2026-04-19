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

func TestExcessivelyBackdated(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Certificate without SCTs",
			path: "excbakdat_sct0_oldx_effx.pem",
			want: lint.NA,
		},
		{
			desc: "Certificate with SCTs and acceptable notBefore",
			path: "excbakdat_sct1_old0_effx.pem",
			want: lint.Pass,
		},
		{
			desc: "Certificate with SCTs and a bad notBefore, issued before Effective Date",
			path: "excbakdat_sct1_old1_eff0.pem",
			want: lint.NE,
		},
		{
			desc: "Certificate with SCTs and a bad notBefore, issued after Effective Date",
			path: "excbakdat_sct1_old1_eff1.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		out := test.TestLint("e_excessively backdated", tc.path)
		if out.Status != tc.want {
			t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
		}
	}
}
