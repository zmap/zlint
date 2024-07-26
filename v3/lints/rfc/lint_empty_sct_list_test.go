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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

/*
   === Pass test cases ===
   empty_sct_list_ok_01.pem    SCTList extension NOT present
   empty_sct_list_ok_02.pem    SCTList extension present, with length > 0

   === NA test cases ===
   empty_sct_list_na_01.pem    Precertificate (Poison extension present)
   empty_sct_list_na_02.pem    CA certificate

   === Fail test cases ===
   empty_sct_list_ko_01.pem    SCTList extension present, with zero length
*/

func TestEmptySCTList(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "empty_sct_list_ok_01.pem",
			want:  lint.Pass,
		},
		{
			input: "empty_sct_list_ok_02.pem",
			want:  lint.Pass,
		},
		{
			input: "empty_sct_list_na_01.pem",
			want:  lint.NA,
		},
		{
			input: "empty_sct_list_na_02.pem",
			want:  lint.NA,
		},
		{
			input: "empty_sct_list_ko_01.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_empty_sct_list", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
