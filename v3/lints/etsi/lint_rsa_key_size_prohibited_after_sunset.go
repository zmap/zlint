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

package etsi

import (
	"fmt"

	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/rsa"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/************************************************
ETSI TS 119 312, clause 8.4:

RSA keys with a length of at least 1 900 bits and less than 3 000 bits shall not be used to issue new
certificates after 2026-12-31. ... After 2026-12-31, only RSA keys with a length of at least 3 000 bits
or signature schemes marked as Recommended (R) in the present document shall be used for newly issued
certificates.
************************************************/

const etsiRsaMinRecommendedModulusLength = 3000
const etsiRsaLegacyModulusLengthLowerBound = 1900

type rsaKeySizeProhibitedAfterSunset struct{}

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_etsi_rsa_key_size_prohibited_after_sunset_date",
			Description:   "Certificates issued after 2026-12-31 shall not use RSA keys with a modulus of less than 3000 bits",
			Citation:      "ETSI TS 119 312 / Clause 8.4",
			Source:        lint.EtsiEsi,
			EffectiveDate: util.EtsiTs119312RsaSunsetDate,
		},
		Lint: NewRsaKeySizeProhibitedAfterSunset,
	})
}

func NewRsaKeySizeProhibitedAfterSunset() lint.LintInterface {
	return &rsaKeySizeProhibitedAfterSunset{}
}

// etsiRsaModulusLength returns the RSA modulus bit length of the certificate's
// public key, or false if the key is not an RSA key.
func etsiRsaModulusLength(c *x509.Certificate) (int, bool) {
	if c.PublicKeyAlgorithm != x509.RSA {
		return 0, false
	}
	key, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok || key.N == nil {
		return 0, false
	}
	return key.N.BitLen(), true
}

// etsiCertificatePolicyOIDs are the certificate policies defined in
// ETSI EN 319 411-1 (0.4.0.2042.1.x) and EN 319 411-2 (0.4.0.194112.1.x).
var etsiCertificatePolicyOIDs = []asn1.ObjectIdentifier{
	{0, 4, 0, 2042, 1, 1}, // NCP
	{0, 4, 0, 2042, 1, 2}, // NCP+
	{0, 4, 0, 2042, 1, 3}, // LCP
	{0, 4, 0, 2042, 1, 4}, // EVCP
	{0, 4, 0, 2042, 1, 5}, // EVCP+ (legacy)
	{0, 4, 0, 2042, 1, 6}, // DVCP
	{0, 4, 0, 2042, 1, 7}, // OVCP
	{0, 4, 0, 2042, 1, 8}, // IVCP
	util.QCPnPolicyOID,
	util.QCPlPolicyOID,
	util.QCPnqscdPolicyOID,
	util.QCPlqscdPolicyOID,
	util.QEVCPwPolicyOID,
	util.QNCPwPolicyOID,
	util.QNCPwgenPolicyOID,
}

// isEtsiCert determines whether a certificate is in scope for ETSI TS 119 312.
// A certificate is in scope if it contains the QcCompliance qcStatement, or if
// it asserts one of the certificate policies defined in ETSI EN 319 411-1/-2.
// The mere presence of other qcStatements (e.g. PSD2, QcType) is not sufficient.
func isEtsiCert(c *x509.Certificate) bool {
	if ext := util.GetExtFromCert(c, util.QcStateOid); ext != nil {
		if util.ParseQcStatem(ext.Value, util.IdEtsiQcsQcCompliance).IsPresent() {
			return true
		}
	}
	for _, policy := range c.PolicyIdentifiers {
		for _, etsiPolicy := range etsiCertificatePolicyOIDs {
			if policy.Equal(etsiPolicy) {
				return true
			}
		}
	}
	return false
}

func (l *rsaKeySizeProhibitedAfterSunset) CheckApplies(c *x509.Certificate) bool {
	_, isRsa := etsiRsaModulusLength(c)
	return isRsa && isEtsiCert(c)
}

func (l *rsaKeySizeProhibitedAfterSunset) Execute(c *x509.Certificate) *lint.LintResult {
	modulusLen, _ := etsiRsaModulusLength(c)
	if modulusLen >= etsiRsaMinRecommendedModulusLength {
		return &lint.LintResult{Status: lint.Pass}
	}
	return &lint.LintResult{
		Status: lint.Error,
		Details: fmt.Sprintf("Certificate was issued with a %d-bit RSA public key, which is less than the "+
			"%d-bit minimum required for certificates issued after 2026-12-31",
			modulusLen, etsiRsaMinRecommendedModulusLength),
	}
}
