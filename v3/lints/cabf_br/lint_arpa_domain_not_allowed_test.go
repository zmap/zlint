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

func TestARPADomainNotAllowed(t *testing.T) {
	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Non-Subscriber certificate",
			path: "arpa_sub0_rev4x_rev6x_effx.pem",
			want: lint.NA,
		},
		{
			desc: "Certificate without any domains that end in an IP Reverse Zone Suffix (RZS)",
			path: "arpa_sub1_rev40_rev60_effx.pem",
			want: lint.Pass,
		},
		{
			desc: "Certificate with a domain that end in an IPv4 RZS, issued before effective date",
			path: "arpa_sub1_rev41_rev60_eff0.pem",
			want: lint.NE,
		},
		{
			desc: "Certificate with a domain that ends in an IPv4 RZS, issued after effective date",
			path: "arpa_sub1_rev41_rev60_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with a domain that end in an IPv6 RZS, issued before effective date",
			path: "arpa_sub1_rev40_rev61_eff0.pem",
			want: lint.NE,
		},
		{
			desc: "Certificate with a domain that ends in an IPv6 RZS, issued after effective date",
			path: "arpa_sub1_rev40_rev61_eff1.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("e_arpa_domain_not_allowed", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
