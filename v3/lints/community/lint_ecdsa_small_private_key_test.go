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
	"fmt"
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
		// d = k^-1 mod n for a small k: k*pub = G, so d is just as recoverable.
		{"ecdsaSmallKeyP256InverseK5.pem", lint.Error},
		{"ecdsaSmallKeyP256NegatedInverseK5.pem", lint.Error},
		// k = 2^16 is the largest inverse searched by default, 2^16 + 1 is the first beyond.
		{"ecdsaSmallKeyP256InverseBoundary.pem", lint.Error},
		{"ecdsaSmallKeyP256NegatedInverseBoundary.pem", lint.Error},
		{"ecdsaSmallKeyP256InverseAboveBoundary.pem", lint.Pass},
		{"ecdsaSmallKeyP384InverseK50000.pem", lint.Error},
		{"ecdsaSmallKeyP521NegatedInverseK3.pem", lint.Error},
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
	c := test.ReadTestCert("ecdsaSmallKeyP256D5.pem")
	for _, l := range []ecdsaSmallPrivateKey{
		{Bits: -1}, {Bits: 0}, {Bits: 1}, {Bits: maxSmallKeyBits + 1},
		{Bits: 16, InverseBits: -1}, {Bits: 16, InverseBits: 1}, {Bits: 16, InverseBits: maxInverseKeyBits + 1},
	} {
		if got := l.Execute(c).Status; got != lint.Fatal {
			t.Errorf("Bits=%d InverseBits=%d: expected %s, got %s", l.Bits, l.InverseBits, lint.Fatal, got)
		}
	}
}

func TestEcdsaSmallPrivateKeyInverseDisabled(t *testing.T) {
	l := &ecdsaSmallPrivateKey{Bits: 16, InverseBits: 0}
	if got := l.Execute(test.ReadTestCert("ecdsaSmallKeyP256InverseK5.pem")).Status; got != lint.Pass {
		t.Errorf("InverseBits=0 on an inverse key: expected %s, got %s", lint.Pass, got)
	}
	// Disabling the inverse search does not disable the small scalar search.
	if got := l.Execute(test.ReadTestCert("ecdsaSmallKeyP256D5.pem")).Status; got != lint.Error {
		t.Errorf("InverseBits=0 on a small key: expected %s, got %s", lint.Error, got)
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

// TestSmallDiscreteLog checks both searches directly on every supported curve:
// the edges of the search windows, exact multiples of the giant step, both signs
// of the key, and values just outside the range.
func TestSmallDiscreteLog(t *testing.T) {
	const bits = 16
	for _, search := range []struct {
		name string
		// m is the number of baby steps the search uses for bits.
		m int64
		// scalar is the private key the search finds as k.
		scalar func(k int64, n *big.Int) *big.Int
		search func(pub *ecdsa.PublicKey, bits int) *big.Int
	}{
		{"small", 1 << (bits/2 + 4), func(k int64, _ *big.Int) *big.Int { return big.NewInt(k) }, smallDiscreteLog},
		{"inverse", 1 << (bits / 2), func(k int64, n *big.Int) *big.Int { return new(big.Int).ModInverse(big.NewInt(k), n) }, smallInverseDiscreteLog},
	} {
		m := search.m
		stride := 2*m + 1
		limit := int64(1) << bits
		inRange := []int64{1, 2, m - 1, m, m + 1, m + 2, stride - 1, stride, stride + 1, 2 * stride,
			3*stride - m - 1, 3*stride - m, 7 * stride, limit - 1, limit}
		for k := int64(1); k <= limit; k += 4999 {
			inRange = append(inRange, k)
		}
		// (limit/stride + 1) * stride is the first giant step past the last window.
		outOfRange := []int64{limit + 1, limit + 2, (limit/stride + 1) * stride, 100000}
		for _, curve := range []elliptic.Curve{elliptic.P256(), elliptic.P384(), elliptic.P521()} {
			n := curve.Params().N
			pub := func(k int64, negate bool) *ecdsa.PublicKey {
				d := search.scalar(k, n)
				if negate {
					d.Sub(n, d)
				}
				x, y := curve.ScalarBaseMult(d.Bytes())
				return &ecdsa.PublicKey{Curve: curve, X: x, Y: y}
			}
			for _, negate := range []bool{false, true} {
				for _, k := range inRange {
					if got := search.search(pub(k, negate), bits); got == nil || got.Int64() != k {
						t.Errorf("%s %s k=%d negated=%v: got %v", search.name, curve.Params().Name, k, negate, got)
					}
				}
				for _, k := range outOfRange {
					if got := search.search(pub(k, negate), bits); got != nil {
						t.Errorf("%s %s k=%d negated=%v: expected no result, got %v", search.name, curve.Params().Name, k, negate, got)
					}
				}
			}
		}
	}
}

func BenchmarkEcdsaSmallPrivateKey(b *testing.B) {
	c := test.ReadTestCert("ecdsaSmallKeyP384Random.pem")
	for _, inverseBits := range []int{0, 16} {
		b.Run(fmt.Sprintf("InverseBits=%d", inverseBits), func(b *testing.B) {
			l := &ecdsaSmallPrivateKey{Bits: 16, InverseBits: inverseBits}
			for b.Loop() {
				l.Execute(c)
			}
		})
	}
}
