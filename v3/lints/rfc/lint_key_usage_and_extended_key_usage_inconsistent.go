/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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
			if !(CheckConsistencyWithEKUServerAuth(c)) {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageClientAuth:
			if !(CheckConsistencyWithEKUClientAuth(c)) {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageCodeSigning:
			if !(CheckConsistencyWithEKUCodeSigning(c)) {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageEmailProtection:
			if !(CheckConsistencyWithEKUEmailProtection(c)) {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageTimeStamping:
			if !(CheckConsistencyWithEKUTimeStamping(c)) {
				return &lint.LintResult{Status: lint.Error}
			}
		case x509.ExtKeyUsageOcspSigning:
			if !(CheckConsistencyWithEKUOcspSigning(c)) {
				return &lint.LintResult{Status: lint.Error}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

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

//  CheckConsistencyWithEKU* functions return false if the certificate being linted
//  has inconsistent Key Usage bits set with a specific Extended Key Usage

//CheckConsistencyWithEKUServerAuth checks if KU bits are consistent with Server Authentication EKU bit
func CheckConsistencyWithEKUServerAuth(c *x509.Certificate) bool {
	//  RFC 5280 4.2.1.12 on KU consistency with Server Authentication EKU:
	//    -- TLS WWW server authentication
	//    -- Key usage bits that may be consistent: digitalSignature,
	//    -- keyEncipherment or keyAgreement

	allowedCombinations := []x509.KeyUsage{
		x509.KeyUsageDigitalSignature,
		x509.KeyUsageKeyEncipherment,
		x509.KeyUsageKeyAgreement,
	}

	return containsCombination(c.KeyUsage, allowedCombinations)
}

//CheckConsistencyWithEKUClientAuth checks if KU bits are consistent with Client Authentication EKU bit
func CheckConsistencyWithEKUClientAuth(c *x509.Certificate) bool {
	// 	RFC 5280 4.2.1.12 on KU consistency with Client Authentication EKU:
	//    -- TLS WWW client authentication
	//    -- Key usage bits that may be consistent: digitalSignature
	//    -- and/or keyAgreement

	allowedCombinations := []x509.KeyUsage{
		x509.KeyUsageDigitalSignature,
		x509.KeyUsageKeyAgreement,
		x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement}

	return containsCombination(c.KeyUsage, allowedCombinations)
}

//CheckConsistencyWithEKUCodeSigning checks if KU bits are consistent with Code Signing EKU bit
func CheckConsistencyWithEKUCodeSigning(c *x509.Certificate) bool {
	// 	RFC 5280 4.2.1.12 on KU consistency with Code Signing EKU:
	//   -- Signing of downloadable executable code
	//   -- Key usage bits that may be consistent: digitalSignature

	allowedCombinations := []x509.KeyUsage{
		x509.KeyUsageDigitalSignature}

	return containsCombination(c.KeyUsage, allowedCombinations)
}

//CheckConsistencyWithEKUEmailProtection checks if KU bits are consistent with Email Protection EKU bit
func CheckConsistencyWithEKUEmailProtection(c *x509.Certificate) bool {
	// 	RFC 5280 4.2.1.12 on KU consistency with Email Protection EKU:
	// 	  -- Email protection
	//    -- Key usage bits that may be consistent: digitalSignature,
	//    -- nonRepudiation, and/or (keyEncipherment or keyAgreement)
	//  Note: Recent editions of X.509 have renamed nonRepudiation bit to contentCommitment

	allowedCombinations := []x509.KeyUsage{
		x509.KeyUsageDigitalSignature,
		x509.KeyUsageContentCommitment,
		x509.KeyUsageKeyEncipherment,
		x509.KeyUsageKeyAgreement,
		x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment,
		x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement,
		x509.KeyUsageContentCommitment | x509.KeyUsageKeyEncipherment,
		x509.KeyUsageContentCommitment | x509.KeyUsageKeyAgreement,
		x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageKeyEncipherment,
		x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageKeyAgreement,
	}

	return containsCombination(c.KeyUsage, allowedCombinations)
}

//CheckConsistencyWithEKUTimeStamping checks if KU bits are consistent with Time Stamping EKU bit
func CheckConsistencyWithEKUTimeStamping(c *x509.Certificate) bool {
	// 	RFC 5280 4.2.1.12 on KU consistency with Time Stamping EKU:
	// 	  -- Binding the hash of an object to a time
	//    -- Key usage bits that may be consistent: digitalSignature
	//    -- and/or nonRepudiation
	//  Note: Recent editions of X.509 have renamed nonRepudiation bit to contentCommitment

	allowedCombinations := []x509.KeyUsage{
		x509.KeyUsageDigitalSignature,
		x509.KeyUsageContentCommitment,
		x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment,
	}

	return containsCombination(c.KeyUsage, allowedCombinations)
}

//CheckConsistencyWithEKUOcspSigning checks if KU bits are consistent with Ocsp Signing EKU bit
func CheckConsistencyWithEKUOcspSigning(c *x509.Certificate) bool {
	// 	RFC 5280 4.2.1.12 on KU consistency with Ocsp Signing EKU:
	// 	  -- Signing OCSP responses
	//    -- Key usage bits that may be consistent: digitalSignature
	//    -- and/or nonRepudiation
	//  Note: Recent editions of X.509 have renamed nonRepudiation bit to contentCommitment

	allowedCombinations := []x509.KeyUsage{
		x509.KeyUsageDigitalSignature,
		x509.KeyUsageContentCommitment,
		x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment,
	}

	return containsCombination(c.KeyUsage, allowedCombinations)
}

//  containsCombination takes the KeyUsage of the certificate being linted and a slice of KU bit
//  combinations that are consistent with a certain EKU. The function returns true if the purposes
//  are consistent with both extensions.
func containsCombination(combination x509.KeyUsage, allowedCombinations []x509.KeyUsage) bool {
	for _, c := range allowedCombinations {
		if combination == c {
			return true
		}
	}
	return false
}
