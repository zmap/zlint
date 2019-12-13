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

func TestQcStatemPsd2NationalScheme(t *testing.T) {
	m := map[string]LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":              NA,
		"QcStmtPsd2Cert02Psd2ExtInvNcaId.pem":           NA,
		"QcStmtPsd2Cert05Valid.pem":                     NA,
		"QcStmtPsd2Cert07MissingRoleName.pem":           NA,
		"QcStmtPsd2Cert08NcaNameMissing.pem":            NA,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":         NA,
		"QcStmtPsd2Cert10RoleNameMissing.pem":           NA,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":        NA,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":    NA,
		"QcStmtPsd2Cert14Valid.pem":                     NA,
		"QcStmtPsd2Cert16RoleIdAndNameInconsistent.pem": NA,
		"QcStmtPsd2Cert47MissingUri.pem":                Error,
		"QcStmtPsd2Cert48LegalPersonSyntaxViolated.pem": NA,
		"QcStmtPsd2Cert49ValidNationalScheme.pem":       Pass,
		"EvAltRegNumCert56JurContryNotMatching.pem":     NA,
		"EvAltRegNumCert52NoOrgId.pem":                  NA,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_qcstatem_psd2_national_scheme"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
