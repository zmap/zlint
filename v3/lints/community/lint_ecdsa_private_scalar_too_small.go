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
	"bytes"
	"crypto/ecdsa"
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "w_ecdsa_private_scalar_too_small",
			Description:   "The private scalar can be trivially computed from the public key",
			Citation:      "https://bugzilla.mozilla.org/show_bug.cgi?id=2083449",
			Source:        lint.Community,
			EffectiveDate: util.ZeroDate,
		},
		Lint: NewEcdsaPrivateScalarTooSmall,
	})
}

type ecdsaPrivateScalarTooSmall struct{}

func NewEcdsaPrivateScalarTooSmall() lint.LintInterface {
	return &ecdsaPrivateScalarTooSmall{}
}

func (l *ecdsaPrivateScalarTooSmall) CheckApplies(c *x509.Certificate) bool {
	fmt.Println(c.PublicKeyAlgorithm)
	return c.PublicKeyAlgorithm == x509.ECDSA
}

func (l *ecdsaPrivateScalarTooSmall) Execute(c *x509.Certificate) *lint.LintResult {
	// Extract the key from the certificate.
	var key *ecdsa.PublicKey
	switch t := c.PublicKey.(type) {
	case *x509.AugmentedECDSA:
		fmt.Println("here")
		key = t.Pub
	case *ecdsa.PublicKey:
		fmt.Println("there")
		key = t
	}

	// Initialize our x and y coordinates to the curve's base point.
	curve := key.Curve
	params := curve.Params()
	x, y := params.Gx, params.Gy

	// Check if the key's point is a small multiple of the curve's base point.
	for i := 1; i <= 255; i++ {
		if bytes.Equal(key.X.Bytes(), x.Bytes()) {
			return &lint.LintResult{
				Status:  lint.Warn,
				Details: fmt.Sprintf("Key has trivially computable private scalar %d", i)}
		}

		// Move to the next base point. Doing this iteratively via addition is much
		// faster than recomputing each point by multiplying the base point by i.
		x, y = curve.Add(x, y, params.Gx, params.Gy)
	}

	return &lint.LintResult{Status: lint.Pass}
}
