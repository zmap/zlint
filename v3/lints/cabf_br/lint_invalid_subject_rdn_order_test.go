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

//nolint:all
/*
   === Proper RDN order test cases
   subject_rdn_order_ok_01.pem             C, ST, L, O, CN
   subject_rdn_order_ok_02.pem             C, ST, L, postalCode, street, O, CN
   subject_rdn_order_ok_03.pem             <empty subject>
   subject_rdn_order_ok_04.pem             DC, DC, C, ST, L, O, CN
   subject_rdn_order_ok_05.pem             C, ST, L, street, O, CN, serialNumber, businessCategory, jurisdictionCountry
   subject_rdn_order_ok_06.pem             C, ST, L, SN, givenName, CN
   subject_rdn_order_ok_07.pem             CN

   === Wrong RDN order test cases
   subject_rdn_order_ko_01.pem             C, ST, L, CN, O
   subject_rdn_order_ko_02.pem             CN, O, L, ST, C
   subject_rdn_order_ko_03.pem             C, ST, L, O, CN, street
   subject_rdn_order_ko_04.pem             C, ST, L, O, CN, DC, DC
   subject_rdn_order_ko_05.pem             C, ST, L, givenName, SN, CN
   subject_rdn_order_ko_06.pem             C, ST, L, street, postalCode, O
   subject_rdn_order_ko_07.pem             CN, C
*/

func TestInvalidSubjectRDNOrder(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "subject_rdn_order_ok_01.pem",
			want:  lint.Pass,
		},
		{
			input: "subject_rdn_order_ok_02.pem",
			want:  lint.Pass,
		},
		{
			input: "subject_rdn_order_ok_03.pem",
			want:  lint.Pass,
		},
		{
			input: "subject_rdn_order_ok_04.pem",
			want:  lint.Pass,
		},
		{
			input: "subject_rdn_order_ok_05.pem",
			want:  lint.Pass,
		},
		{
			input: "subject_rdn_order_ok_06.pem",
			want:  lint.Pass,
		},
		{
			input: "subject_rdn_order_ok_07.pem",
			want:  lint.Pass,
		},
		{
			input: "subject_rdn_order_ko_01.pem",
			want:  lint.Error,
		},
		{
			input: "subject_rdn_order_ko_02.pem",
			want:  lint.Error,
		},
		{
			input: "subject_rdn_order_ko_03.pem",
			want:  lint.Error,
		},
		{
			input: "subject_rdn_order_ko_04.pem",
			want:  lint.Error,
		},
		{
			input: "subject_rdn_order_ko_05.pem",
			want:  lint.Error,
		},
		{
			input: "subject_rdn_order_ko_06.pem",
			want:  lint.Error,
		},
		{
			input: "subject_rdn_order_ko_07.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_invalid_subject_rdn_order", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
