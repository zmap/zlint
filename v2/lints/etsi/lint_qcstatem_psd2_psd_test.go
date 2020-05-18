package etsi

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

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestQcStatemPsd2Psd(t *testing.T) {
	m := map[string]lint.LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":                lint.Pass,
		"QcStmtPsd2Cert02Psd2ExtInvNcaId.pem":             lint.Error,
		"QcStmtPsd2Cert04Psd2ExtInvNcaId.pem":             lint.Error,
		"QcStmtPsd2Cert05Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert06Psd2ExtWrongNcaId.pem":           lint.Error,
		"QcStmtPsd2Cert07MissingRoleName.pem":             lint.Error,
		"QcStmtPsd2Cert08NcaNameMissing.pem":              lint.Error,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":           lint.Error,
		"QcStmtPsd2Cert10RoleNameMissing.pem":             lint.Error,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":          lint.Error,
		"QcStmtPsd2Cert12NcaIdInconsistent.pem":           lint.Error,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":      lint.Error,
		"QcStmtPsd2Cert14Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert15NcaIdInconsistent.pem":           lint.Error,
		"QcStmtPsd2Cert16RoleIdAndNameInconsistent.pem":   lint.Pass,
		"QcStmtPsd2Cert17NcaIdInconsistent.pem":           lint.Error,
		"QcStmtPsd2Cert18Psd2ExtNcaIdInvalid.pem":         lint.Error,
		"QcStmtPsd2Cert19ValidTwoRoles.pem":               lint.Pass,
		"QcStmtPsd2Cert22NcaNameWrongStringType.pem":      lint.Error,
		"QcStmtPsd2Cert23Psd2ExtNcaIdWrongStringType.pem": lint.Error,
		"QcStmtPsd2Cert24RoleNameIllegalChars.pem":        lint.Error,
		"QcStmtPsd2Cert25Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert26RoleOidAsUtf8Str.pem":            lint.Error,
		"QcStmtPsd2Cert27RoleNameNull.pem":                lint.Error,
		"QcStmtPsd2Cert28NcaNameIa5Str.pem":               lint.Error,
		"QcStmtPsd2Cert30Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert31Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert32NcaIdInconsistent.pem":           lint.Error,
		"QcStmtPsd2Cert34MalformedOrgId.pem":              lint.Error,
		"QcStmtPsd2Cert35MalformedOrgId.pem":              lint.Error,
		"QcStmtPsd2Cert36MalformedOrgId.pem":              lint.Error,
		"QcStmtPsd2Cert37Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert38NcaIdLowerCase.pem":              lint.Error,
		"QcStmtPsd2Cert39Valid.pem":                       lint.Pass,
		"QcStmtPsd2Cert40Valid.pem":                       lint.NA,
		"QcStmtPsd2Cert41ValidLei.pem":                    lint.NA,
		"QcStmtPsd2Cert42OrgIdInvalid.pem":                lint.NA,
		"QcStmtPsd2Cert43LeiNotXg.pem":                    lint.NA,
		"QcStmtPsd2Cert44ValidNtr.pem":                    lint.NA,
		"QcStmtPsd2Cert47MissingUri.pem":                  lint.NA,
		"QcStmtPsd2Cert48LegalPersonSyntaxViolated.pem":   lint.NA,
		"QcStmtPsd2Cert50CabfNtrSopSyntaxInPsd2Cert.pem":  lint.NA,
		"QcStmtPsd2Cert51ValidNoQcs2QcStmt.pem":           lint.NA, // cert has "PDS" not "PSD" prefix
		"EvAltRegNumCert56JurContryNotMatching.pem":       lint.NA,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_qcstatem_psd2_psd", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
