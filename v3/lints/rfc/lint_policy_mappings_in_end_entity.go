package rfc

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

type policyMappingsInEndEntity struct{}

/************************************************
RFC 5280: 4.2.1.5

   This extension is used in CA certificates.  It lists one or more
   pairs of OIDs; each pair includes an issuerDomainPolicy and a
   subjectDomainPolicy.
************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_policy_mappings_in_end_entity",
			Description:   "The policyMappings extension is used in CA certificates and MUST NOT appear in end entity certificates",
			Citation:      "RFC 5280: 4.2.1.5",
			Source:        lint.RFC5280,
			EffectiveDate: util.RFC5280Date,
		},
		Lint: NewPolicyMappingsInEndEntity,
	})
}

func NewPolicyMappingsInEndEntity() lint.LintInterface {
	return &policyMappingsInEndEntity{}
}

func (l *policyMappingsInEndEntity) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.PolicyMapOID) && !util.IsCACert(c)
}

func (l *policyMappingsInEndEntity) Execute(c *x509.Certificate) *lint.LintResult {
	return &lint.LintResult{Status: lint.Error}
}
