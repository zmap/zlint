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

func TestLegacyGenerationDeprecated(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	// Test files naming scheme:
	// leg1: legacy policy asserted
	// leg0: legacy policy NOT asserted
	// mv|iv|sv|ov: certificate type
	// eff1: issued on or after effective date
	// eff0: issued before effective date

	data := []Data{
		{
			input: "smime_leg1_mv_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "smime_leg1_iv_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "smime_leg1_sv_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "smime_leg1_ov_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "smime_leg1_xx_eff0.pem",
			want:  lint.NE,
		},
		{
			input: "smime_leg0_xx_eff1.pem",
			want:  lint.NA,
		},
	}

	for _, testData := range data {
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_legacy_generation_deprecated", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
