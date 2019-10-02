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
Section 5.2 - Forbidden and Required Practices
CAs MUST NOT issue certificates that have:
- invalid public keys (e.g., RSA certificates with public exponent equal to 1);
********************************************************************/

package lints

import (
	"crypto/rsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type allowedRSAKey struct{}

func (l *allowedRSAKey) Initialize() error {
	return nil
}

func (l *allowedRSAKey) CheckApplies(c *x509.Certificate) bool {
	switch c.SignatureAlgorithm {
	case x509.SHA1WithRSA, x509.SHA256WithRSA, x509.SHA384WithRSA, x509.SHA512WithRSA, x509.SHA256WithRSAPSS, x509.SHA384WithRSAPSS, x509.SHA512WithRSAPSS:
		return true
	}

	return false
}

func (l *allowedRSAKey) Execute(c *x509.Certificate) *LintResult {
	pubKey := c.PublicKey.(*rsa.PublicKey)
	bitLen := pubKey.N.BitLen()

	if (bitLen%8) != 0 || bitLen < 2048 || pubKey.E == 1 {
		return &LintResult{Status: Error}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_mp_allowed_rsa_keys",
		Description:   "RSA keys whose modulus size in bits is divisible by 8, is at least 2048 and public exponent is not equal to 1",
		Citation:      "Mozilla Root Store Policy / Section 5.3",
		Source:        MozillaRootStorePolicy,
		EffectiveDate: util.MozillaPolicy24Date,
		Lint:          &allowedRSAKey{},
	})
}
