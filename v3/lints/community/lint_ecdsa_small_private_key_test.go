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
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestEcdsaSmallPrivateKey(t *testing.T) {
	data := []struct {
		input    string
		expected lint.LintStatus
	}{
		{"ecdsaSmallKeyP256D5.pem", lint.Error},
		{"ecdsaSmallKeyP384D200CA.pem", lint.Error},
		// d = n - 5 is just as recoverable as d = 5.
		{"ecdsaSmallKeyP256NegatedD5.pem", lint.Error},
		// d = 2^16 is the largest scalar searched by default, 2^16 + 1 is the first beyond.
		{"ecdsaSmallKeyP256Boundary.pem", lint.Error},
		{"ecdsaSmallKeyP256D50000.pem", lint.Error},
		{"ecdsaSmallKeyP256NegatedD50000.pem", lint.Error},
		{"ecdsaSmallKeyP256NegatedBoundary.pem", lint.Error},
		{"ecdsaSmallKeyP256AboveBoundary.pem", lint.Pass},
		{"ecdsaSmallKeyP256NegatedAboveBoundary.pem", lint.Pass},
		{"ecdsaSmallKeyP384Random.pem", lint.Pass},
		// Every compromised eMudhra intermediate, code signing CAs included from bug 2083449, as issued by the emSign roots.
		{"emudhraEvG3BIntermediate.pem", lint.Error},
		{"emudhraIntermediateEvG3C.pem", lint.Error},
		{"emudhraIntermediateOvG3B.pem", lint.Error},
		{"emudhraIntermediateOvG3C.pem", lint.Error},
		{"emudhraIntermediateDvG3B.pem", lint.Error},
		{"emudhraIntermediateDvG3C.pem", lint.Error},
		{"emudhraIntermediateCsG3B.pem", lint.Error},
		{"emudhraIntermediateCsG3C.pem", lint.Error},
		// ...and the same keys as cross-signed by Sectigo.
		{"emudhraSectigoCrossEvG3B.pem", lint.Error},
		{"emudhraSectigoCrossEvG3C.pem", lint.Error},
		{"emudhraSectigoCrossOvG3B.pem", lint.Error},
		{"emudhraSectigoCrossOvG3C.pem", lint.Error},
		{"emudhraSectigoCrossDvG3B.pem", lint.Error},
		{"emudhraSectigoCrossDvG3C.pem", lint.Error},
		{"emudhraSectigoCrossCsG3B.pem", lint.Error},
		{"emudhraSectigoCrossCsG3C.pem", lint.Error},
		{"emudhraEvG3BSubscriber.pem", lint.Pass},
		{"RSASHA1Good.pem", lint.NA},
		// A scalar that lands exactly on a giant step (7 * 8193), which is
		// found through the point at infinity.
		{"ecdsaSmallKeyP256StrideMultiple.pem", lint.Error},
		{"ecdsaSmallKeyP521D3.pem", lint.Error},
		// Every kind of CA, whatever its extendedKeyUsage, must be checked.
		{"ecdsaSmallKeyP256SelfSignedRoot.pem", lint.Error},
		{"ecdsaSmallKeyP256NoEKUCA.pem", lint.Error},
		{"ecdsaSmallKeyP256ServerAuthCA.pem", lint.Error},
		{"ecdsaSmallKeyP256CodeSigningCA.pem", lint.Error},
		{"ecdsaSmallKeyP256EmailProtectionCA.pem", lint.Error},
		// There is no effective date, so even a very old certificate is checked.
		{"ecdsaSmallKeyP256NotBefore2021.pem", lint.Error},
	}
	for _, d := range data {
		t.Run(d.input, func(t *testing.T) {
			out := test.TestLint("e_ecdsa_small_private_key", d.input)
			if out.Status != d.expected {
				t.Errorf("expected %s, got %s (%s)", d.expected, out.Status, out.Details)
			}
		})
	}
}

func TestEcdsaSmallPrivateKeyBitsOutOfRange(t *testing.T) {
	for _, bits := range []int{-1, 0, 1, maxSmallKeyBits + 1} {
		l := &ecdsaSmallPrivateKey{Bits: bits}
		c := test.ReadTestCert("ecdsaSmallKeyP256D5.pem")
		if got := l.Execute(c).Status; got != lint.Fatal {
			t.Errorf("Bits=%d: expected %s, got %s", bits, lint.Fatal, got)
		}
	}
}

func TestEcdsaSmallPrivateKeyMalformedKeyIsNotApplicable(t *testing.T) {
	c := elliptic.P256()
	// (1, 1) is not on P-256.
	offCurve := &ecdsa.PublicKey{Curve: c, X: big.NewInt(1), Y: big.NewInt(1)}
	for name, cert := range map[string]*x509.Certificate{
		"off curve":     {PublicKeyAlgorithm: x509.ECDSA, PublicKey: offCurve},
		"missing point": {PublicKeyAlgorithm: x509.ECDSA, PublicKey: &ecdsa.PublicKey{Curve: c}},
		"missing curve": {PublicKeyAlgorithm: x509.ECDSA, PublicKey: &ecdsa.PublicKey{X: big.NewInt(1), Y: big.NewInt(1)}},
		"not ecdsa":     {PublicKeyAlgorithm: x509.RSA},
	} {
		if (&ecdsaSmallPrivateKey{Bits: 16}).CheckApplies(cert) {
			t.Errorf("%s: expected CheckApplies to be false", name)
		}
	}
}

// TestSmallDiscreteLog checks the search directly on every supported curve:
// the edges of the search windows, exact multiples of the giant step, both signs
// of the key, and values just outside the range.
func TestSmallDiscreteLog(t *testing.T) {
	const bits = 16
	m := int64(1) << (bits/2 + 4)
	stride := 2*m + 1
	inRange := []int64{1, 2, m - 1, m, m + 1, m + 2, stride - 1, stride, stride + 1, 2 * stride,
		3*stride - m - 1, 3*stride - m, 7 * stride, 1<<bits - 1, 1 << bits}
	for d := int64(1); d <= 1<<bits; d += 4999 {
		inRange = append(inRange, d)
	}
	outOfRange := []int64{1<<bits + 1, 1<<bits + 2, 8 * stride, 100000}
	for _, curve := range []elliptic.Curve{elliptic.P256(), elliptic.P384(), elliptic.P521()} {
		n := curve.Params().N
		pub := func(d int64, negate bool) *ecdsa.PublicKey {
			k := big.NewInt(d)
			if negate {
				k.Sub(n, k)
			}
			x, y := curve.ScalarBaseMult(k.Bytes())
			return &ecdsa.PublicKey{Curve: curve, X: x, Y: y}
		}
		for _, negate := range []bool{false, true} {
			for _, d := range inRange {
				if got := smallDiscreteLog(pub(d, negate), bits); got == nil || got.Int64() != d {
					t.Errorf("%s d=%d negated=%v: got %v", curve.Params().Name, d, negate, got)
				}
			}
			for _, d := range outOfRange {
				if got := smallDiscreteLog(pub(d, negate), bits); got != nil {
					t.Errorf("%s d=%d negated=%v: expected no result, got %v", curve.Params().Name, d, negate, got)
				}
			}
		}
	}
}
