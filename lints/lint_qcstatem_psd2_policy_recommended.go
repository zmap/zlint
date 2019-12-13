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

type qcStatemPsd2PolicyRecomm struct{}

func (l *qcStatemPsd2PolicyRecomm) Initialize() error {
	return nil
}

func (l *qcStatemPsd2PolicyRecomm) CheckApplies(c *x509.Certificate) bool {
	isPresent := util.HasCertAnyEtsiQcStatement(c)
	if !isPresent {
		return false
	}
	return true
}

func (l *qcStatemPsd2PolicyRecomm) Execute(c *x509.Certificate) *LintResult {
	if util.HasCertAnyEtsiQcStatement(c) && !util.HasCertAnyEtsiQcpPolicy(c) {
		return &LintResult{Status: Warn, Details: "EU qualified certificate missing QCP policy identifier"}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_qcstatem_psd2_policy_recommended",
		Description:   "Check the requirement that EU Qualified Certificates include at least one of the QCP qualifiers.",
		Citation:      "FETSI EN 319 412-4",
		Source:        EtsiEn_319_412_4_CertProfileWeb,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2PolicyRecomm{},
	})
}
