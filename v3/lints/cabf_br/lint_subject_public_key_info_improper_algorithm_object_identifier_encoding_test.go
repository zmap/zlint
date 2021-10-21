package cabf_br

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

func TestWrongSubjectPublicKeyAlgorithmIdentifierObjectEncoding(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Wrong subject public key algorithm identifier object algorithm",
			InputFilename:  "dsaCert.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Correct subject public key algorithm identifier for RSA",
			InputFilename:  "publicKeyIsRSAWithCorrectEncoding.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Correct subject public key algorithm identifier for P256",
			InputFilename:  "publicKeyIsECCP256WithCorrectEncoding.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Correct subject public key algorithm identifier for P384",
			InputFilename:  "publicKeyIsECCP384WithCorrectEncoding.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Correct subject public key algorithm identifier for P521",
			InputFilename:  "publicKeyIsECCP521WithCorrectEncoding.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "Public Key is RSA but the explicit NULL is missing from the parameters",
			InputFilename:   "publicKeyIsRSAExplicitNullMissing.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "The encoded AlgorithmObjectIdentifier for RSA inside the the SubjectPublicKeyInfo field is \"300b06092a864886f70d010101\" but the expected one is \"300d06092a864886f70d0101010500\".",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_algorithm_identifier_improper_encoding", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
