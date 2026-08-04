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

package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestOrgValidatedInvalidCN(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "CA certificate",
			path: "smime/sub0_smx_ovx_cnex_cnox_effx.pem",
			want: lint.NA,
		},
		{
			desc: "Non-S/MIME certificate",
			path: "smime/sub1_sm0_ovx_cnex_cnox_effx.pem",
			want: lint.NA,
		},
		{
			desc: "Non-OV S/MIME certificate",
			path: "smime/sub1_sm1_ov0_cnex_cnox_effx.pem",
			want: lint.NA,
		},
		{
			desc: "OV S/MIME certificate with email in CN",
			path: "smime/sub1_sm1_ov1_cne1_cno0_eff1.pem",
			want: lint.Pass,
		},
		{
			desc: "OV S/MIME certificate with CN matching O",
			path: "smime/sub1_sm1_ov1_cne0_cno1_eff1.pem",
			want: lint.Pass,
		},
		{
			desc: "OV S/MIME certificate with bad CN, issued before effective date",
			path: "smime/sub1_sm1_ov1_cne0_cno0_eff0.pem",
			want: lint.NE,
		},
		{
			desc: "OV S/MIME certificate with bad CN, issued after effective date",
			path: "smime/sub1_sm1_ov1_cne0_cno0_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "OV S/MIME certificate with no CN",
			path: "smime/sub1_sm1_ov1_cn_absent_eff1.pem",
			want: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("e_org_validated_invalid_cn", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
