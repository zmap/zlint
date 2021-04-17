package rfc

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubjectSurnameMaxLength(t *testing.T) {
	data := []struct {
		input string
		want  lint.LintStatus
	}{
		{"surnameUnder64.pem", lint.Pass},
		{"surnameOver32768.pem", lint.Error},
	}
	for _, d := range data {
		input := d.input
		want := d.want
		t.Run(input, func(t *testing.T) {
			got := test.TestLint("e_subject_surname_max_length", input).Status
			if want != got {
				t.Errorf("%s: expected %s, got %s", input, want, got)
			}
		})
	}
}
