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

func TestQcStatemPsd2OrgIdTestCerts(t *testing.T) {
	m := map[string]LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":                Pass,
		"QcStmtPsd2Cert02Psd2ExtInvNcaId.pem":             Pass,
		"QcStmtPsd2Cert05Valid.pem":                       Pass,
		"QcStmtPsd2Cert07MissingRoleName.pem":             Pass,
		"QcStmtPsd2Cert08NcaNameMissing.pem":              Pass,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":           Pass,
		"QcStmtPsd2Cert10RoleNameMissing.pem":             Pass,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":          Pass,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":      Pass,
		"QcStmtPsd2Cert14Valid.pem":                       Pass,
		"QcStmtPsd2Cert16RoleIdAndNameInconsistent.pem":   Pass,
		"QcStmtPsd2Cert19ValidTwoRoles.pem":               Pass,
		"QcStmtPsd2Cert20ValidTwoTimesSameRole.pem":       Pass,
		"QcStmtPsd2Cert21ValidFourRoles.pem":              Pass,
		"QcStmtPsd2Cert22NcaNameWrongStringType.pem":      Pass,
		"QcStmtPsd2Cert23Psd2ExtNcaIdWrongStringType.pem": Pass,
		"QcStmtPsd2Cert24RoleNameIllegalChars.pem":        Pass,
		"QcStmtPsd2Cert26RoleOidAsUtf8Str.pem":            Pass,
		"QcStmtPsd2Cert27RoleNameNull.pem":                Pass,
		"QcStmtPsd2Cert28NcaNameIa5Str.pem":               Pass,
		"QcStmtPsd2Cert29subjOrgIdMissing.pem":            Error,
		"QcStmtPsd2Cert30Valid.pem":                       Pass,
		"QcStmtPsd2Cert31Valid.pem":                       Pass,
		"QcStmtPsd2Cert33LegalPersonSyntaxViolated.pem":   Error,
		"QcStmtPsd2Cert41ValidLei.pem":                    Pass,
		"QcStmtPsd2Cert42OrgIdInvalid.pem":                Pass,
		"QcStmtPsd2Cert43LeiNotXg.pem":                    Pass,
		"QcStmtPsd2Cert44ValidNtr.pem":                    Pass,
		"QcStmtPsd2Cert45LegalPersonSyntaxViolated.pem":   Error,
		"QcStmtPsd2Cert46InvalidOrgIdNtr.pem":             Pass,
		"QcStmtPsd2Cert48LegalPersonSyntaxViolated.pem":   Error,
		"EvAltRegNumCert56JurContryNotMatching.pem":       NA,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_qcstatem_psd2_orgid"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
