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

type qcStatemPsd2Pd2StatemEnc struct{}

func (l *qcStatemPsd2Pd2StatemEnc) Initialize() error {
	return nil
}

func (l *qcStatemPsd2Pd2StatemEnc) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	_, isPresent := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	return isPresent
}

func (l *qcStatemPsd2Pd2StatemEnc) Execute(c *x509.Certificate) *lint.LintResult {
	qcs := util.ParseQcStatem(util.GetExtFromCert(c, util.QcStateOid).Value, util.IdEtsiPsd2Statem)
	if qcs.GetErrorInfo() != "" {
		return &lint.LintResult{Status: lint.Error, Details: qcs.GetErrorInfo()}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_psd2_psd2statem_encoding",
		Description:   "This test checks that a PSD2 QcStatement has the correct encoding.",
		Citation:      "ETSI TS 119 495, 'Annex A (normative): ASN.1 Declaration'",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemPsd2Pd2StatemEnc{},
	})
}
