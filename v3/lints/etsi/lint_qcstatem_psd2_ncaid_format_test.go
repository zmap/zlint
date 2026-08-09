package etsi

/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

func TestEtsiQcStatemPsd2NcaIdFormat(t *testing.T) {
	m := map[string]lint.LintStatus{
		"QcStmtEtsiPsd2ValidCert01.pem":             lint.Pass,
		"QcStmtEtsiPsd2NcaIdBadFormatCert01.pem":    lint.Error,
		"QcStmtEtsiPsd2NcaIdBadCountryCert01.pem":   lint.Error,
		"QcStmtEtsiPsd2NcaIdPspPaValidCert01.pem":   lint.Pass,
		"QcStmtEtsiPsd2NcaIdPspPaInvalidCert01.pem": lint.Error,
		"QcStmtEtsiPsd2NcaIdPspCbValidCert01.pem":   lint.Pass,
		"QcStmtEtsiPsd2WrongEncodingCert01.pem":     lint.Error,
		"QcStmtEtsiValidCert11.pem":                 lint.NA,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_qcstatem_psd2_ncaid_format", inputPath)
		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}

	detailsCases := map[string]string{
		"QcStmtEtsiPsd2NcaIdBadFormatCert01.pem":    "NCAId must be a 2-letter ISO 3166-1 country code, a hyphen, and a 2-8 character uppercase identifier",
		"QcStmtEtsiPsd2NcaIdBadCountryCert01.pem":   "NCAId country code prefix is not an assigned ISO 3166-1 country code",
		"QcStmtEtsiPsd2NcaIdPspPaInvalidCert01.pem": "NCAId must be 'NA' for a PSD2 QcStatement declaring a PSP_CB or PSP_PA role",
	}
	for inputPath, expectedDetails := range detailsCases {
		out := test.TestLint("e_qcstatem_psd2_ncaid_format", inputPath)
		if out.Details != expectedDetails {
			t.Errorf("%s: expected details %q, got %q", inputPath, expectedDetails, out.Details)
		}
	}
}
