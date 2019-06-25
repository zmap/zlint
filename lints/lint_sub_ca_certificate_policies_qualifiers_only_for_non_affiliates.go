package lints

/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

/************************************************
BRs: 7.1.2.2a certificatePolicies
The following fields MAY be present if the Subordinate CA is not an Affiliate of the entity that controls the Root CA.
certificatePolicies:policyQualifiers:policyQualifierId (Optional)
• id-qt 1 [RFC 5280].
certificatePolicies:policyQualifiers:qualifier:cPSuri (Optional)
• HTTP URL for the Root CA's Certificate Policies, Certification Practice Statement, Relying Party Agreement, or other
pointer to online policy information provided by the CA.
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCACertPolicyQualifiersOnlyForNonAffiliates struct{}

func (l *subCACertPolicyQualifiersOnlyForNonAffiliates) Initialize() error {
	return nil
}

func (l *subCACertPolicyQualifiersOnlyForNonAffiliates) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *subCACertPolicyQualifiersOnlyForNonAffiliates) Execute(c *x509.Certificate) *LintResult {
	for _, policyQualifierIds := range c.QualifierId {
		if policyQualifierIds != nil {
			return &LintResult{Status: Notice}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_sub_ca_certificate_policies_qualifiers_only_for_non_affiliates",
		Description:   "Policy qualifiers shouldn't be present if the subordinate CA is an Affiliate of the entity that controls the Root CA",
		Citation:      "BRs: 7.1.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCACertPolicyQualifiersOnlyForNonAffiliates{},
	})
}
