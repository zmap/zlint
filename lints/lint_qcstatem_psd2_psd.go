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

type qcStatemPsd2Psd struct{}

func (l *qcStatemPsd2Psd) Initialize() error {
	return nil
}

func (l *qcStatemPsd2Psd) CheckApplies(c *x509.Certificate) bool {
	_, isPresent := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	if !isPresent {
		return false
	}
	if util.CertHasSubjectOrgIdWithPrefix(c, "PSD") {
		return true
	}
	return false
}

func (l *qcStatemPsd2Psd) Execute(c *x509.Certificate) *LintResult {
	orgId := util.GetSubjectOrgId(c.RawSubject)
	if orgId.ErrorString != "" {

		return &LintResult{Status: Error, Details: orgId.ErrorString}
	}
	errStr, parsedOi := util.ParseEtsiPsd2OrgId(&orgId.Value)
	if errStr != "" {
		return &LintResult{Status: Error, Details: errStr}

	}
	ext := util.GetExtFromCert(c, util.QcStateOid)
	s := util.ParseQcStatem(ext.Value, util.IdEtsiPsd2Statem)
	if s.GetErrorInfo() != "" {
		return &LintResult{Status: Error, Details: "parsing error for PSD2 QcStatement, cannot properly apply this lint: " + s.GetErrorInfo()}
	}
	psd2Statem, ok := s.(util.EtsiPsd2)
	if !ok {
		return &LintResult{Status: Fatal, Details: "parsed QcStatem is not of type EtsiPsd2"}
	}
	if parsedOi.NcaId != psd2Statem.GetNcaId() {
		return &LintResult{Status: Error, Details: "NcaId in subject:OrganizationIdentifier and PSD2 QcStatement do not match"}
	}
	if parsedOi.Country != psd2Statem.GetNcaCountry() {
		return &LintResult{Status: Error, Details: "Country in subject:OrganizationIdentifier and PSD2 QcStatement do not match"}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_qcstatem_psd2_psd",
		Description:   "Applies to ETSI PSD2 certificates the subject:OrganizationIdentifier of which is of the form 'PSD...' and checks whether the format of the subject:OrganizationIdentifier field is correct and whether the NCAId therein matches the one in the PSD2 statement.",
		Citation:      "ETSI TS 119 495, '5.2.1 PSD2 Authorization Number or other recognized identifier'",
		Source:        EtsiEsi,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2Psd{},
	})
}
