package rfc

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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

func TestKeyUsageAndExtendedKeyUsageInconsistent(t *testing.T) {
	testCases := []struct {
		Name           string
		Filename       string
		ExpectedResult lint.LintStatus
	}{

		// Legend for the nameing:

		// a1 --> EKU anyExtendedKeyUsage set, a0 not set
		// s1 --> EKU serverAuth set, s0 not set
		// c1 --> EKU clientAuth set, c0 not set
		// cs1 --> EKU codeSigning set; cs0 not set
		// ep1 --> EKU emailProtection set, ep0 not set
		// ts1 --> EKU timeStamping set; ts0 not set
		// o1 --> EKU OCSPSigning set; o0 not set

		// nc1 --> noCheck set, nc0 not set

		// ds1 --> KU digitalSignature set
		// cc1 --> KU contentCommitment set
		// ke1 --> KU keyEncipherment set
		// de1 --> KU dataEncipherment set
		// ka1 --> KU keyAgreement set
		// c1 --> KU CertSign set
		// crl1 --> KU CRLSign set
		// eo1 --> KU encipherOnly set
		// do1 --> KU decipherOnly set

		// Tests for 1 EKU bit set
		{
			Name:           "TestConsistentKUBitWithServerAuthEKUBit",
			Filename:       "a1s1c0cs0ep0ts0o0nc1_ds1.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Test2KUXorBitsServerAuthEKUBit",
			Filename:       "a1s1c0cs0ep0ts0o0nc1_ds1ke1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Test1ConsistentAnd1InconsistentKUBitWithServerAuthEKUBit",
			Filename:       "a1s1c0cs0ep0ts0o0nc1_cc1ke1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Test2ConsistentKUBitsWithClientAuthEKUBit",
			Filename:       "a1s0c1cs0ep0ts0o0nc1_ds1ka1.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "TestInconsistentKUBitWithClientAuthEKUBit",
			Filename:       "a1s0c1cs0ep0ts0o0nc1_de1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestConsistentKUBitWithCodeSigningEKUBit",
			Filename:       "a1s0c0cs1ep0ts0o0nc1_ds1.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "TestInconsistentKUBitsWithCodeSigningEKUBit",
			Filename:       "a1s0c0cs1ep0ts0o0nc1_ds1cc1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestConsistentKUBitsWithEmailProtectionEKUBit",
			Filename:       "a1s0c0cs0ep1ts0o0nc1_ds1cc1ke1.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "TestInconsistentKUBitsWithEmailProtectionEKUBit",
			Filename:       "a1s0c0cs0ep1ts0o0nc1_ds1cc1ke1ka1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestConsistentKUBitWithOCSPSigningEKUBit",
			Filename:       "a1s0c0cs0ep0ts0o1nc1_ds1.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "TestInconsistentKUBitsWithOCSPSigningEKUBit",
			Filename:       "a1s0c0cs0ep0ts0o1nc1_ds1cc1ke1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TestConsistentKUBitsWithTimeStampingEKUBit",
			Filename:       "a1s0c0cs0ep0ts1o0nc1_ds1cc1.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "TestInconsistentKUBitsWithTimeStampingEKUBit",
			Filename:       "a1s0c0cs0ep0ts1o0nc1_ke1.pem",
			ExpectedResult: lint.Error,
		},
		// Test for all EKU bits set
		{
			Name:           "TestAllEKUBitsSetWithConsistentKUBits",
			Filename:       "a1s1c1cs1ep1ts1o1nc1_ds1.pem",
			ExpectedResult: lint.Pass,
		},
		// Tests for multiple EKU bits set
		{
			Name:           "Test2EKUBitsSetWithConsistentKUBits",
			Filename:       "a1s0c0cs0ep0ts1o1nc1_ds1cc1.pem",
			ExpectedResult: lint.Pass,
		},

		{
			Name:           "Test2EKUBitsSetWith1Consistent1InconsistentKUBits",
			Filename:       "a1s0c1cs1ep0ts0o0nc1_ka1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Test2EKUBitsSetWith2InconsistentKUBits",
			Filename:       "a1s1c1cs0ep0ts0o0nc1_cc1c1.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_key_usage_and_extended_key_usage_inconsistent", tc.Filename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
