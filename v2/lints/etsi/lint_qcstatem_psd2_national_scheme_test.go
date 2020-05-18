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

func TestQcStatemPsd2NationalScheme(t *testing.T) {
	m := map[string]lint.LintStatus{
		"QcStmtPsd2Cert01InvalidRoles.pem":              lint.NA,
		"QcStmtPsd2Cert02Psd2ExtInvNcaId.pem":           lint.NA,
		"QcStmtPsd2Cert05Valid.pem":                     lint.NA,
		"QcStmtPsd2Cert07MissingRoleName.pem":           lint.NA,
		"QcStmtPsd2Cert08NcaNameMissing.pem":            lint.NA,
		"QcStmtPsd2Cert09NcaNameZeroLength.pem":         lint.NA,
		"QcStmtPsd2Cert10RoleNameMissing.pem":           lint.NA,
		"QcStmtPsd2Cert11RoleNameZeroLenght.pem":        lint.NA,
		"QcStmtPsd2Cert13Psd2ExtNcaIdZeroLength.pem":    lint.NA,
		"QcStmtPsd2Cert14Valid.pem":                     lint.NA,
		"QcStmtPsd2Cert16RoleIdAndNameInconsistent.pem": lint.NA,
		"QcStmtPsd2Cert47MissingUri.pem":                lint.Error,
		"QcStmtPsd2Cert48LegalPersonSyntaxViolated.pem": lint.NA,
		"QcStmtPsd2Cert49ValidNationalScheme.pem":       lint.Pass,
		"EvAltRegNumCert56JurContryNotMatching.pem":     lint.NA,
		"EvAltRegNumCert52NoOrgId.pem":                  lint.NA,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_qcstatem_psd2_national_scheme", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
