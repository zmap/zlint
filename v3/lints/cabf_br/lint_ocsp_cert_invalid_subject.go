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

package cabf_br

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:                    "e_ocsp_cert_invalid_subject",
			Description:             "Checks that OCSP Responder certs contain the required Subject attribs",
			Citation:                "TLS BRs §7.1.2.8 and §7.1.2.10.2",
			Source:                  lint.CABFBaselineRequirements,
			EffectiveDate:           util.SC62EffectiveDate,
			OverrideFrameworkFilter: true,
		},
		Lint: NewOcspCertInvalidSubject,
	})
}

type OcspCertInvalidSubject struct{}

func NewOcspCertInvalidSubject() lint.LintInterface {
	return &OcspCertInvalidSubject{}
}

func (l *OcspCertInvalidSubject) CheckApplies(c *x509.Certificate) bool {
	return util.HasEKU(c, x509.ExtKeyUsageOcspSigning)
}

func (l *OcspCertInvalidSubject) Execute(c *x509.Certificate) *lint.LintResult {

	var errDetails string

	if c.Subject.CommonName == "" {
		errDetails = "missing commonName"
	}
	if len(c.Subject.Organization) == 0 {
		errDetails = "missing organizationName"
	}
	if len(c.Subject.Country) == 0 {
		errDetails = "missing countryName"
	}

	if errDetails != "" {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "Invalid Subject for an OCSP Responder Certificate: " + errDetails,
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
