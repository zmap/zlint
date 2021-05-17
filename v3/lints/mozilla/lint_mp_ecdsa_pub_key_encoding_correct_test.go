package mozilla

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

func TestECDSAPubKeyAidEncoding(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Standard ECC certificate with a P-256 key signed by a P-256 key",
			InputFilename:  "eccP256.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard ECC certificate with a P-384 key signed by a P-384 key",
			InputFilename:  "eccP384.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Standard ECC certificate with a P-521 key signed by a P-521 key",
			InputFilename:  "eccP521.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Certificate with an RSA key",
			InputFilename:  "evAllGood.pem",
			ExpectedResult: lint.NA,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mp_ecdsa_pub_key_encoding_correct", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
