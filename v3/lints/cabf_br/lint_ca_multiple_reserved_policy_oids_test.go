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

//  TEST CASES - file naming scheme:
//
//      ca:     CA certificate
//      sub:    Subordinate CA certificate
//      cr:     Cross-certificate (based on config)
//      any:    Contains the anyPolicy OID
//      mul:    Contains multiple CABF Reserved Policy Identifiers
//      var:    Variant (i.e. combination of Reserved Policy Identifiers)
//      eff:    Effective (i.e. certficate was issued on or after the Effective Date)
//
//      0 = No
//      1 = Yes
//      x = Does not matter

func TestCAMultipleReservedPolicyOIDs(t *testing.T) {
	type Data struct {
		input  string
		config string
		want   lint.LintStatus
	}
	data := []Data{
		{
			input: "ca0_subx_crx_anyx_mulx_varx_effx.pem",
			want:  lint.NA,
		},
		{
			input: "ca1_sub0_crx_anyx_mulx_varx_effx.pem",
			want:  lint.NA,
		},
		{
			input: "ca1_sub1_cr1_anyx_mulx_varx_effx.pem",
			config: `
                [e_ca_multiple_reserved_policy_oids]
                CrossCert = true
                `,
			want: lint.NA,
		},
		{
			input: "ca1_sub1_cr0_any1_mulx_varx_effx.pem",
			want:  lint.NA,
		},
		{
			input: "ca1_sub1_cr0_any0_mul0_varx_eff1.pem",
			want:  lint.Pass,
		},
		{
			input: "ca1_sub1_cr0_any0_mul1_varx_eff0.pem",
			want:  lint.NE,
		},
		{
			input: "ca1_sub1_cr0_any0_mul1_vara_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "ca1_sub1_cr0_any0_mul1_varb_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "ca1_sub1_cr0_any0_mul1_varc_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "ca1_sub1_cr0_any0_mul1_vard_eff1.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLintWithConfig("e_ca_multiple_reserved_policy_oids", testData.input, testData.config)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
