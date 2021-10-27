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
			Name:            "Wrong subject public key algorithm identifier object algorithm",
			InputFilename:   "dsaCert.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "The encoded AlgorithmObjectIdentifier \"3082012b06072a8648ce3804013082011e02818100931d0880233aece9e2b816fb0e2daecc2b044e6131f401d784266b16fdf12992bb098f19f108ce4395f323859e7dfd19c88c2e75c976ca76c4ec61ec39efe745124683b726436926b79a36acac5ed9a02cd55bed1653912e10b5422823cf6d6b80057c88fe2da1fba521642142303a9f76c5cfcdf6d79dc4da1a6678f7d8cde3021500d13f595a85e4b55fe6f4c4b58090a979c03f212d02818029cc9723232468277f26e5148324661b0d2f54099cda8bbdd455f3f6faf33e72b99ed49b04358d82213d6ef4c3a70ed4f604d04814d60ff69c8307edaf3d49c596bebb0198797469d15422efcdb68a028c8aba632539576e9d5d077bd61b4abb6496cb58ea18e998c5123e551dc78a7c1bdd064dec12ef138be63a98159fa898\" inside the SubjectPublicKeyInfo field is not allowed",
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
			ExpectedDetails: "The encoded AlgorithmObjectIdentifier \"300b06092a864886f70d010101\" inside the SubjectPublicKeyInfo field is not allowed",
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
