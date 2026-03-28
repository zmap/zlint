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

package chrome

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestClientAuthNotAllowed(t *testing.T) {
	
	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Non-Subscriber certificate",
			path: "cana_sub0_srvx_clix_effx.pem",
			want: lint.NA,
		},
		{
			desc: "Subscriber certificate but not for serverAuth use",
			path: "cana_sub1_srv0_clix_effx.pem",
			want: lint.NA,
		},
		{
			desc: "Server certificate without clientAuth in EKU",
			path: "cana_sub1_srv1_cli0_effx.pem",
			want: lint.Pass,
		},
		{
			desc: "Server cert with clientAuth in EKU, issued before effective date",
			path: "cana_sub1_srv1_cli1_eff0.pem",
			want: lint.NE,
		},
		{
			desc: "Server cert with clientAuth in EKU, issued after effective date",
			path: "cana_sub1_srv1_cli1_eff1.pem",
			want: lint.Error,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("e_client_auth_not_allowed", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
