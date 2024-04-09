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
   === Pass test cases ===
   invalid_business_cat_ok_01.pem       EV cert with valid businessCategory == "Private Organization"
   invalid_business_cat_ok_04.pem       EV cert with valid businessCategory == "Government Entity"
   invalid_business_cat_ok_05.pem       EV cert with valid businessCategory == "Business Entity"
   invalid_business_cat_ok_06.pem       EV cert with valid businessCategory == "Non‐Commercial Entity"

   === NA test cases ===
   invalid_business_cat_ok_02.pem       EV cert without businessCategory
   invalid_business_cat_ok_03.pem       OV cert with invalid businessCategory

   === Fail test cases ===
   invalid_business_cat_ko_01.pem       EV cert with slightly invalid businessCategory
   invalid_business_cat_ko_02.pem       EV cert with grossly invalid businessCategory
*/

package cabf_ev

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestInvalidBusinessCategory(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "invalid_business_cat_ok_01.pem",
			want:  lint.Pass,
		},
		{
			input: "invalid_business_cat_ok_04.pem",
			want:  lint.Pass,
		},
		{
			input: "invalid_business_cat_ok_05.pem",
			want:  lint.Pass,
		},
		{
			input: "invalid_business_cat_ok_06.pem",
			want:  lint.Pass,
		},
		{
			input: "invalid_business_cat_ok_02.pem",
			want:  lint.NA,
		},
		{
			input: "invalid_business_cat_ok_03.pem",
			want:  lint.NA,
		},
		{
			input: "invalid_business_cat_ko_01.pem",
			want:  lint.Error,
		},
		{
			input: "invalid_business_cat_ko_02.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_ev_invalid_business_category", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
