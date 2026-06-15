package cabf_cs_br

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
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type csEVSubjectOrgMaxLength struct{}

/************************************************
CSBRs: 7.1.4.2.4
Certificate Field: subject:organizationName (OID 2.5.4.10)

If the combination of names or the organization name by itself exceeds 64 characters, the CA
MAY abbreviate parts of the organization name, and/or omit non-material words in the
organization name in such a way that the text in this field does not exceed the 64-character
limit; provided that the CA checks this field in accordance with the High Risk Certificate
Request requirements of Section 4.2.1 and a Relying Party will not be misled into thinking
that they are dealing with a different organization. In cases where this is not possible, the
CA MUST NOT issue the EV Code Signing Certificate.
************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_cs_ev_subject_org_max_length",
			Description:   "EV Code Signing certificate organizationName MUST NOT exceed 64 characters",
			Citation:      "CSBRs: 7.1.4.2.4",
			Source:        lint.CABFCSBaselineRequirements,
			EffectiveDate: util.CABF_CS_BRs_1_2_Date,
		},
		Lint: NewCsEVSubjectOrgMaxLength,
	})
}

func NewCsEVSubjectOrgMaxLength() lint.LintInterface {
	return &csEVSubjectOrgMaxLength{}
}

func (l *csEVSubjectOrgMaxLength) CheckApplies(cert *x509.Certificate) bool {
	return util.IsSubscriberCert(cert) && util.IsEVCodeSigning(cert.PolicyIdentifiers) && len(cert.Subject.Organization) > 0
}

func (l *csEVSubjectOrgMaxLength) Execute(cert *x509.Certificate) *lint.LintResult {
	for _, org := range cert.Subject.Organization {
		if utf8.RuneCountInString(org) > 64 {
			return &lint.LintResult{Status: lint.Error, Details: "EV Code Signing certificate organizationName exceeds 64 characters."}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
