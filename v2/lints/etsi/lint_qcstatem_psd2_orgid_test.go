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

func TestQcStatemPsd2OrgIdTestCerts(t *testing.T) {
	m := map[string]lint.LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":                lint.Pass,
		"QcStmtPsd2Cert02Psd2ExtInvNcaId.pem":             lint.Pass,
		"QcStmtPsd2Cert05Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert07MissingRoleName.pem":             lint.Pass,
		"QcStmtPsd2Cert08NcaNameMissing.pem":              lint.Pass,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":           lint.Pass,
		"QcStmtPsd2Cert10RoleNameMissing.pem":             lint.Pass,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":          lint.Pass,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":      lint.Pass,
		"QcStmtPsd2Cert14Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert16RoleIdAndNameInconsistent.pem":   lint.Pass,
		"QcStmtPsd2Cert19ValidTwoRoles.pem":               lint.Pass,
		"QcStmtPsd2Cert20ValidTwoTimesSameRole.pem":       lint.Pass,
		"QcStmtPsd2Cert21ValidFourRoles.pem":              lint.Pass,
		"QcStmtPsd2Cert22NcaNameWrongStringType.pem":      lint.Pass,
		"QcStmtPsd2Cert23Psd2ExtNcaIdWrongStringType.pem": lint.Pass,
		"QcStmtPsd2Cert24RoleNameIllegalChars.pem":        lint.Pass,
		"QcStmtPsd2Cert26RoleOidAsUtf8Str.pem":            lint.Pass,
		"QcStmtPsd2Cert27RoleNameNull.pem":                lint.Pass,
		"QcStmtPsd2Cert28NcaNameIa5Str.pem":               lint.Pass,
		"QcStmtPsd2Cert29subjOrgIdMissing.pem":            lint.Error,
		"QcStmtPsd2Cert30Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert31Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert33LegalPersonSyntaxViolated.pem":   lint.Error,
		"QcStmtPsd2Cert41ValidLei.pem":                    lint.Pass,
		"QcStmtPsd2Cert42OrgIdInvalid.pem":                lint.Pass,
		"QcStmtPsd2Cert43LeiNotXg.pem":                    lint.Pass,
		"QcStmtPsd2Cert44ValidNtr.pem":                    lint.Pass,
		"QcStmtPsd2Cert45LegalPersonSyntaxViolated.pem":   lint.Error,
		"QcStmtPsd2Cert46InvalidOrgIdNtr.pem":             lint.Pass,
		"QcStmtPsd2Cert48LegalPersonSyntaxViolated.pem":   lint.Error,
		"EvAltRegNumCert56JurContryNotMatching.pem":       lint.NA,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_qcstatem_psd2_orgid", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
