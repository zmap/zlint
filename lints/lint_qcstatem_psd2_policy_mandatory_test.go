package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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
)

func TestQcStatemPsd2PolicyMandatory(t *testing.T) {
	m := map[string]LintStatus{
		"evAllGood.pem": NA,
		"QcStmtPsd2Cert80EuQcCertWoAnyPolicy.pem":        Error,
		"QcStmtPsd2Cert81EuQcCertWoAnyPolicy.pem":        Error,
		"QcStmtPsd2Cert82EuQcCertWoAnyPolicy.pem":        Error,
		"QcStmtPsd2Cert83EuQcCertWoAnyPolicy.pem":        Error,
		"QcStmtPsd2Cert84InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert85InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert86ValidWebPolicyConsistent.pem":   Pass,
		"QcStmtPsd2Cert87EuQcCertWoQcPolicy.pem":         Error,
		"QcStmtPsd2Cert88EuQcCertWoQcPolicy.pem":         Error,
		"QcStmtPsd2Cert89ValidQcTypeEsign.pem":           Pass,
		"QcStmtPsd2Cert90MissingQcSscdStmt.pem":          Error,
		"QcStmtPsd2Cert91InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert92InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert93ValidEsealPolicyConsistent.pem": Pass,
		"QcStmtPsd2Cert94MissingQcSscdStmt.pem":          Error,
		"QcStmtPsd2Cert95InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert96InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert97InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert98InconsistenPolicy.pem":          Error,
		"QcStmtPsd2Cert102InconsistentPolicy.pem":        Error,
		"QcStmtPsd2Cert103InconsistentPolicy.pem":        Error,
		"QcStmtPsd2Cert104ValidEsignQcSSCD.pem":          Pass,
		"QcStmtPsd2Cert105InconsistentPolicy.pem":        Error,
		"QcStmtPsd2Cert106ValidEsealQcSSCD.pem":          Pass,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_qcstatem_psd2_policy_mandatory"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
