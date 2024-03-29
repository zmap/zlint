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

type NCReservedIPNet struct{}

/************************************************
BRs: 7.1.5
(b) For each iPAddress range in permittedSubtrees, the CA MUST confirm that the
Applicant has been assigned the iPAddress range or has been authorized by the
assigner to act on the assignee's behalf.

BRs: 7.1.4.2.1
CAs SHALL NOT issue certificates with a subjectAlternativeName extension or
Subject commonName field containing a Reserved IP Address or Internal Name.
************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_ext_nc_intersects_reserved_ip",
			Description:   "iPAddress name constraint intersects an IANA reserved network",
			Citation:      "BRs: 7.1.5 / 7.1.4.2.1",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABEffectiveDate,
		},
		Lint: NewNCReservedIPNet,
	})
}

func NewNCReservedIPNet() lint.LintInterface {
	return &NCReservedIPNet{}
}

func (l *NCReservedIPNet) CheckApplies(c *x509.Certificate) bool {
	return c.NotAfter.After(util.NoReservedIP) && util.IsExtInCert(c, util.NameConstOID)
}

func (l *NCReservedIPNet) Execute(c *x509.Certificate) *lint.LintResult {
	for _, constraint := range c.PermittedIPAddresses {
		if util.IntersectsIANAReserved(constraint.Data) {
			return &lint.LintResult{Status: lint.Error}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
