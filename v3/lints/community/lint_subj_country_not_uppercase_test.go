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
 * Test cases:
 *
 *      country_not_upcase_ok1.pem      Normal
 *      country_not_upcase_ko1.pem      Country code is in mixed case
 *      country_not_upcase_ko2.pem      Country code is all lowercase
 *      country_not_upcase_ko3.pem      Two country codes, one OK and one bad
 */

func TestSubjCountryNotUppercase(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "country_not_upcase_ok1.pem",
			want:  lint.Pass,
		},
		{
			input: "country_not_upcase_ko1.pem",
			want:  lint.Error,
		},
		{
			input: "country_not_upcase_ko2.pem",
			want:  lint.Error,
		},
		{
			input: "country_not_upcase_ko3.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_subj_country_not_uppercase", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
