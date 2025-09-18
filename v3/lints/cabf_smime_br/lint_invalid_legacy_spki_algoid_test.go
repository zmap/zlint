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

/*
   TEST CASES - file naming convention
   ===================================
   smc1/0      Certificate is/is not for S/MIME
   alg_xxx     Self-explanatory (followed by algo name)
   eff1/0      Certificate issued after/before Effective Date
   pqf1/0      Certificate issued after/before Effective Date for PQC algorithms
*/

func TestInvalidLegacySPKIAlgoId(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	data := []Data{
		{
			input: "smime/sm0_alg_xxx_effx_pqfx.pem",
			want:  lint.NA,
		},
		{
			input: "smime/sm1_alg_xxx_eff0_pqfx.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_alg_rsa_eff1_pqfx.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_p224_eff1_pqfx.pem",
			want:  lint.Error,
		},
		{
			input: "smime/sm1_alg_p256_eff1_pqfx.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_p384_eff1_pqfx.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_p521_eff1_pqfx.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_ed25519_eff1_pqfx.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_mld44_eff1_pqf0.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_mld65_eff1_pqf0.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_mld87_eff1_pqf0.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_mlk512_eff1_pqf0.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_mlk768_eff1_pqf0.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_mlk1024_eff1_pqf0.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_alg_mld44_eff1_pqf1.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_alg_mld65_eff1_pqf1.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_alg_mld87_eff1_pqf1.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_alg_mlk512_eff1_pqf1.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_alg_mlk768_eff1_pqf1.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_alg_mlk1024_eff1_pqf1.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_alg_gost_eff1_pqfx.pem",
			want:  lint.Error,
		},
	}

	for _, testData := range data {
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_invalid_legacy_spki_algoid", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
