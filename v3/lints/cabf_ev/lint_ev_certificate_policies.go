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

package cabf_ev

import (
	"fmt"
	"net/url"
	"time"

	"github.com/zmap/zcrypto/x509"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_ev_certificate_policies",
			Description:   "EV Certificates issued to Subscribers MUST include a CPS URI policy qualifier",
			Citation:      "CA/Browser Forum EV Guidelines v1.8.1, Sec. 9.7.3",
			Source:        lint.CABFEVGuidelines,
			EffectiveDate: util.ZeroDate, // TODO: Find effective date
		},
		Lint: NewEvTechnicalRequirements,
	})
}

type EvCertificatePolicies struct{}

func NewEvTechnicalRequirements() lint.LintInterface {
	return &EvCertificatePolicies{}
}

func (l *EvCertificatePolicies) CheckApplies(c *x509.Certificate) bool {
	return c.ValidationLevel == x509.EV
}

func (l *EvCertificatePolicies) Execute(c *x509.Certificate) *lint.LintResult {
	for _, uris := range c.CPSuri {
		if uris != nil {
			// Policy i is the CPS URI.
			// c.CPSuri[i] is populated only if the policyQualifierID is the correct { id-qt 1 }
			if len(uris) != 1 {
				return &lint.LintResult{Status: lint.Error}
			}
			uri := uris[0]

			// The CPS URI is a "HTTP URL for the Root CA’s Certification Practice Statement"
			cps, err := url.Parse(uri)
			if err != nil {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Error parsing CPS URI: %v", err)}
			}
			if cps.Scheme != "http" && cps.Scheme != "https" {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("CPS URI scheme not http(s): %s", cps.Scheme)}
			}

			if !util.HasValidTLD(cps.Hostname(), time.Now()) {
				return &lint.LintResult{Status: lint.Error}
			}

			return &lint.LintResult{Status: lint.Pass}
		}
	}

	return &lint.LintResult{Status: lint.Error}
}
