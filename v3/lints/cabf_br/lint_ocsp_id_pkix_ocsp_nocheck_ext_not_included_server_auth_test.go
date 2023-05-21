package cabf_br

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestOCSPIDPKIXOCSPNocheckExtNotIncludedServerAuth(t *testing.T) {
	testCases := []struct {
		Name           string
		Filename       string
		ExpectedResult lint.LintStatus
	}{

		// Legend for the nameing:
		// o1 --> EKU OCSPSigning set; o0 not set
		// s1 --> EKU serverAuth set, s0 not set
		// ep1 --> EKU emailProtection set, ep0 not set
		// a1 --> EKU anyExtendedKeyUsage set, a0 not set
		// nc1 --> noCheck set, nc0 not set

		{
			Name:           "o0s0ep0a0nc0",
			Filename:       "o0s0ep0a0nc0.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s0ep0a0nc1",
			Filename:       "o0s0ep0a0nc1.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s0ep0a1nc0",
			Filename:       "o0s0ep0a1nc0.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s0ep0a1nc1",
			Filename:       "o0s0ep0a1nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o0s1ep0a0nc0",
			Filename:       "o0s1ep0a0nc0.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s1ep0a0nc1",
			Filename:       "o0s1ep0a0nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o0s1ep0a1nc0",
			Filename:       "o0s1ep0a1nc0.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o0s1ep0a1nc1",
			Filename:       "o0s1ep0a1nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o1s0ep0a0nc0",
			Filename:       "o1s0ep0a0nc0.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o1s0ep0a0nc1",
			Filename:       "o1s0ep0a0nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o1s0ep0a1nc0",
			Filename:       "o1s0ep0a1nc0.pem",
			ExpectedResult: lint.Error,
		}, {
			Name:           "o1s0ep0a1nc1",
			Filename:       "o1s0ep0a1nc1.pem",
			ExpectedResult: lint.Pass,
		}, {
			Name:           "o1s1ep0a0nc0",
			Filename:       "o1s1ep0a0nc0.pem",
			ExpectedResult: lint.Error,
		}, {
			Name:           "o1s1ep0a0nc1",
			Filename:       "o1s1ep0a0nc1.pem",
			ExpectedResult: lint.Pass,
		}, {
			Name:           "o1s1ep0a1nc0",
			Filename:       "o1s1ep0a1nc0.pem",
			ExpectedResult: lint.Error,
		}, {
			Name:           "o1s1ep0a1nc1",
			Filename:       "o1s1ep0a1nc1.pem",
			ExpectedResult: lint.Pass,
		}, {
			Name:           "o0s0ep1a0nc0",
			Filename:       "o0s0ep1a0nc0.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s0ep1a0nc1",
			Filename:       "o0s0ep1a0nc1.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s0ep1a1nc0",
			Filename:       "o0s0ep1a1nc0.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s0ep1a1nc1",
			Filename:       "o0s0ep1a1nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o0s1ep1a0nc0",
			Filename:       "o0s1ep1a0nc0.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "o0s1ep1a0nc1",
			Filename:       "o0s1ep1a0nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o0s1ep1a1nc0",
			Filename:       "o0s1ep1a1nc0.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o0s1ep1a1nc1",
			Filename:       "o0s1ep1a1nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o1s0ep1a0nc0",
			Filename:       "o1s0ep1a0nc0.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o1s0ep1a0nc1",
			Filename:       "o1s0ep1a0nc1.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "o1s0ep1a1nc0",
			Filename:       "o1s0ep1a1nc0.pem",
			ExpectedResult: lint.Error,
		}, {
			Name:           "o1s0ep1a1nc1",
			Filename:       "o1s0ep1a1nc1.pem",
			ExpectedResult: lint.Pass,
		}, {
			Name:           "o1s1ep1a0nc0",
			Filename:       "o1s1ep1a0nc0.pem",
			ExpectedResult: lint.Error,
		}, {
			Name:           "o1s1ep1a0nc1",
			Filename:       "o1s1ep1a0nc1.pem",
			ExpectedResult: lint.Pass,
		}, {
			Name:           "o1s1ep1a1nc0",
			Filename:       "o1s1ep1a1nc0.pem",
			ExpectedResult: lint.Error,
		}, {
			Name:           "o1s1ep1a1nc1",
			Filename:       "o1s1ep1a1nc1.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ocsp_id_pkix_ocsp_nocheck_ext_not_included_server_auth", tc.Filename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
