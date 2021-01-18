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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type KUAndEKUInconsistent struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_key_usage_and_extended_key_usage_inconsistent",
		Description:   "The certificate MUST only be used for a purpose consistent with both key usage extension and extended key usage extension.",
		Citation:      "RFC 5280, Section 4.2.1.12.",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &KUAndEKUInconsistent{},
	})
}

func (l *KUAndEKUInconsistent) Initialize() error {
	return nil
}

// CheckApplies returns true when the certificate contains both a key usage
// extension and an extended key usage extension.
func (l *KUAndEKUInconsistent) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.EkuSynOid) && util.IsExtInCert(c, util.KeyUsageOID)
}

// Execute returns an Error level lint.LintResult if the purposes of the certificate
// being linted is not consistent with both extensions.
func (l *KUAndEKUInconsistent) Execute(c *x509.Certificate) *lint.LintResult {
	for _, extKeyUsage := range c.ExtKeyUsage {
		switch extKeyUsage {
		case x509.ExtKeyUsageServerAuth:
			if !serverAuth[c.KeyUsage] {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageClientAuth:
			if !clientAuth[c.KeyUsage] {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageCodeSigning:
			if !codeSigning[c.KeyUsage] {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageEmailProtection:
			if !emailProtection[c.KeyUsage] {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageTimeStamping:
			if !timeStamping[c.KeyUsage] {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageOcspSigning:
			if !ocspSigning[c.KeyUsage] {
				return &lint.LintResult{Status: lint.Error}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

// Variable containing the consistent KU combinations with Server Authentication EKU:
//  RFC 5280 4.2.1.12 on KU consistency with Server Authentication EKU:
//    -- TLS WWW server authentication
//    -- Key usage bits that may be consistent: digitalSignature,
//    -- keyEncipherment or keyAgreement
var serverAuth = map[x509.KeyUsage]bool{
	x509.KeyUsageDigitalSignature: true,
	x509.KeyUsageKeyEncipherment:  true,
	x509.KeyUsageKeyAgreement:     true,
}

// Variable containing the consistent KU combinations with Client Authentication EKU:
// 	RFC 5280 4.2.1.12 on KU consistency with Client Authentication EKU:
//    -- TLS WWW client authentication
//    -- Key usage bits that may be consistent: digitalSignature
//    -- and/or keyAgreement
var clientAuth = map[x509.KeyUsage]bool{
	x509.KeyUsageDigitalSignature:                             true,
	x509.KeyUsageKeyAgreement:                                 true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement: true,
}

// Variable containing the consistent KU combinations with Code Signing EKU:
// 	RFC 5280 4.2.1.12 on KU consistency with Code Signing EKU:
//   -- Signing of downloadable executable code
//   -- Key usage bits that may be consistent: digitalSignature
var codeSigning = map[x509.KeyUsage]bool{
	x509.KeyUsageDigitalSignature: true,
}

// Variable containing the consistent KU combinations with Email Protection EKU:
// 	RFC 5280 4.2.1.12 on KU consistency with Email Protection EKU:
// 	  -- Email protection
//    -- Key usage bits that may be consistent: digitalSignature,
//    -- nonRepudiation, and/or (keyEncipherment or keyAgreement)
//  Note: Recent editions of X.509 have renamed nonRepudiation bit to contentCommitment
var emailProtection = map[x509.KeyUsage]bool{
	x509.KeyUsageDigitalSignature:                                                                 true,
	x509.KeyUsageContentCommitment:                                                                true,
	x509.KeyUsageKeyEncipherment:                                                                  true,
	x509.KeyUsageKeyAgreement:                                                                     true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment:                                true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment:                                  true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement:                                     true,
	x509.KeyUsageContentCommitment | x509.KeyUsageKeyEncipherment:                                 true,
	x509.KeyUsageContentCommitment | x509.KeyUsageKeyAgreement:                                    true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageKeyEncipherment: true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageKeyAgreement:    true,
}

// Variable containing the consistent KU combinations with Time Stamping EKU:
// 	RFC 5280 4.2.1.12 on KU consistency with Time Stamping EKU:
// 	  -- Binding the hash of an object to a time
//    -- Key usage bits that may be consistent: digitalSignature
//    -- and/or nonRepudiation
//  Note: Recent editions of X.509 have renamed nonRepudiation bit to contentCommitment
var timeStamping = map[x509.KeyUsage]bool{
	x509.KeyUsageDigitalSignature:                                  true,
	x509.KeyUsageContentCommitment:                                 true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment: true,
}

// Variable containing the consistent KU combinations with Ocsp Signing EKU:
// 	RFC 5280 4.2.1.12 on KU consistency with Ocsp Signing EKU:
// 	  -- Signing OCSP responses
//    -- Key usage bits that may be consistent: digitalSignature
//    -- and/or nonRepudiation
//  Note: Recent editions of X.509 have renamed nonRepudiation bit to contentCommitment
var ocspSigning = map[x509.KeyUsage]bool{
	x509.KeyUsageDigitalSignature:                                  true,
	x509.KeyUsageContentCommitment:                                 true,
	x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment: true,
}
