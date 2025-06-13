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

// TEST CASES
// File naming scheme:
//      sub  = Subscriber certificate
//      val  = Validity period (days)
//      iat  = Issued After Threshold date (15/3/2026)
//      cdp  = Certificate contains CRL Distribution Point(s)
//      ocps = Certificate contains OCSP pointer(s)
//      eff  = Certificate issued after this lint's effective date

func TestMissingCRLDistribPoint(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "sub0_valxx_iatx_cdpx_ocspx_effx.pem",
			want:  lint.NA,
		},
		{
			input: "sub1_val30_iatx_cdp0_ocsp0_eff0.pem",
			want:  lint.NE,
		},
		{
			input: "sub1_val30_iatx_cdp1_ocsp0_eff1.pem",
			want:  lint.Pass,
		},
		{
			input: "sub1_val30_iatx_cdp0_ocsp1_eff1.pem",
			want:  lint.Pass,
		},
		{
			input: "sub1_val30_iatx_cdp0_ocsp0_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "sub1_val10_iat1_cdp0_ocsp0_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "sub1_val10_iat0_cdp0_ocsp0_eff1.pem",
			want:  lint.NA,
		},
		{
			input: "sub1_val07_iatx_cdp0_ocsp0_eff1.pem",
			want:  lint.NA,
		},
		{
			input: "sub1_val10_iat0_cdp0_ocsp1_eff1.pem",
			want:  lint.NA,
		},
		{
			input: "sub1_val07_iatx_cdp0_ocsp1_eff1.pem",
			want:  lint.NA,
		},
	}

	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_missing_crl_distrib_point", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
