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

func TestInvalidOrgIDRegistrationScheme(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "iors_ee_vat_e1_x1.pem",
			want:  lint.Pass,
		},
		{
			input: "iors_ee_ntr_e1_x1.pem",
			want:  lint.Pass,
		},
		{
			input: "iors_ee_psd_e1_x1.pem",
			want:  lint.Pass,
		},
		{
			input: "iors_ee_abc_e1_x1.pem",
			want:  lint.Error,
		},
		{
			input: "iors_ee_abc_e1_x0.pem",
			want:  lint.NA,
		},
		{
			input: "iors_ee_nul_e1_x1.pem",
			want:  lint.NA,
		},
		{
			input: "iors_ee_abc_e0_x1.pem",
			want:  lint.NE,
		},
		{
			input: "iors_ca_abc_e1_x1.pem",
			want:  lint.NA,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_ev_invalid_orgid_reg_scheme", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}

}
