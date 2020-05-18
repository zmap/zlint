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

type qcStatemPsd2PolicyRecomm struct{}

func (l *qcStatemPsd2PolicyRecomm) Initialize() error {
	return nil
}

func (l *qcStatemPsd2PolicyRecomm) CheckApplies(c *x509.Certificate) bool {
	isPresent := util.HasCertAnyEtsiQcStatement(c)
	return isPresent
}

func (l *qcStatemPsd2PolicyRecomm) Execute(c *x509.Certificate) *lint.LintResult {
	if util.HasCertAnyEtsiQcStatement(c) && !util.HasCertAnyEtsiQcpPolicy(c) {
		return &lint.LintResult{Status: lint.Warn, Details: "EU qualified certificate missing QCP policy identifier"}
	}

	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_qcstatem_psd2_policy_recommended",
		Description:   "Check the requirement that EU Qualified Certificates include at least one of the QCP qualifiers.",
		Citation:      "ETSI EN 319 412-4",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2PolicyRecomm{},
	})
}
