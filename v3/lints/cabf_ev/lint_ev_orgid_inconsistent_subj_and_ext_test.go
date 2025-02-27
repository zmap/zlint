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

/*
 * Contributed by Adriano Santoni <adriano.santoni@staff.aruba.it>
 * of ACTALIS S.p.A. (www.actalis.com).
 */

/*
   === Pass test cases ===
   orgid_subj_and_ext_ok_01.pem         EV cert with orgId=="VATIT-1234567890" and cabfOrgId consistent

   === NA test cases ===
   orgid_subj_and_ext_ok_02.pem         OV cert with orgId=="VATIT-1234567890" and cabfOrgId NOT consistent
   orgid_subj_and_ext_ok_04.pem         EV cert without orgId
   orgid_subj_and_ext_ok_05.pem         EV cert with orgId but NO cabfOrgId (which is wrong, but not this lint's business)

   === NE test cases ===
   orgid_subj_and_ext_ok_03.pem         EV cert with orgId and cabfOrgId NOT consistent, but issued before 31/1/2020

   === Fail test cases ===
   orgid_subj_and_ext_ko_01.pem         EV cert with orgId=="NTRUS+CA-1234567890" and cabfOrgId NOT consistent
   orgid_subj_and_ext_ko_02.pem         EV cert with orgId=="PSDAT-FMA-1234567890" and cabfOrgId NOT consistent
   orgid_subj_and_ext_ko_03.pem         EV cert with invalid orgId ("VATBEE-12345")
*/

package cabf_ev

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestOrgIdInconsistentSubjAndExt(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	data := []Data{
		{
			input: "orgid_subj_and_ext_ok_01.pem",
			want:  lint.Pass,
		},
		{
			input: "orgid_subj_and_ext_ok_02.pem",
			want:  lint.NA,
		},
		{
			input: "orgid_subj_and_ext_ok_03.pem",
			want:  lint.NE,
		},
		{
			input: "orgid_subj_and_ext_ok_04.pem",
			want:  lint.NA,
		},
		{
			input: "orgid_subj_and_ext_ok_05.pem",
			want:  lint.NA,
		},
		{
			input: "orgid_subj_and_ext_ok_06.pem",
			want:  lint.Pass,
		},
		{
			input: "orgid_subj_and_ext_ko_01.pem",
			want:  lint.Error,
		},
		{
			input: "orgid_subj_and_ext_ko_02.pem",
			want:  lint.Error,
		},
		{
			input: "orgid_subj_and_ext_ko_03.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_ev_orgid_inconsistent_subj_and_ext", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
