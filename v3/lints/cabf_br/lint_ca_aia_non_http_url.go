/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

	"strings"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_ca_aia_non_http_url",
			Description:   "Within the AIA extension of CA certificates, accessLocations must contain HTTP URLs",
			Citation:      "CABF BRs section 7.1.2.10.3 (CA Certificate Authority Information Access)",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.SC62EffectiveDate,
		},
		Lint: NewCAAIANonHTTPURL,
	})
}

type CAAIANonHTTPURL struct{}

func NewCAAIANonHTTPURL() lint.LintInterface {
	return &CAAIANonHTTPURL{}
}

func (l *CAAIANonHTTPURL) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) &&
		(len(c.IssuingCertificateURL) > 0 || len(c.OCSPServer) > 0)
}

func (l *CAAIANonHTTPURL) Execute(c *x509.Certificate) *lint.LintResult {
	for _, url := range c.IssuingCertificateURL {
		if !strings.HasPrefix(strings.ToLower(url), "http://") {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: "For the 'caIssuers' accessMethod within the AIA extension, accessLocation must contain an HTTP URL",
			}
		}
	}

	for _, url := range c.OCSPServer {
		if !strings.HasPrefix(strings.ToLower(url), "http://") {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: "For the 'ocsp' accessMethod within the AIA extension, accessLocation must contain an HTTP URL",
			}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
