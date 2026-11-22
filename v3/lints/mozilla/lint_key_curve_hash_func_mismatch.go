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

package mozilla

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"

	"fmt"
	"time"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_key_curve_hash_func_mismatch",
			Description:   "Checks that the ECC curve used for signing is allowed in combination with hash function",
			Citation:      "Mozilla Root Store Policy, section 5.1",
			Source:        lint.MozillaRootStorePolicy,
			EffectiveDate: util.MozillaPolicy25Date,
		},
		Lint: NewKeyCurveHashFuncMismatch,
	})
}

type KeyCurveHashFuncMismatch struct{}

func NewKeyCurveHashFuncMismatch() lint.LintInterface {
	return &KeyCurveHashFuncMismatch{}
}

func (l *KeyCurveHashFuncMismatch) CheckApplies(c *x509.Certificate) bool {
	return c.SignatureAlgorithm == x509.ECDSAWithSHA256 ||
		c.SignatureAlgorithm == x509.ECDSAWithSHA384 ||
		c.SignatureAlgorithm == x509.ECDSAWithSHA512
}

// Minimum and maximum expected length of the DER encoding of an ECDSA signature
// computed using various ECDSA-based algorithms, assuming the signature is created
// using the ECDSA curve required by Mozilla Root Store Policy for each hash function.
// These lengths refer to the DER encoding of the entire SEQUENCE that contains the two
// components (r, s) of the signature (per RFC 3279 §2.2.3).
// The maximum lengths are very specific values which only depends on the curve.
// The minimum lengths are based on heuristics (lower values are highly unlikely).
var allowedECDSASigLengths = map[x509.SignatureAlgorithm]struct{ Min, Max int }{
	// Assuming the ECDSA key is P256
	x509.ECDSAWithSHA256: {68, 72},
	// Assuming the ECDSA key is P384
	x509.ECDSAWithSHA384: {100, 104},
	// Assuming the ECDSA key is P521
	x509.ECDSAWithSHA512: {137, 139},
}

func (l *KeyCurveHashFuncMismatch) Execute(c *x509.Certificate) *lint.LintResult {

	// Support for P521 was first introduced in Mozilla Root Store Policy v3.0
	if c.SignatureAlgorithm == x509.ECDSAWithSHA512 && c.NotBefore.Before(util.MozillaPolicy30Date) {
		return &lint.LintResult{
			Status: lint.Error,
			Details: "Signature algorithm ECDSAWithSHA512 not allowed before " +
				util.MozillaPolicy30Date.Format(time.DateOnly),
		}
	}

	if len(c.Signature) > allowedECDSASigLengths[c.SignatureAlgorithm].Max {
		return &lint.LintResult{
			Status: lint.Error,
			Details: fmt.Sprintf("Signature was made with a curve not permitted with algorithm %v",
				c.SignatureAlgorithm),
		}
	}
	if len(c.Signature) < allowedECDSASigLengths[c.SignatureAlgorithm].Min {
		return &lint.LintResult{
			Status: lint.Error,
			Details: fmt.Sprintf("Signature was probably made with a curve not permitted with algorithm %v",
				c.SignatureAlgorithm),
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
