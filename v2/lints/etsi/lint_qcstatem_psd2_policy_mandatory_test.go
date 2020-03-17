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
	"testing"

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestQcStatemPsd2PolicyMandatory(t *testing.T) {
	m := map[string]lint.LintStatus{
		"evAllGood.pem": lint.NA,
		"QcStmtPsd2Cert80EuQcCertWoAnyPolicy.pem":        lint.Error,
		"QcStmtPsd2Cert81EuQcCertWoAnyPolicy.pem":        lint.Error,
		"QcStmtPsd2Cert82EuQcCertWoAnyPolicy.pem":        lint.Error,
		"QcStmtPsd2Cert83EuQcCertWoAnyPolicy.pem":        lint.Error,
		"QcStmtPsd2Cert84InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert85InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert86ValidWebPolicyConsistent.pem":   lint.Pass,
		"QcStmtPsd2Cert87EuQcCertWoQcPolicy.pem":         lint.Error,
		"QcStmtPsd2Cert88EuQcCertWoQcPolicy.pem":         lint.Error,
		"QcStmtPsd2Cert89ValidQcTypeEsign.pem":           lint.Pass,
		"QcStmtPsd2Cert90MissingQcSscdStmt.pem":          lint.Error,
		"QcStmtPsd2Cert91InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert92InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert93ValidEsealPolicyConsistent.pem": lint.Pass,
		"QcStmtPsd2Cert94MissingQcSscdStmt.pem":          lint.Error,
		"QcStmtPsd2Cert95InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert96InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert97InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert98InconsistenPolicy.pem":          lint.Error,
		"QcStmtPsd2Cert102InconsistentPolicy.pem":        lint.Error,
		"QcStmtPsd2Cert103InconsistentPolicy.pem":        lint.Error,
		"QcStmtPsd2Cert104ValidEsignQcSSCD.pem":          lint.Pass,
		"QcStmtPsd2Cert105InconsistentPolicy.pem":        lint.Error,
		"QcStmtPsd2Cert106ValidEsealQcSSCD.pem":          lint.Pass,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_qcstatem_psd2_policy_mandatory", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
