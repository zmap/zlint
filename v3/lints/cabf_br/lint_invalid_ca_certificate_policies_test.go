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
 * Test file naming convention: (cacp_cXrXpXmXaXbXeX.pem)
 *      X = 0/1 for no/yes
 *      c = CA certificate
 *      r = Root CA (self-signed)
 *      p = CertificatePolicies is present
 *      m = Multiple OIDs in CertificatePolicies
 *      a = The AnyPolicy OID is present
 *      b = At least one CABF policy OID is present
 *      e = Certificate issued after effective date
 */

func TestInvalidCACertificatePolicies(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "cacp_c0r0p0m0a0b0e0.pem",
			want:  lint.NA,
		},
		{
			input: "cacp_c1r1p0m0a0b0e0.pem",
			want:  lint.NA,
		},
		{
			input: "cacp_c1r0p0m0a0b0e0.pem",
			want:  lint.Error,
		},
		{
			input: "cacp_c1r0p1m0a0b0e1.pem",
			want:  lint.Error,
		},
		{
			input: "cacp_c1r0p1m0a0b1e1.pem",
			want:  lint.Pass,
		},
		{
			input: "cacp_c1r0p1m0a1b0e1.pem",
			want:  lint.Pass,
		},
		{
			input: "cacp_c1r0p1m1a1b0e0.pem",
			want:  lint.Error,
		},
		{
			input: "cacp_c1r0p1m1a0b0e0.pem",
			want:  lint.NE,
		},
		{
			input: "cacp_c1r0p1m1a0b0e1.pem",
			want:  lint.Error,
		},
		{
			input: "cacp_c1r0p1m1a0b1e1.pem",
			want:  lint.Pass,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_invalid_ca_certificate_policies", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
