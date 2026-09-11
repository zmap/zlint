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

func TestEtsiQcStatemPsd2OrgIdFormatError(t *testing.T) {
	cases := map[string]struct {
		status  lint.LintStatus
		details string // empty means "don't check Details"
	}{
		"QcStmtEtsiPsd2OrgIdShallValidCert01.pem":         {status: lint.Pass},
		"QcStmtEtsiPsd2OrgIdShouldValidCert01.pem":        {status: lint.NE},
		"QcStmtEtsiPsd2OrgIdShallMalformedCert01.pem":     {status: lint.Error, details: "subject:organizationIdentifier \"PSDES-bde-3DFD21\" does not match the required PSD<country>-<NCAid>-<PSPid> structure"},
		"QcStmtEtsiPsd2OrgIdShouldMalformedCert01.pem":    {status: lint.NE},
		"QcStmtEtsiPsd2OrgIdShallBadCountryCert01.pem":    {status: lint.Error, details: "subject:organizationIdentifier \"PSDZZ-BDE-3DFD21\" has a country code that is not an assigned ISO 3166-1 country"},
		"QcStmtEtsiValidCert11.pem":                       {status: lint.NA},
		"QcStmtEtsiPsd2OrgIdJustBeforeBoundaryCert01.pem": {status: lint.Pass},
		"QcStmtEtsiPsd2OrgIdOnBoundaryCert01.pem":         {status: lint.NE},
	}
	for inputPath, tc := range cases {
		out := test.TestLint("e_qcstatem_psd2_orgid_format", inputPath)
		if out.Status != tc.status {
			t.Errorf("%s: expected %s, got %s", inputPath, tc.status, out.Status)
		}
		if tc.details != "" && out.Details != tc.details {
			t.Errorf("%s: expected details %q, got %q", inputPath, tc.details, out.Details)
		}
	}
}
