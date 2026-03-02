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

func TestEmailAddrNotInSAN(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Certificate without email in Subject",
			path: "semnis_subj0_sanx_matx_effx.pem",
			want: lint.NA,
		},
		{
			desc: "Certificate with email in Subject, issued before effective date",
			path: "semnis_subj1_sanx_matx_eff0.pem",
			want: lint.NE,
		},
		{
			desc: "Certificate with email in Subject, no SAN, issued after effective date",
			path: "semnis_subj1_san0_matx_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with email in Subject, no match in SAN, issued after effective date",
			path: "semnis_subj1_san1_mat0_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with email in Subject, match in SAN, issued after effective date",
			path: "semnis_subj1_san1_mat1_eff1.pem",
			want: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("e_subj_email_not_in_san", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
