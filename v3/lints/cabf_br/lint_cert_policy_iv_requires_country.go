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

type CertPolicyIVRequiresCountry struct{}

/************************************************
BRs: 7.1.6.4
Certificate Policy Identifier: 2.23.140.1.2.3
If the Certificate complies with these Requirements and includes Subject Identity Information
that is verified in accordance with Section 3.2.3.
Such Certificates MUST also include either organizationName or both givenName and
surname, localityName (to the extent such field is required under Section 7.1.4.2.2),
stateOrProvinceName (to the extent required under Section 7.1.4.2.2), and countryName in
the Subject field.
************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_cert_policy_iv_requires_country",
			Description:   "If certificate policy 2.23.140.1.2.3 is included, countryName MUST be included in subject",
			Citation:      "BRs: 7.1.6.4",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABV131Date,
		},
		Lint: NewCertPolicyIVRequiresCountry,
	})
}

func NewCertPolicyIVRequiresCountry() lint.LintInterface {
	return &CertPolicyIVRequiresCountry{}
}

func (l *CertPolicyIVRequiresCountry) CheckApplies(cert *x509.Certificate) bool {
	return util.SliceContainsOID(cert.PolicyIdentifiers, util.BRIndividualValidatedOID)
}

func (l *CertPolicyIVRequiresCountry) Execute(cert *x509.Certificate) *lint.LintResult {
	var out lint.LintResult
	if util.TypeInName(&cert.Subject, util.CountryNameOID) {
		out.Status = lint.Pass
	} else {
		out.Status = lint.Error
	}
	return &out
}
