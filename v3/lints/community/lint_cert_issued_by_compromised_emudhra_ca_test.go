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

package community

import (
	"crypto/elliptic"
	"crypto/sha1"
	"encoding/hex"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCertIssuedByCompromisedEmudhraCA(t *testing.T) {
	testCases := []struct {
		name      string
		inputPath string
		expected  lint.LintStatus
	}{
		// https://crt.sh/?id=30010070587
		{"EV G3B subscriber, signed with the compromised key", "emudhraEvG3BSubscriber.pem", lint.Error},
		// https://crt.sh/?id=29977123837
		{"OV G3B subscriber, signed with the compromised key", "emudhraOvG3BSubscriber.pem", lint.Error},
		// https://crt.sh/?id=29977124909
		{"EV G3C subscriber, signed with the compromised key", "emudhraEvG3CSubscriber.pem", lint.Error},
		// Not in CT: minted for this test with the recovered private keys, using a
		// throwaway leaf key and a "zlint.invalid" subject.
		{"OV G3C, minted with the recovered key", "emudhraOvG3CMinted.pem", lint.Error},
		{"DV G3C, minted with the recovered key", "emudhraDvG3CMinted.pem", lint.Error},
		{"CS G3B, minted with the recovered key", "emudhraCsG3BMinted.pem", lint.Error},
		// No AKI at all.
		{"DV G3B, minted, no AKI", "emudhraDvG3BMintedNoAKI.pem", lint.Error},
		// AKI of 20 bytes of 0x11, which is not any eMudhra SKI.
		{"CS G3C, minted, wrong AKI", "emudhraCsG3CMintedWrongAKI.pem", lint.Error},
		// Pre-existing test certificates in testdata/, not eMudhra certificates.
		{"ECDSA-signed by an unrelated issuer", "aiaWithValidNames.pem", lint.Pass},
		{"RSA-signed is filtered out", "sub1_val10_iat0_cdp0_ocsp0_eff1.pem", lint.NA},
		// https://crt.sh/?sha256=E7568FE6E12331F4FD7DE61DA4C7559D0F8EB192A218D4741A3BA75184333838
		{"real intermediate, signed by the emSign root", "emudhraEvG3BIntermediate.pem", lint.Pass},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := test.TestLint("e_cert_issued_by_compromised_emudhra_ca", tc.inputPath)
			if out.Status != tc.expected {
				t.Errorf("%s: expected %s, got %s", tc.inputPath, tc.expected, out.Status)
			}
		})
	}
}

// The embedded public keys are checked against the Subject Key Identifiers
// listed in the Bugzilla report. eMudhra computed each SKI as the SHA-1 of the
// key's public point, so a mistyped or swapped key will not match.
func TestCertIssuedByCompromisedEmudhraCAEmbeddedKeys(t *testing.T) {
	expectedSKIs := map[string]string{
		"EM EV TLS CA - G3B": "3edc20d9894698877b873fd88147b25c15d84706",
		"EM OV TLS CA - G3B": "9bd86551afe7e72c37ad08ee6919d9bf3784ae0a",
		"EM DV TLS CA - G3B": "e21e2fe71c19390e554b62c120245735248a2514",
		"EM EV TLS CA - G3C": "a6255e80a94db7f3967c1bc791df92bcee4beb5f",
		"EM OV TLS CA - G3C": "34eaea0cebc90adaf9d0c5e2ed9b78c9879e1b6c",
		"EM DV TLS CA - G3C": "ae7e105c3c2996c1a847270559f1e03dd4e50af3",
		"EM OV CS CA - G3B":  "79fae8eccfdb0d758da458d69fb6eddc7c49dee6",
		"EM OV CS CA - G3C":  "ee6f73298d6fbfa7b31bc39534f61556447866b7",
	}
	if len(compromisedEmudhraKeys) != len(expectedSKIs) {
		t.Fatalf("expected %d embedded keys, got %d", len(expectedSKIs), len(compromisedEmudhraKeys))
	}
	for name, pub := range compromisedEmudhraKeys {
		want, ok := expectedSKIs[name]
		if !ok {
			t.Errorf("unexpected embedded key %q", name)
			continue
		}
		if pub.X == nil || pub.Y == nil || !pub.Curve.IsOnCurve(pub.X, pub.Y) {
			t.Errorf("%s: not a valid P-384 point", name)
			continue
		}
		sum := sha1.Sum(elliptic.Marshal(pub.Curve, pub.X, pub.Y))
		if got := hex.EncodeToString(sum[:]); got != want {
			t.Errorf("%s: key hashes to SKI %s, expected %s", name, got, want)
		}
	}
}

func TestCertIssuedByCompromisedEmudhraCACheckApplies(t *testing.T) {
	l := NewCertIssuedByCompromisedEmudhraCA()
	for algo, want := range map[x509.SignatureAlgorithm]bool{
		x509.ECDSAWithSHA256:           true,
		x509.ECDSAWithSHA384:           true,
		x509.SHA256WithRSA:             false,
		x509.UnknownSignatureAlgorithm: false,
	} {
		if got := l.CheckApplies(&x509.Certificate{SignatureAlgorithm: algo}); got != want {
			t.Errorf("%s: expected CheckApplies %v, got %v", algo, want, got)
		}
	}
}
