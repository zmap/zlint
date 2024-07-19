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
 * Test cases:
 *
 *      Input file              Want        Description
 *      ==========              ====        ===========
 *      cdp_not_http_na1.pem    NA          Subscriber cert with no CDPs at all
 *      cdp_not_http_ne1.pem    NE          Subscriber cert with an LDAP CDP, but issued before effective date
 *      cdp_not_http_ok1.pem    Pass        Subscriber cert with single HTTP CDP
 *      cdp_not_http_ok2.pem    Pass        Subscriber cert with double HTTP CDP
 *      cdp_not_http_ko1.pem    Error       Subscriber cert with single LDAP CDP, issued after effective date
 *      cdp_not_http_ko2.pem    Error       Subscriber cert with an HTTP CDP and an LDAP CDP, issued after effective date
 */

func TestCrlDistribPointsNotHTTP(t *testing.T) {
	type Data struct {
		input  string
		config string
		want   lint.LintStatus
	}
	data := []Data{
		{
			input: "cdp_not_http_na1.pem",
			want:  lint.NA,
		},
		{
			input: "cdp_not_http_ne1.pem",
			want:  lint.NE,
		},
		{
			input: "cdp_not_http_ok1.pem",
			want:  lint.Pass,
		},
		{
			input: "cdp_not_http_ok2.pem",
			want:  lint.Pass,
		},
		{
			input: "cdp_not_http_ko1.pem",
			want:  lint.Error,
		},
		{
			input: "cdp_not_http_ko2.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLintWithConfig("e_crl_distrib_points_not_http", testData.input, testData.config)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
