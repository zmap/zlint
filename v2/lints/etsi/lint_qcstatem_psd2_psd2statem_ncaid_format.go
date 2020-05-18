package etsi

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemPsd2Psd2StatemNcaidFormat struct{}

func (l *qcStatemPsd2Psd2StatemNcaidFormat) Initialize() error {
	return nil
}

func (l *qcStatemPsd2Psd2StatemNcaidFormat) CheckApplies(c *x509.Certificate) bool {
	_, isPresent := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	return isPresent
}

func (l *qcStatemPsd2Psd2StatemNcaidFormat) Execute(c *x509.Certificate) *lint.LintResult {
	ext := util.GetExtFromCert(c, util.QcStateOid)
	s := util.ParseQcStatem(ext.Value, util.IdEtsiPsd2Statem)
	if s.GetErrorInfo() != "" {
		return &lint.LintResult{Status: lint.Error, Details: "parsing error for PSD2 QcStatement, cannot properly apply this lint: " + s.GetErrorInfo()}
	}
	psd2Statem, ok := s.(util.EtsiPsd2)
	if !ok {
		return &lint.LintResult{Status: lint.Fatal, Details: "parsed QcStatem is not of type EtsiPsd2"}
	}
	if psd2Statem.GetNcaCountry() == "" || psd2Statem.GetNcaId() == "" {
		return &lint.LintResult{Status: lint.Error, Details: "NCAId field (country-NcaId) in PSD2 QcStatement has invalid format"}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_psd2_psd2statem_ncaid_format",
		Description:   "Checks that the NCAId field of the PSD2 QcStatement has the correct syntax.",
		Citation:      "ETSI TS 119 495, '5.2.3 Name and identifier of the competent authority'",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemPsd2Psd2StatemNcaidFormat{},
	})
}
