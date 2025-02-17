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

package cabf_ev

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

/*
 * Test cases
 *
 *      File                         Description
 *      ------------------------     -------------
 *      extra_subj_attrs_na1.pem     CA certificate
 *      extra_subj_attrs_na2.pem     OV Subscriber certificate
 *      extra_subj_attrs_ok1.pem     EV Subscriber certificate with valid Subject
 *      extra_subj_attrs_ne1.pem     EV Subscriber certificate with invalid Subject, issued before effective date
 *      extra_subj_attrs_ko1.pem     EV Subscriber certificate with invalid Subject, issued after effective date
 *
 */

func TestExtraSubjectAttribs(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "extra_subj_attrs_na1.pem",
			want:  lint.NA,
		},
		{
			input: "extra_subj_attrs_na2.pem",
			want:  lint.NA,
		},
		{
			input: "extra_subj_attrs_ok1.pem",
			want:  lint.Pass,
		},
		{
			input: "extra_subj_attrs_with_ou_ok2.pem",
			want:  lint.Pass,
		},
		{
			input: "extra_subj_attrs_ne1.pem",
			want:  lint.NE,
		},
		{
			input: "extra_subj_attrs_ko1.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_ev_extra_subject_attribs", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
