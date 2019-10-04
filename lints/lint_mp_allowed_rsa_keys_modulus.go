/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

/********************************************************************
Section 5.1 - Algorithms
RSA keys whose modulus size in bits is divisible by 8, and is at least 2048.
********************************************************************/

package lints

import (
	"crypto/rsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type allowedRSAKeyModulus struct{}

func (l *allowedRSAKeyModulus) Initialize() error {
	return nil
}

func (l *allowedRSAKeyModulus) CheckApplies(c *x509.Certificate) bool {
	if c.PublicKeyAlgorithm == x509.RSA {
		return true
	}

	return false
}

func (l *allowedRSAKeyModulus) Execute(c *x509.Certificate) *LintResult {
	pubKey, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		return &LintResult{Status: Fatal, Details: "certificate public key was not an RSA public key"}
	}

	bitLen := pubKey.N.BitLen()
	if (bitLen%8) != 0 || bitLen < 2048 {
		return &LintResult{Status: Error}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_mp_allowed_rsa_keys_modulus",
		Description:   "RSA keys whose modulus size in bits is divisible by 8, is at least 2048 and public exponent is not equal to 1",
		Citation:      "Mozilla Root Store Policy / Section 5.1",
		Source:        MozillaRootStorePolicy,
		EffectiveDate: util.MozillaPolicy24Date,
		Lint:          &allowedRSAKeyModulus{},
	})
}
