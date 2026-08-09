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

func TestEtsiQcStatemPsd2NcaIdEulist(t *testing.T) {
	cases := map[string]struct {
		status  lint.LintStatus
		details string // empty means "don't check Details"
	}{
		"QcStmtEtsiPsd2NcaIdEuMatchCert01.pem": {status: lint.Pass},
		"QcStmtEtsiPsd2NcaIdNonEuCert01.pem":   {status: lint.Pass},
		// QcStmtEtsiPsd2ValidCert01.pem is structurally valid (hence the name) but
		// carries NCAId "ES-BDE" — ETSI's own now-stale §5.2.1 example value. Spain's
		// current EBA code is "ES-BE", so a Warn here is the intended result.
		"QcStmtEtsiPsd2ValidCert01.pem":         {status: lint.Warn, details: "NCAId \"ES-BDE\" does not match the current EU NCA identifier \"ES-BE\" for country ES"},
		"QcStmtEtsiPsd2WrongEncodingCert01.pem": {status: lint.Error, details: "error with ASN.1 encoding, possibly a wrong ASN.1 string type was used"},
		"QcStmtEtsiValidCert11.pem":             {status: lint.NA},
	}
	for inputPath, tc := range cases {
		out := test.TestLint("w_qcstatem_psd2_ncaid_eulist", inputPath)
		if out.Status != tc.status {
			t.Errorf("%s: expected %s, got %s", inputPath, tc.status, out.Status)
		}
		if tc.details != "" && out.Details != tc.details {
			t.Errorf("%s: expected details %q, got %q", inputPath, tc.details, out.Details)
		}
	}
}
