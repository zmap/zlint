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
 *      Naming convention for unit test files:
 *
 *      c - CA certificate (any kind of)
 *      r - Root CA certificate (i.e. self-signed)
 *      e - EKU extension present in certificate
 *      a - EKU contains anyExtendedKeyUsage
 *      s - EKU contains serverAuth
 *      m - EKU contains emailProtection
 *      b - Certificate issued before effective date
 */

func TestCaInvalidEKU(t *testing.T) {
	type Data struct {
		input string
		want  lint.LintStatus
	}
	data := []Data{
		{
			input: "c0r0e0a0s0m0b0.pem",
			want:  lint.NA,
		},
		{
			input: "c1r1e0a0s0m0b0.pem",
			want:  lint.NA,
		},
		{
			input: "c1r0e0a0s0m0b0.pem",
			want:  lint.NA,
		},
		{
			input: "c1r0e1a1s0m0b0.pem",
			want:  lint.Pass,
		},
		{
			input: "c1r0e1a1s0m1b0.pem",
			want:  lint.Error,
		},
		{
			input: "c1r0e1a0s1m0b0.pem",
			want:  lint.Pass,
		},
		{
			input: "c1r0e1a0s0m1b0.pem",
			want:  lint.NA,
		},
		{
			input: "c1r0e1a0s1m1b0.pem",
			want:  lint.Error,
		},
		{
			input: "c1r0e1a0s1m1b1.pem",
			want:  lint.NE,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_ca_invalid_eku", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}
