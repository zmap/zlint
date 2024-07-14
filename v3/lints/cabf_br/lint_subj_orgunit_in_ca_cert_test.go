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
 * Test cases:
 *
 *      Input file              Config      Want        Description
 *      ==========              ======      ====        ===========
 *      orgunit_in_ca_ok1.pem   (none)      NA          Subscriber cert with OU, issued before effective date
 *      orgunit_in_ca_ok4.pem   (none)      NA          Non-TLS CA cert with OU, issued before effective date
 *      orgunit_in_ca_ok2.pem   (none)      Pass        TLS CA cert without OU
 *      orgunit_in_ca_ok3.pem   (none)      NE          TLS CA cert with OU, issued before effective date
 *      orgunit_in_ca_ko1.pem   (none)      Error       TLS CA cert with OU, issued after effective date
 *      orgunit_in_ca_ko1.pem   CrossCert   Pass        TLS CA cert with OU, issued after effective date
 */

package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubjectOrgUnitInCACert(t *testing.T) {
	type Data struct {
		input  string
		config string
		want   lint.LintStatus
	}
	data := []Data{
		{
			input: "orgunit_in_ca_ok1.pem",
			want:  lint.NA,
		},
		{
			input: "orgunit_in_ca_ok2.pem",
			want:  lint.Pass,
		},
		{
			input: "orgunit_in_ca_ok3.pem",
			want:  lint.NE,
		},
		{
			input: "orgunit_in_ca_ok4.pem",
			want:  lint.NA,
		},
		{
			input: "orgunit_in_ca_ko1.pem",
			want:  lint.Error,
		},
		{
			input: "orgunit_in_ca_ko1.pem",
			config: `
                [e_subj_orgunit_in_ca_cert]
                CrossCert = true
                `,
			want: lint.Pass,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLintWithConfig("e_subj_orgunit_in_ca_cert", testData.input, testData.config)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
