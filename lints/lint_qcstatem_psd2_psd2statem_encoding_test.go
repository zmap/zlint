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

func TestQcStatemPsd2Encoding(t *testing.T) {
	m := map[string]LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":                Pass,
		"QcStmtPsd2Cert03MissingRolesOid.pem":             Error,
		"QcStmtPsd2Cert05Valid.pem":                       Pass,
		"QcStmtPsd2Cert07MissingRoleName.pem":             Error,
		"QcStmtPsd2Cert08NcaNameMissing.pem":              Error,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":           Error,
		"QcStmtPsd2Cert10RoleNameMissing.pem":             Error,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":          Error,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":      Error,
		"QcStmtPsd2Cert14Valid.pem":                       Pass,
		"QcStmtPsd2Cert15NcaIdInconsistent.pem":           Pass,
		"QcStmtPsd2Cert17NcaIdInconsistent.pem":           Pass,
		"QcStmtPsd2Cert22NcaNameWrongStringType.pem":      Error,
		"QcStmtPsd2Cert23Psd2ExtNcaIdWrongStringType.pem": Error,
		"QcStmtPsd2Cert24RoleNameIllegalChars.pem":        Error,
		"QcStmtPsd2Cert26RoleOidAsUtf8Str.pem":            Error,
		"QcStmtPsd2Cert27RoleNameNull.pem":                Error,
		"QcStmtPsd2Cert28NcaNameIa5Str.pem":               Error,
		"QcStmtPsd2Cert30Valid.pem":                       Pass,
		"QcStmtPsd2Cert31Valid.pem":                       Pass,
		"QcStmtPsd2Cert39Valid.pem":                       Pass,
		"QcStmtPsd2Cert40Valid.pem":                       Pass,
		"EvAltRegNumCert56JurContryNotMatching.pem":       NA,
		"EvAltRegNumCert57NtrJurSopMissing.pem":           NA,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_qcstatem_psd2_psd2statem_encoding"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
