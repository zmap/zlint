/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

package cabf_smime_br

import (
	"crypto/ecdsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ecpublickey_key_usages",
		Description:   "For signing only, bit positions SHALL be set for digitalSignature and MAY be set for nonRepudiation. For key management only, bit positions SHALL be set for keyEncipherment.For dual use, bit positions SHALL be set for digitalSignature and keyEncipherment and MAY be set for nonRepudiation.",
		Citation:      "7.1.2.3.e",
		Source:        lint.CABFSMIMEBaselineRequirements,
		EffectiveDate: util.CABF_SMIME_BRs_1_0_0_Date,
		Lint:          NewECPublicKeyKeyUsages,
	})
}

type ecPublicKeyKeyUsages struct{}

func NewECPublicKeyKeyUsages() lint.LintInterface {
	return &ecPublicKeyKeyUsages{}
}

func (l *ecPublicKeyKeyUsages) CheckApplies(c *x509.Certificate) bool {
	if !(util.IsSubscriberCert(c) && util.IsSMIMEBRCertificate(c) && util.IsExtInCert(c, util.KeyUsageOID)) {
		return false
	}

	_, ok := c.PublicKey.(*ecdsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.ECDSA
}

func (l *ecPublicKeyKeyUsages) Execute(c *x509.Certificate) *lint.LintResult {
	certType := 0
	if c.KeyUsage&x509.KeyUsageDigitalSignature != 0 {
		certType |= 1
	}
	if c.KeyUsage&x509.KeyUsageKeyAgreement != 0 {
		certType |= 2
	}

	switch certType {
	case 1:
		// signing only
		mask := 0x1FF ^ (x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment)
		if c.KeyUsage&mask != 0 {
			return &lint.LintResult{Status: lint.Error}
		}
	case 2:
		// key management only
		mask := 0x1FF ^ (x509.KeyUsageKeyAgreement | x509.KeyUsageEncipherOnly | x509.KeyUsageDecipherOnly)
		if c.KeyUsage&mask != 0 {
			return &lint.LintResult{Status: lint.Error}
		}
	case 3:
		// dual use
		mask := 0x1FF ^ (x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageKeyAgreement | x509.KeyUsageEncipherOnly | x509.KeyUsageDecipherOnly)
		if c.KeyUsage&mask != 0 {
			return &lint.LintResult{Status: lint.Error}
		}
	default:
		return &lint.LintResult{Status: lint.Error}
	}

	return &lint.LintResult{Status: lint.Pass}
}