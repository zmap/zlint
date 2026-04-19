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

func TestEtsiQcType(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "NE - correct data and before 2.5.0 Version of ETSI EN 319 412-5",
			InputFilename:  "QcStmtEtsiValidCert03.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Pass - QcStmtEtsiValidCert11",
			InputFilename:  "QcStmtEtsiValidCert11.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Error - QcStmtEtsiValidAddLangCert13",
			InputFilename:  "QcStmtEtsiValidAddLangCert13.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Pass - QcStmtEtsiEsealValidCert02",
			InputFilename:  "QcStmtEtsiEsealValidCert02.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - Certificate has no QcStatements",
			InputFilename:  "QcStmtEtsiNoQcStatmentsCert22.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NE - certificate has only eseal qc type and issued on  01. May 2025",
			InputFilename:  "qctWithEseal.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Error - certificate has a wrong qcType in QcStatements and is issued in 2024",
			InputFilename:  "qctWithWrongType_2024.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_qcstatem_qctype_valid", tc.InputFilename)

			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %v was %v", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
