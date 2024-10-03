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

/*
 * Contributed by asantoni64@gmail.com
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
			Name:          "e_invalid_ca_certificate_policies",
			Description:   "Checks compliance with CABF section 7.1.2.10.5",
			Citation:      "CABF BRs §7.1.2.10.5",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABFBRs_2_0_0_Date,
		},
		Lint: NewInvalidCACertificatePolicies,
	})
}

type invalidCACertificatePolicies struct{}

func NewInvalidCACertificatePolicies() lint.LintInterface {
	return &invalidCACertificatePolicies{}
}

func (l *invalidCACertificatePolicies) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c) && !util.IsRootCA(c)
}

func (l *invalidCACertificatePolicies) Execute(c *x509.Certificate) *lint.LintResult {
	if !util.IsExtInCert(c, util.CertPolicyOID) {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "In a non-Root CA certificate, the CertificatePolicies extension is mandatory",
		}
	}

	if len(c.PolicyIdentifiers) == 1 {
		if !c.PolicyIdentifiers[0].Equal(util.AnyPolicyOID) {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: "If a CA cert contains only one policy OID, then it MUST be AnyPolicy",
			}
		}
		return &lint.LintResult{Status: lint.Pass}
	}

	// There are two or more policy OIDs, so check for the presence of
	// at least one reserved policy OID and ensure AnyPolicy is absent
	reservedOIDFound := false
	for _, oid := range c.PolicyIdentifiers {
		if oid.Equal(util.AnyPolicyOID) {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: "The AnyPolicy OID must not be accompanied by any other policy OIDs",
			}
		}
		if oid.Equal(util.BROrganizationValidatedOID) ||
			oid.Equal(util.BRExtendedValidatedOID) ||
			oid.Equal(util.BRDomainValidatedOID) ||
			oid.Equal(util.BRIndividualValidatedOID) {
			reservedOIDFound = true
		}
	}
	if !reservedOIDFound {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "At least one CABF reserved policy OIDs MUST be present in a policy-restricted CA cert",
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
