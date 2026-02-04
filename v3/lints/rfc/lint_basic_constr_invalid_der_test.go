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

func TestBasicConstraintsInvalidDER(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	data := []Data{
		{
			// Certificate without BasicConstraints
			input: "basic_constr_invalid_der_na.pem",
			want:  lint.NA,
		},
		{
			// Cert with correctly encoded BasicConstraints
			input: "basic_constr_invalid_der_ok.pem",
			want:  lint.Pass,
		},
		{
			// Cert with wrongly encoded BasicConstraints
			input: "basic_constr_invalid_der_ko.pem",
			want:  lint.Error,
		},
	}

	for _, testData := range data {
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_basic_constr_invalid_der", testData.input)
			if out.Status != testData.want {
				t.Errorf("%s: expected %s, got %s", testData.input, testData.want, out.Status)
			}
		})
	}
}
