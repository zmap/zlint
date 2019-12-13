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

func (l *qcStatemPsd2Pd2StatemEnc) Execute(c *x509.Certificate) *LintResult {
	qcs := util.ParseQcStatem(util.GetExtFromCert(c, util.QcStateOid).Value, util.IdEtsiPsd2Statem)
	if qcs.GetErrorInfo() != "" {
		return &LintResult{Status: Error, Details: qcs.GetErrorInfo()}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_qcstatem_psd2_psd2statem_encoding",
		Description:   "This test checks that a PSD2 QcStatement has the correct encoding.",
		Citation:      "ETSI TS 119 495, 'Annex A (normative): ASN.1 Declaration'",
		Source:        EtsiTs_119_495_EsiPsd,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemPsd2Pd2StatemEnc{},
	})
}
