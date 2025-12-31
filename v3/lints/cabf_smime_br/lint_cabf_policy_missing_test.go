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
   sm1/0      Certificate for S/MIME: yes/no
   sub1/0/x   Subscriber certificate: yes/no/don't care
   cp1/0/x    Contains a CABF S/MIME BR reserved policy OID: yes/no/don't care
   ef1/0/x    Certificate issued after Effective Date: yes/no/don't care
*/

func TestCABFPolicyMissing(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	data := []Data{
		{
			input: "smime/sm0_subx_cpx_efx.pem",
			want:  lint.NA,
		},
		{
			input: "smime/sm1_sub0_cpx_efx.pem",
			want:  lint.NA,
		},
		{
			input: "smime/sm1_sub1_cp0_ef0.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_sub1_cp0_ef1.pem",
			want:  lint.Error,
		},
		{
			input: "smime/sm1_sub1_cp1_ef1.pem",
			want:  lint.Pass,
		},
	}

	for _, testData := range data {
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_cabf_policy_missing", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
