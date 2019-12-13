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

func TestQcStatemPsd2NcaIdEuList(t *testing.T) {
	m := map[string]LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":                Pass,
		"QcStmtPsd2Cert04Psd2ExtInvNcaId.pem":             Warn,
		"QcStmtPsd2Cert05Valid.pem":                       Pass,
		"QcStmtPsd2Cert06Psd2ExtWrongNcaId.pem":           Warn,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":           Warn,
		"QcStmtPsd2Cert10RoleNameMissing.pem":             Warn,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":          Warn,
		"QcStmtPsd2Cert12NcaIdInconsistent.pem":           Warn,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":      Warn,
		"QcStmtPsd2Cert14Valid.pem":                       Pass,
		"QcStmtPsd2Cert16RoleIdAndNameInconsistent.pem":   Pass,
		"QcStmtPsd2Cert18Psd2ExtNcaIdInvalid.pem":         Warn,
		"QcStmtPsd2Cert19ValidTwoRoles.pem":               Pass,
		"QcStmtPsd2Cert22NcaNameWrongStringType.pem":      Warn,
		"QcStmtPsd2Cert23Psd2ExtNcaIdWrongStringType.pem": Warn,
		"QcStmtPsd2Cert24RoleNameIllegalChars.pem":        Warn,
		"QcStmtPsd2Cert25Valid.pem":                       Pass,
		"QcStmtPsd2Cert26RoleOidAsUtf8Str.pem":            Warn,
		"QcStmtPsd2Cert27RoleNameNull.pem":                Warn,
		"QcStmtPsd2Cert28NcaNameIa5Str.pem":               Warn,
		"QcStmtPsd2Cert30Valid.pem":                       Pass,
		"QcStmtPsd2Cert37Valid.pem":                       Pass,
		"QcStmtPsd2Cert38NcaIdLowerCase.pem":              Warn,
		"QcStmtPsd2Cert44ValidNtr.pem":                    Pass,
		"QcStmtPsd2Cert46InvalidOrgIdNtr.pem":             Pass,
		"QcStmtPsd2Cert47MissingUri.pem":                  Pass,
		"QcStmtPsd2Cert51ValidNoQcs2QcStmt.pem":           Pass,
		"EvAltRegNumCert56JurContryNotMatching.pem":       NA,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["w_qcstatem_psd2_psd2statem_ncaid_eulist"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
