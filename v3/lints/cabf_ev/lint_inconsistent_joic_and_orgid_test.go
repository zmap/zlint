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

func TestInconsistentJoiCountryAndOrgID(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Non-EV certificate",
			path: "incons_orgid_joic_ev0_ox_jx_bx_nx_cx_efx.pem",
			want: lint.NA,
		},
		{
			desc: "EV certificate without organizationIdentifier",
			path: "incons_orgid_joic_ev1_o0_jx_bx_nx_cx_efx.pem",
			want: lint.NA,
		},
		{
			desc: "EV certificate with organizationIdentifier but without joiCountry",
			path: "incons_orgid_joic_ev1_o1_j0_bx_nx_cx_efx.pem",
			want: lint.NA,
		},
		{
			desc: "EV cert with organizationIdentifier (non-NTR scheme) and joiCountry",
			path: "incons_orgid_joic_ev1_o1_j1_b0_n0_cx_efx.pem",
			want: lint.Pass,
		},
		{
			desc: "EV cert with consistent organizationIdentifier and joiCountry",
			path: "incons_orgid_joic_ev1_o1_j1_b0_n1_c1_ef1.pem",
			want: lint.Pass,
		},
		{
			desc: "EV cert with inconsistent orgId and joiCountry, issued before effective date",
			path: "incons_orgid_joic_ev1_o1_j1_b0_n1_c0_ef0.pem",
			want: lint.NE,
		},
		{
			desc: "EV cert with inconsistent orgId and joiCountry, issued after effective date",
			path: "incons_orgid_joic_ev1_o1_j1_b0_n1_c0_ef1.pem",
			want: lint.Error,
		},
		{
			desc: "EV cert with invalid value in orgId, issued before effective date",
			path: "incons_orgid_joic_ev1_o1_jx_b1_nx_cx_ef0.pem",
			want: lint.NE,
		},
		{
			desc: "EV cert with invalid value in orgId (case A), issued after effective date",
			path: "incons_orgid_joic_ev1_o1_jx_b1_nx_cx_ef1_a.pem",
			want: lint.Error,
		},
		{
			desc: "EV cert with invalid value in orgId (case B), issued after effective date",
			path: "incons_orgid_joic_ev1_o1_jx_b1_nx_cx_ef1_b.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("e_inconsistent_joic_and_orgid", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
