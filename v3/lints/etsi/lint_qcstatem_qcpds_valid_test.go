package etsi

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestEtsiQcPds(t *testing.T) {
	m := map[string]lint.LintStatus{
		"QcStmtEtsiNumberInLangCodeCert21.pem":      lint.Error,
		"QcStmtEtsiMissingEnglishPdsCert04.pem":     lint.Error,
		"QcStmtEtsiTwoEnglPdsCert12.pem":            lint.Error,
		"QcStmtEtsiWrongEncodingLangCodeCert07.pem": lint.Error,
		"QcStmtEtsiWrongLangCodeCert05.pem":         lint.Error,
		"QcStmtEtsiLangCodeUpperCaseCert23.pem":     lint.Pass,
		"QcStmtEtsiWrongEncodingUrlCert08.pem":      lint.Error,
		"QcStmtEtsiTwoLangCodesCert17.pem":          lint.Error,
		"QcStmtEtsiValidCert03.pem":                 lint.Pass,
		"QcStmtEtsiValidCert11.pem":                 lint.Pass,
		"QcStmtEtsiValidAddLangCert13.pem":          lint.Pass,
		"QcStmtEtsiEsealValidCert02.pem":            lint.Pass,
		"QcStmtEtsiNoQcStatmentsCert22.pem":         lint.NA,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_qcstatem_qcpds_valid", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
