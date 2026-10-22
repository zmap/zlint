package cabf_br

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

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/************************************************************************

https://cabforum.org/uploads/CA-Browser-Forum-BR-v2.0.0.pdf

7.1.2.7.1 Subscriber Certificate Types
There are four types of Subscriber Certificates that may be issued, which vary based on the
amount of Subject Information that is included. Each of these certificate types shares a common
profile, with three exceptions: the subject name fields that may occur, how those fields are
validated, and the contents of the certificatePolicies extension.

7.1.2.7.2 Domain Validated
For a Subscriber Certificate to be Domain Validated, it MUST meet the following profile:....

certificatePolicies MUST be present. MUST assert the Reserved Certificate Policy Identifier of
2.23.140.1.2.1 as a policyIdentifier. See Section 7.1.2.7.9
--
7.1.2.7.3 Individual Validated
For a Subscriber Certificate to be Individual Validated, it MUST meet the following profile:

certificatePolicies MUST be present. MUST assert the Reserved Certificate Policy Identifier of
2.23.140.1.2.3 as a policyIdentifier. See Section 7.1.2.7.9.
--
7.1.2.7.4 Organization Validated
For a Subscriber Certificate to be Organization Validated, it MUST meet the following profile:

certificatePolicies MUST be present. MUST assert the Reserved Certificate Policy Identifier of
2.23.140.1.2.2 as a policyIdentifier. See Section 7.1.2.7.9.
--
7.1.2.7.5 Extended Validation
For a Subscriber Certificate to be Extended Validation, it MUST comply with the Certificate Profile
specified in the then‐current version of the Guidelines for the Issuance and Management of
Extended Validation Certificates. In addition, it MUST meet the following profile:

certificatePolicies MUST be present. MUST assert the Reserved Certificate Policy Identifier of
2.23.140.1.1 as a policyIdentifier. See Section 7.1.2.7.9.

*************************************************************************/

type subCertReservedPolicyMissing struct {
	Skip bool `comment:"Set this to true to skip this lint when the certificate is outside the CA/B Forum TLS Baseline Requirements scope"`
}

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_sub_cert_reserved_policy_missing",
			Description:   "Subscriber (Server) Certificate must contain a reserved certificate policy identifier",
			Citation:      "BRs: 7.1.2.7.9",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABFBRs_2_0_0_Date,
		},
		Lint: NewSubCertReservedPolicyMissing,
	})
}

func NewSubCertReservedPolicyMissing() lint.LintInterface {
	return &subCertReservedPolicyMissing{}
}

func (l *subCertReservedPolicyMissing) Configure() any {
	return l
}

func (l *subCertReservedPolicyMissing) CheckApplies(c *x509.Certificate) bool {
	return !l.Skip && util.IsSubscriberCert(c)
}

func (l *subCertReservedPolicyMissing) Execute(c *x509.Certificate) *lint.LintResult {

	for _, oid := range c.PolicyIdentifiers {
		if oid.Equal(util.BRDomainValidatedOID) ||
			oid.Equal(util.BROrganizationValidatedOID) ||
			oid.Equal(util.BRIndividualValidatedOID) ||
			oid.Equal(util.BRExtendedValidatedOID) {
			return &lint.LintResult{Status: lint.Pass}
		}
	}

	return &lint.LintResult{
		Status:  lint.Error,
		Details: "Missing a CA/B Forum TLS reserved certificate policy identifier",
	}
}
