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

package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

/*
   === Pass test cases ===
   invalid_cps_uri_ok_01.pem       Certificate with a well-formed CPS URI
   invalid_cps_uri_ok_02.pem       Certificate without a CPS URI

   === NE test cases ===
   invalid_cps_uri_ok_03.pem       Certificate with an invalid CPS URI, but issued before effective date

   === Fail test cases ===
   invalid_cps_uri_ko_01.pem       Certificate with an invalid CPS URI (disallowed scheme)
   invalid_cps_uri_ko_02.pem       Certificate with an invalid CPS URI (syntax error)
   invalid_cps_uri_ko_03.pem       Certificate with two CPS URIs, one good and one bad
*/

func TestInvalidCPSUri(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "invalid_cps_uri_ok_01.pem",
			want:  lint.Pass,
		},
		{
			input: "invalid_cps_uri_ok_02.pem",
			want:  lint.Pass,
		},
		{
			input: "invalid_cps_uri_ok_03.pem",
			want:  lint.NE,
		},
		{
			input: "invalid_cps_uri_ko_01.pem",
			want:  lint.Error,
		},
		{
			input: "invalid_cps_uri_ko_02.pem",
			want:  lint.Error,
		},
		{
			input: "invalid_cps_uri_ko_03.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_invalid_cps_uri", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
