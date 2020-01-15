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

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type qcStatemPsd2VatNtr struct{}

func (l *qcStatemPsd2VatNtr) Initialize() error {
	return nil
}

func (l *qcStatemPsd2VatNtr) CheckApplies(c *x509.Certificate) bool {
	_, isPresent := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	if !isPresent {
		return false
	}
	if util.CertHasSubjectOrgIdWithPrefix(c, "VAT") || util.CertHasSubjectOrgIdWithPrefix(c, "NTR") || util.CertHasSubjectOrgIdWithPrefix(c, "LEI") {
		return true
	}

	return false
}

func (l *qcStatemPsd2VatNtr) Execute(c *x509.Certificate) *LintResult {
	orgId := util.GetSubjectOrgId(c.RawSubject)
	if orgId.ErrorString != "" {

		return &LintResult{Status: Error, Details: orgId.ErrorString}
	}
	errStr, parsedOi := util.ParseEtsiPsd2OrgId(&orgId.Value)
	if errStr != "" {
		return &LintResult{Status: Error, Details: errStr}

	}
	if util.CertHasSubjectOrgIdWithPrefix(c, "LEI") {
		if parsedOi.Country != "XG" {
			return &LintResult{Status: Error, Details: "country not equal to 'XG' as required in case of 'LEI...' "}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_qcstatem_psd2_vat_ntr",
		Description:   "Applies to ETSI PSD2 certificates the subject:OrganizationIdentifier of which is of the form 'VAT...' or 'NTR...' or 'LEI... 'and checks whether the format of the subject:OrganizationIdentifier field is correct",
		Citation:      "ETSI EN 319 412-1, Sec. '5.1.4 Legal person semantics identifier' ",
		Source:        EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemPsd2VatNtr{},
	})
}
