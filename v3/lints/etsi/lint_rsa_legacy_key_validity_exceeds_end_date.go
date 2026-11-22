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

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/************************************************
ETSI TS 119 312, clause 8.4:

RSA keys with a length of at least 1 900 bits and less than 3 000 bits shall not be used to issue new
certificates after 2026-12-31. Certificates based on such keys that were issued on or before 2026-12-31
shall have a validity period ending no later than 2028-12-31.

************************************************/

type rsaLegacyKeyValidityExceedsEndDate struct{}

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:            "e_etsi_rsa_legacy_key_validity_period_exceeds_end_date",
			Description:     "Certificates issued on or before 2026-12-31 with an RSA key of at least 1900 and less than 3000 bits shall have a validity period ending no later than 2028-12-31",
			Citation:        "ETSI TS 119 312 / Clause 8.4",
			Source:          lint.EtsiEsi,
			EffectiveDate:   util.ZeroDate,
			IneffectiveDate: util.EtsiTs119312RsaSunsetDate,
		},
		Lint: NewRsaLegacyKeyValidityExceedsEndDate,
	})
}

func NewRsaLegacyKeyValidityExceedsEndDate() lint.LintInterface {
	return &rsaLegacyKeyValidityExceedsEndDate{}
}

func (l *rsaLegacyKeyValidityExceedsEndDate) CheckApplies(c *x509.Certificate) bool {
	modulusLen, isRsa := etsiRsaModulusLength(c)
	return isRsa &&
		modulusLen >= etsiRsaLegacyModulusLengthLowerBound &&
		modulusLen < etsiRsaMinRecommendedModulusLength &&
		isEtsiCert(c)
}

func (l *rsaLegacyKeyValidityExceedsEndDate) Execute(c *x509.Certificate) *lint.LintResult {
	if c.NotAfter.Before(util.EtsiTs119312RsaLegacyMaxValidityEndDate) {
		return &lint.LintResult{Status: lint.Pass}
	}
	modulusLen, _ := etsiRsaModulusLength(c)
	return &lint.LintResult{
		Status: lint.Error,
		Details: fmt.Sprintf("Certificate uses a %d-bit RSA public key and was issued on or before 2026-12-31, "+
			"so its validity period is required to end no later than 2028-12-31", modulusLen),
	}
}
