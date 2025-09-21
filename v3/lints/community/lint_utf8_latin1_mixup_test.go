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

package community

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

/*
   TEST CASES

   File                     Result      Description
   ===================      ======      ===========
   utf8_lat1_mixup_ok1      Pass        Clean TLS certificate (no mixup anywhere)
   utf8_lat1_mixup_ko1      Error       TLS certificate with mixup in stateOrProvince
   utf8_lat1_mixup_ko2      Error       S/MIME certificate with mixup in givenName
   utf8_lat1_mixup_ko3      Error       Code Signing certificate with mixup in organizationName
*/

func TestUTF8Latin1Mixup(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	data := []Data{
		{
			input: "utf8_lat1_mixup_ok1.pem",
			want:  lint.Pass,
		},
		{
			input: "utf8_lat1_mixup_ko1.pem",
			want:  lint.Error,
		},
		{
			input: "utf8_lat1_mixup_ko2.pem",
			want:  lint.Error,
		},
		{
			input: "utf8_lat1_mixup_ko3.pem",
			want:  lint.Error,
		},
	}

	for _, testData := range data {
		//		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_utf8_latin1_mixup", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}

}
