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

type subCertAiaShouldNotBeMissing struct{}

/*************************************************************************************************
BRs: 7.1.2.7.6 Subscriber Certificate Extensions
authorityInformationAccess SHOULD be present.

Ballot SC-104 (BRs v2.3.1, effective 2026-10-04) relaxed the requirement from MUST to SHOULD;
e_sub_cert_aia_missing covers certificates issued before that date.
*************************************************************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "w_sub_cert_aia_missing",
			Description:   "Subscriber Certificate: authorityInformationAccess SHOULD be present.",
			Citation:      "BRs: 7.1.2.7.6",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABFBRs_2_3_1_Date,
		},
		Lint: NewSubCertAiaShouldNotBeMissing,
	})
}

func NewSubCertAiaShouldNotBeMissing() lint.LintInterface {
	return &subCertAiaShouldNotBeMissing{}
}

func (l *subCertAiaShouldNotBeMissing) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c)
}

func (l *subCertAiaShouldNotBeMissing) Execute(c *x509.Certificate) *lint.LintResult {
	if util.IsExtInCert(c, util.AiaOID) {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Warn}
	}
}
