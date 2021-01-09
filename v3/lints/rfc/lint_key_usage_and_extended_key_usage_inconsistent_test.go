package rfc

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

	"github.com/zmap/zcrypto/x509"
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

//Tests for verifying the truth tables
func TestEKUServerAuth(t *testing.T) {
	got := KeyUsage(x509.KeyUsageDigitalSignature).
		Xor(KeyUsage(x509.KeyUsageKeyEncipherment).
			Xor(KeyUsage(x509.KeyUsageKeyAgreement)))
	for w := range serverAuth {
		if !got[KeyUsage(w)] {
			t.Errorf("expected %d to be present in the computed map", w)
		}
	}
	for g := range got {
		if !serverAuth[x509.KeyUsage(g)] {
			t.Errorf("expected %d to be present in the pre-computed map", g)
		}
	}
}

func TestEKUClientAuth(t *testing.T) {
	got := KeyUsage(x509.KeyUsageDigitalSignature).
		Or(KeyUsage(x509.KeyUsageKeyAgreement))
	for w := range clientAuth {
		if !got[KeyUsage(w)] {
			t.Errorf("expected %d to be present in the computed map", w)
		}
	}
	for g := range got {
		if !clientAuth[x509.KeyUsage(g)] {
			t.Errorf("expected %d to be present in the pre-computed map", g)
		}
	}
}

func TestEKUCodeSigning(t *testing.T) {
	got := Accepted{KeyUsage(x509.KeyUsageDigitalSignature): true}
	for w := range codeSigning {
		if !got[KeyUsage(w)] {
			t.Errorf("expected %d to be present in the computed map", w)
		}
	}
	for g := range got {
		if !codeSigning[x509.KeyUsage(g)] {
			t.Errorf("expected %d to be present in the pre-computed map", g)
		}
	}
}

func TestEKUEmailProtection(t *testing.T) {
	got := KeyUsage(x509.KeyUsageDigitalSignature).
		Or(KeyUsage(x509.KeyUsageContentCommitment).
			Or(KeyUsage(x509.KeyUsageKeyEncipherment).Xor(KeyUsage(x509.KeyUsageKeyAgreement))))
	for w := range emailProtection {
		if !got[KeyUsage(w)] {
			t.Errorf("expected %d to be present in the computed map", w)
		}
	}
	for g := range got {
		if !emailProtection[x509.KeyUsage(g)] {
			t.Errorf("expected %d to be present in the pre-computed map", g)
		}
	}
}

func TestEKUTimeStamping(t *testing.T) {
	got := KeyUsage(x509.KeyUsageDigitalSignature).
		Or(KeyUsage(x509.KeyUsageContentCommitment))
	for w := range timeStamping {
		if !got[KeyUsage(w)] {
			t.Errorf("expected %d to be present in the computed map", w)
		}
	}
	for g := range got {
		if !timeStamping[x509.KeyUsage(g)] {
			t.Errorf("expected %d to be present in the pre-computed map", g)
		}
	}
}

func TestEKUOCSPSigning(t *testing.T) {
	got := KeyUsage(x509.KeyUsageDigitalSignature).
		Or(KeyUsage(x509.KeyUsageContentCommitment))
	for w := range ocspSigning {
		if !got[KeyUsage(w)] {
			t.Errorf("expected %d to be present in the computed map", w)
		}
	}
	for g := range got {
		if !ocspSigning[x509.KeyUsage(g)] {
			t.Errorf("expected %d to be present in the pre-computed map", g)
		}
	}
}

// Util for truth table tests
type Truther interface {
	Or(a Truther) Accepted
	Xor(a Truther) Accepted
}

type KeyUsage x509.KeyUsage
type Accepted map[KeyUsage]bool

func (usage KeyUsage) Or(truther Truther) Accepted {
	if truther == nil {
		return Accepted{usage: true}
	}
	switch t := truther.(type) {
	case KeyUsage:
		return map[KeyUsage]bool{
			usage:     true,
			t:         true,
			usage | t: true,
		}
	case Accepted:
		t.Or(usage)
		return t
	default:
		panic("")
	}
}

func (usage KeyUsage) Xor(truther Truther) Accepted {
	if truther == nil {
		return Accepted{usage: true}
	}
	switch t := truther.(type) {
	case KeyUsage:
		return Accepted{
			usage: true,
			t:     true,
		}
	case Accepted:
		t.Xor(usage)
		return t
	default:
		panic("")
	}
}

func (accepted Accepted) Or(truther Truther) Accepted {
	if truther == nil {
		return accepted
	}
	switch other := truther.(type) {
	case KeyUsage:
		accepted[other] = true
		for current := range accepted {
			accepted[current|other] = true
		}
	case Accepted:
		for key := range other {
			accepted[key] = true
			for inner := range accepted {
				accepted[key|inner] = true
			}
		}
	}
	return accepted
}

func (accepted Accepted) Xor(truther Truther) Accepted {
	if truther == nil {
		return accepted
	}
	switch other := truther.(type) {
	case KeyUsage:
		accepted[other] = true
	case Accepted:
		for key := range other {
			accepted.Xor(key)
		}
	}
	return accepted
}
