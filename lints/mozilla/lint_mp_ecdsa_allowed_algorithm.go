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
ECDSA keys using one of the following curve-hash pairs: P‐256 with SHA-256, P‐384 with SHA-384
********************************************************************/

package lints

import (
	"encoding/asn1"
	"errors"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ecdsaSignature struct {
	R, S *big.Int
}

func getSigningKeySize(cert *x509.Certificate) (int, error) {
	sig := new(ecdsaSignature)
	if _, err := asn1.Unmarshal(cert.Signature, sig); err != nil {
		return -1, err
	}

	rsize := sig.R.BitLen()
	ssize := sig.S.BitLen()

	switch {
	case rsize <= 112 && ssize <= 112:
		return 112, nil
	case rsize <= 128 && ssize <= 128:
		return 128, nil
	case rsize <= 160 && ssize <= 160:
		return 160, nil
	case rsize <= 192 && ssize <= 192:
		return 192, nil
	case rsize <= 224 && ssize <= 224:
		return 224, nil
	case rsize <= 239 && ssize <= 239:
		return 239, nil
	case rsize <= 256 && ssize <= 256:
		return 256, nil
	case rsize <= 320 && ssize <= 320:
		return 320, nil
	case rsize <= 384 && ssize <= 384:
		return 384, nil
	case rsize <= 512 && ssize <= 512:
		return 512, nil
	case rsize <= 521 && ssize <= 521:
		return 521, nil
	}

	return -1, errors.New("cannot identify signing ECDSA key length")
}

type ecdsaAllowedAlgorithm struct{}

func (l *ecdsaAllowedAlgorithm) Initialize() error {
	return nil
}

func (l *ecdsaAllowedAlgorithm) CheckApplies(c *x509.Certificate) bool {
	if c.SignatureAlgorithm == x509.ECDSAWithSHA1 || c.SignatureAlgorithm == x509.ECDSAWithSHA256 ||
		c.SignatureAlgorithm == x509.ECDSAWithSHA384 || c.SignatureAlgorithm == x509.ECDSAWithSHA512 {
		return true
	}

	return false
}

func (l *ecdsaAllowedAlgorithm) Execute(c *x509.Certificate) *LintResult {
	signKeySize, err := getSigningKeySize(c)
	if err != nil {
		return &LintResult{
			Status:  Fatal,
			Details: "cannot identify signing ECDSA key length",
		}
	}

	switch {
	case c.SignatureAlgorithm == x509.ECDSAWithSHA256 && signKeySize == 256:
		return &LintResult{
			Status:  Pass,
			Details: "Detected ECDSAWithSHA256 and 256 bit signing key.",
		}
	case c.SignatureAlgorithm == x509.ECDSAWithSHA384 && signKeySize == 384:
		return &LintResult{
			Status:  Pass,
			Details: "Detected ECDSAWithSHA384 and 384 bit signing key.",
		}
	}

	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_mp_ecdsa_allowed_algorithm",
		Description:   "ECDSA keys using one of the following curve-hash pairs: P‐256 with SHA-256, P‐384 with SHA-384",
		Citation:      "Mozilla Root Store Policy / Section 5.1",
		Source:        MozillaRootStorePolicy,
		EffectiveDate: util.MozillaPolicy24Date,
		Lint:          &ecdsaAllowedAlgorithm{},
	})
}
