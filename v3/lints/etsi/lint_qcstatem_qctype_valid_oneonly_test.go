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

func TestEtsiQcTypeOneOnly(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "NE - correct data and before 2.5.0 Version of ETSI EN 319 412-5",
			InputFilename:  "QcStmtEtsiValidCert11.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Pass - certificate has only eseal qc type",
			InputFilename:  "qctWithEseal.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "Error - certificate has eseal and web qc types",
			InputFilename:   "qctWithEsealAndWeb.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "more than one QcType present, sequence must have exactly size 1",
		},
		{
			Name:            "Error - certificate has a wrong qcType in QcStatements",
			InputFilename:   "qctWithWrongType.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "encountered invalid ETSI QcType OID: 0.4.0.1862.1.2",
		},
		{
			Name:            "Error - certificate has a wrong qcType in QcStatements",
			InputFilename:   "qctWithWrongType.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "encountered invalid ETSI QcType OID: 0.4.0.1862.1.2",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_qcstatem_qctype_valid_oneonly", tc.InputFilename)

			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %v was %v", tc.ExpectedDetails, result.Details)
			}

			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
