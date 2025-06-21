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

func TestCANonHTTPCAIssuersURL(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	// === TEST FILES NAMING SCHEME ===
	// subca: Subordinate CA certificate
	// ocsp: Contains the ocsp accessMethod in the AIA extension
	// cais: Contains the caIssuers accessMethod in the AIA extension
	// http: accessLocation contains an 'http://' URL
	// eff: issued after effective date

	data := []Data{
		{
			input: "subca0_ocspx_httpx_caisx_httpx_effx.pem",
			want:  lint.NA,
		},
		{
			input: "subca1_ocsp0_httpx_cais0_httpx_effx.pem",
			want:  lint.NA,
		},
		{
			input: "subca1_ocsp1_http0_cais0_httpx_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "subca1_ocsp1_http1_cais0_httpx_eff1.pem",
			want:  lint.Pass,
		},
		{
			input: "subca1_ocsp0_httpx_cais1_http0_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "subca1_ocsp0_httpx_cais1_http1_eff1.pem",
			want:  lint.Pass,
		},
		{
			input: "subca1_ocsp1_http0_cais0_httpx_eff0.pem",
			want:  lint.NE,
		},
		{
			input: "subca1_ocsp0_httpx_cais1_http0_eff0.pem",
			want:  lint.NE,
		},
		{
			input: "subca1_ocsp1_http1_cais1_http1_eff1.pem",
			want:  lint.Pass,
		},
	}

	for _, testData := range data {
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_ca_aia_non_http_url", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
