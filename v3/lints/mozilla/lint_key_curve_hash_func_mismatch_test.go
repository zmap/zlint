/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

package mozilla

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestKeyCurveHashFuncMismatch(t *testing.T) {
	testCases := []struct {
		Name string
		Path string
		Want lint.LintStatus
	}{
		{
			Name: "Certificate not signed with ECDSA",
			Path: "ecc0_pxxx_shaxxx_eff1.pem",
			Want: lint.NA,
		},
		{
			Name: "Certificate signed with P256 and SHA256, issued after effective date",
			Path: "ecc1_p256_sha256_eff1.pem",
			Want: lint.Pass,
		},
		{
			Name: "Certificate signed with P256 and SHA384, issued before effective date",
			Path: "ecc1_p256_sha384_eff0.pem",
			Want: lint.NE,
		},
		{
			Name: "Certificate signed with P256 and SHA384, issued after effective date",
			Path: "ecc1_p256_sha384_eff1.pem",
			Want: lint.Error,
		},
		{
			Name: "Certificate signed with P384 and SHA256, issued before effective date",
			Path: "ecc1_p384_sha256_eff0.pem",
			Want: lint.NE,
		},
		{
			Name: "Certificate signed with P384 and SHA256, issued after effective date",
			Path: "ecc1_p384_sha256_eff1.pem",
			Want: lint.Error,
		},
		{
			Name: "Certificate signed with P384 and SHA384, issued after effective date",
			Path: "ecc1_p384_sha384_eff1.pem",
			Want: lint.Pass,
		},
		{
			Name: "Certificate signed with P521 and SHA256, issued after effective date",
			Path: "ecc1_p521_sha256_eff1.pem",
			Want: lint.Error,
		},
		{
			Name: "Certificate signed with P521 and SHA384, issued after effective date",
			Path: "ecc1_p521_sha384_eff1.pem",
			Want: lint.Error,
		},
		{
			Name: "Certificate signed with P521 and SHA512, issued after effective date",
			Path: "ecc1_p521_sha512_eff1.pem",
			Want: lint.Pass,
		},
		{
			Name: "Certificate signed with P521 and SHA512, issued before effective date",
			Path: "ecc1_p521_sha512_eff0.pem",
			Want: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_key_curve_hash_func_mismatch", tc.Path)
			if result.Status != tc.Want {
				t.Errorf("expected %v, got %v", tc.Want, result.Status)
			}
		})
	}
}
