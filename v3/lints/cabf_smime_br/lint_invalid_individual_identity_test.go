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

func TestInvalidIndividualIdentity(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}

	// Test files naming scheme:
	// leg1/leg0: legacy policy asserted (1) or not asserted (0)
	// polxx: policy xx is asserted (where xx is one of iv, ov, mv, sv)
	// eff1/eff0: issued on or after effective date (1), or before (0)
	// pseu1/pseu0: pseudonym present (1) or absent (0)
	// pers1/pers0: personal name present (1) or absent (0)
	// personal name present means givenName and/or surname are present

	data := []Data{
		{
			input: "leg1_poliv_eff1_pseu0_pers0.pem",
			want:  lint.NA,
		},
		{
			input: "leg0_polmv_eff1_pseu0_pers0.pem",
			want:  lint.NA,
		},
		{
			input: "leg0_polov_eff1_pseu0_pers0.pem",
			want:  lint.NA,
		},
		{
			input: "leg0_poliv_eff0_pseu0_pers0.pem",
			want:  lint.NE,
		},
		{
			input: "leg0_poliv_eff1_pseu0_pers0.pem",
			want:  lint.Error,
		},
		{
			input: "leg0_poliv_eff1_pseu1_pers0.pem",
			want:  lint.Pass,
		},
		{
			input: "leg0_poliv_eff1_pseu0_pers1.pem",
			want:  lint.Pass,
		},
		{
			input: "leg0_polsv_eff0_pseu0_pers0.pem",
			want:  lint.NE,
		},
		{
			input: "leg0_polsv_eff1_pseu0_pers0.pem",
			want:  lint.Error,
		},
		{
			input: "leg0_polsv_eff1_pseu1_pers0.pem",
			want:  lint.Pass,
		},
		{
			input: "leg0_polsv_eff1_pseu0_pers1.pem",
			want:  lint.Pass,
		},
	}

	for _, testData := range data {
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_invalid_individual_identity", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
