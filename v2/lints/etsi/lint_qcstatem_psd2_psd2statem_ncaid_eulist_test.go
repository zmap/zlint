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

func TestQcStatemPsd2NcaIdEuList(t *testing.T) {
	m := map[string]lint.LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":                lint.Pass,
		"QcStmtPsd2Cert04Psd2ExtInvNcaId.pem":             lint.Warn,
		"QcStmtPsd2Cert05Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert06Psd2ExtWrongNcaId.pem":           lint.Warn,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":           lint.Warn,
		"QcStmtPsd2Cert10RoleNameMissing.pem":             lint.Warn,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":          lint.Warn,
		"QcStmtPsd2Cert12NcaIdInconsistent.pem":           lint.Warn,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":      lint.Warn,
		"QcStmtPsd2Cert14Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert16RoleIdAndNameInconsistent.pem":   lint.Pass,
		"QcStmtPsd2Cert18Psd2ExtNcaIdInvalid.pem":         lint.Warn,
		"QcStmtPsd2Cert19ValidTwoRoles.pem":               lint.Pass,
		"QcStmtPsd2Cert22NcaNameWrongStringType.pem":      lint.Warn,
		"QcStmtPsd2Cert23Psd2ExtNcaIdWrongStringType.pem": lint.Warn,
		"QcStmtPsd2Cert24RoleNameIllegalChars.pem":        lint.Warn,
		"QcStmtPsd2Cert25Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert26RoleOidAsUtf8Str.pem":            lint.Warn,
		"QcStmtPsd2Cert27RoleNameNull.pem":                lint.Warn,
		"QcStmtPsd2Cert28NcaNameIa5Str.pem":               lint.Warn,
		"QcStmtPsd2Cert30Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert37Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert38NcaIdLowerCase.pem":              lint.Warn,
		"QcStmtPsd2Cert44ValidNtr.pem":                    lint.Pass,
		"QcStmtPsd2Cert46InvalidOrgIdNtr.pem":             lint.Pass,
		"QcStmtPsd2Cert47MissingUri.pem":                  lint.Pass,
		"QcStmtPsd2Cert51ValidNoQcs2QcStmt.pem":           lint.Pass,
		"EvAltRegNumCert56JurContryNotMatching.pem":       lint.NA,
	}
	for inputPath, expected := range m {
		out := test.TestLint("w_qcstatem_psd2_psd2statem_ncaid_eulist", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
