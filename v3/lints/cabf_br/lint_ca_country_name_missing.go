package cabf_br

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

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type caCountryNameMissing struct{}

/************************************************
--- Citation History of this Requirement ---
v1.0 to v1.2.4:   9.1.4
v1.2.5:           Appendix B 1E, 2H
v1.3.0 to v1.4.7: 7.1.2.1e, 7.1.2.2h
v1.4.8 to v1.8.7: 7.1.4.3.1c
v2.0.0 to v2.1.6: 7.1.2.10.2

--- Version Notes ---
As of v2.0.0, this requirement no longer applies to CA certificates that conform to the
Cross-Certified Subordinate CA Certificate Profile. This lint does not implement this exemption
because it is infeasible to identify certificates to which it applies from only the certificate.

This requirement was baselined at v2.1.6 and is current.

--- Requirements Language ---
BRs: 7.1.2
If the CA asserts compliance with these Baseline Requirements, all certificates that it issues MUST
comply with one of the following certificate profiles

[Each of the CA profiles, excepting the Cross-Certified Subordinate CA Certificate Profile,
specifies the subject follows 7.1.2.10.2 in the profile table.]

BRs: 7.1.2.10.2
The following table details the acceptable AttributeTypes that may appear within the type
field of an AttributeTypeAndValue, as well as the contents permitted within the value field.
+----------------+----------+--------------------------------------------------------+-----------------+
| Attribute Name | Presence | Value                                                  | Verification    |
+----------------+----------+--------------------------------------------------------+-----------------+
| countryName    | MUST     | The two‐letter ISO 3166‐1 country code for the country | Section 3.2.2.3 |
|                |          | in which the CA’s place of business is located.        |                 |
+----------------+----------+--------------------------------------------------------+-----------------+
************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_ca_country_name_missing",
			Description:   "Root and Subordinate CA certificates MUST have a countryName present in subject information",
			Citation:      "BRs: 7.1.2.10.2",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABEffectiveDate,
		},
		Lint: NewCaCountryNameMissing,
	})
}

func NewCaCountryNameMissing() lint.LintInterface {
	return &caCountryNameMissing{}
}

func (l *caCountryNameMissing) CheckApplies(c *x509.Certificate) bool {
	return c.IsCA
}

func (l *caCountryNameMissing) Execute(c *x509.Certificate) *lint.LintResult {
	if c.Subject.Country != nil && c.Subject.Country[0] != "" {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Error}
	}
}
