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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCRLMissingNullRSAParams(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "CRL signed with sha224WithRSAEncryption, with NULL params",
			path: "crl_rsa_sha224_nul1_effx.pem",
			want: lint.Pass,
		},
		{
			desc: "CRL signed with sha224WithRSAEncryption, without NULL params",
			path: "crl_rsa_sha224_nul0_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "CRL signed with sha256WithRSAEncryption, with NULL params",
			path: "crl_rsa_sha256_nul1_effx.pem",
			want: lint.Pass,
		},
		{
			desc: "CRL signed with sha256WithRSAEncryption, without NULL params",
			path: "crl_rsa_sha256_nul0_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "CRL signed with sha256WithRSAEncryption, without NULL params, before Effective Date",
			path: "crl_rsa_sha256_nul0_eff0.pem",
			want: lint.NE,
		},
		{
			desc: "CRL signed with sha384WithRSAEncryption, with NULL params",
			path: "crl_rsa_sha384_nul1_effx.pem",
			want: lint.Pass,
		},
		{
			desc: "CRL signed with sha384WithRSAEncryption, without NULL params",
			path: "crl_rsa_sha384_nul0_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "CRL signed with sha512WithRSAEncryption, with NULL params",
			path: "crl_rsa_sha512_nul1_effx.pem",
			want: lint.Pass,
		},
		{
			desc: "CRL signed with sha512WithRSAEncryption, without NULL params",
			path: "crl_rsa_sha512_nul0_eff1.pem",
			want: lint.Error,
		},
		{
			desc: "CRL signed with ecdsa-with-SHA384",
			path: "crl_ecc_sha384_nulx_effx.pem",
			want: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestRevocationListLint(t, "e_crl_sigalgo_missing_null_params", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
