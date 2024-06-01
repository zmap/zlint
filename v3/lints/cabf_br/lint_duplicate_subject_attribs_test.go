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

package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

/*
 *      TEST CASES
 *
 *      Filename                        Exp res     Certificate features
 *      --------                        -------     --------------------
 *      dupl_subj_attrs_ok_01.pem       Pass        Normal TLS certificate
 *      dupl_subj_attrs_na_01.pem       NA          Normal S/MIME certificate
 *      dupl_subj_attrs_ne_01.pem       NE          With duplicate O, but issued before effective date
 *      dupl_subj_attrs_ko_01.pem       Error       With duplicate O (in separate RDNs), issued after effective date
 *      dupl_subj_attrs_ko_02.pem       Error       With duplicate C (in the same RDN), issued after effective date
 *      dupl_subj_attrs_ko_03.pem       Error       With duplicate L (in separate RDNs), issued after effective date
 */

func TestDuplicateSubjectAttribs(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "dupl_subj_attrs_ok_01.pem",
			want:  lint.Pass,
		},
		{
			input: "dupl_subj_attrs_na_01.pem",
			want:  lint.NA,
		},
		{
			input: "dupl_subj_attrs_ne_01.pem",
			want:  lint.NE,
		},
		{
			input: "dupl_subj_attrs_ko_01.pem",
			want:  lint.Error,
		},
		{
			input: "dupl_subj_attrs_ko_02.pem",
			want:  lint.Error,
		},
		{
			input: "dupl_subj_attrs_ko_03.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_duplicate_subject_attribs", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
