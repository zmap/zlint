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

package rfc

import (
	"sort"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type mldsaKeyUsages struct{}

/*******************************************************************************************************
RFC 9881, Section 5:
If the keyUsage extension is present in a certificate that includes id-ml-dsa-* (where * is 44, 65,
or 87) in the SubjectPublicKeyInfo, then the subject public key can only be used for verifying digital
signatures on certificates or CRLs, or those used to provide a non-repudiation service that protects
against the signing entity falsely denying some action. This means that the keyUsage extension MUST
have at least one of the following bits set:

  * digitalSignature
  * nonRepudiation
  * keyCertSign
  * cRLSign

ML-DSA subject public keys cannot be used to establish keys or encrypt data, so the keyUsage extension
MUST NOT have any of the following bits set:

  * keyEncipherment
  * dataEncipherment
  * keyAgreement
  * encipherOnly
  * decipherOnly
*******************************************************************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_mldsa_key_usages",
			Description:   "ML-DSA: keyUsage MUST have at least one of digitalSignature, nonRepudiation, keyCertSign or cRLSign set and MUST NOT have keyEncipherment, dataEncipherment, keyAgreement, encipherOnly or decipherOnly set",
			Citation:      "RFC 9881, Section 5",
			Source:        lint.RFC9881,
			EffectiveDate: util.RFC9881Date,
		},
		Lint: NewMldsaKeyUsages,
	})
}

func NewMldsaKeyUsages() lint.LintInterface {
	return &mldsaKeyUsages{}
}

// CheckApplies returns true when the certificate has an ML-DSA public key and a key usage extension.
func (l *mldsaKeyUsages) CheckApplies(c *x509.Certificate) bool {
	return util.IsMLDSAOID(c.PublicKeyAlgorithmOID) && util.HasKeyUsageOID(c)
}

func (l *mldsaKeyUsages) Execute(c *x509.Certificate) *lint.LintResult {
	var invalidKUs []string
	for _, ku := range []x509.KeyUsage{
		x509.KeyUsageKeyEncipherment,
		x509.KeyUsageDataEncipherment,
		x509.KeyUsageKeyAgreement,
		x509.KeyUsageEncipherOnly,
		x509.KeyUsageDecipherOnly,
	} {
		if util.HasKeyUsage(c, ku) {
			invalidKUs = append(invalidKUs, util.KeyUsageToString[ku])
		}
	}

	if len(invalidKUs) > 0 {
		// Sort the invalid KUs to allow consistent ordering of Details messages for
		// unit testing
		sort.Strings(invalidKUs)
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "Certificate contains invalid key usage(s): " + strings.Join(invalidKUs, ", "),
		}
	}

	requiredKUs := x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	if c.KeyUsage&requiredKUs == 0 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "Certificate does not contain any of the key usages: KeyUsageDigitalSignature, KeyUsageContentCommitment, KeyUsageCertSign, KeyUsageCRLSign",
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
