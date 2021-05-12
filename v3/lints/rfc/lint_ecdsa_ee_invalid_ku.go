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

package rfc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type ecdsaInvalidKU struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ecdsa_ee_invalid_ku",
		Description:   "Key usage values keyEncipherment or dataEncipherment MUST NOT be present in certificates with ECDSA public keys",
		Citation:      "RFC 8813 Section 3",
		Source:        lint.RFC8813,
		EffectiveDate: util.RFC8813Date,
		Lint:          &ecdsaInvalidKU{},
	})
}

// Initialize is a no-op for this lint.
func (l *ecdsaInvalidKU) Initialize() error {
	return nil
}

// CheckApplies returns true when the certificate has an ECDSA public key and a key usage extension.
func (l *ecdsaInvalidKU) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.ECDSA && util.IsExtInCert(c, util.KeyUsageOID)
}

// Execute returns an Error level lint.LintResult if the ECDSA certificate
// being linted has the following Key Usage bits set: keyEncipherment or dataEncipherment.
func (l *ecdsaInvalidKU) Execute(c *x509.Certificate) *lint.LintResult {
	// RFC 8813, Section 3 "Updates to Section 3" reads:
	//
	// If the keyUsage extension is present in a certificate that indicates
	// id-ecPublicKey in SubjectPublicKeyInfo, then the following values
	// MUST NOT be present:
	//
	//    keyEncipherment; and
	//    dataEncipherment.
	forbiddenKUs := map[x509.KeyUsage]bool{
		x509.KeyUsageKeyEncipherment:  true,
		x509.KeyUsageDataEncipherment: true,
	}

	var invalidKUs []string
	for ku, kuName := range util.KeyUsageToString {
		if c.KeyUsage&ku != 0 {
			if forbiddenKUs[ku] {
				invalidKUs = append(invalidKUs, kuName)
			}
		}
	}

	if len(invalidKUs) > 0 {
		// Sort the invalid KUs to allow consistent ordering of Details messages for
		// unit testing
		sort.Strings(invalidKUs)
		return &lint.LintResult{
			Status: lint.Error,
			Details: fmt.Sprintf(
				"Certificate had invalid key usage(s): %s",
				strings.Join(invalidKUs, ", ")),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
