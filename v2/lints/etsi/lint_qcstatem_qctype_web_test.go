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

func TestEtsiQcTypeWeb(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "QcStmtEtsiValidCert11.pem",
			InputFilename:  "QcStmtEtsiValidCert11.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtEtsiEsealValidCert02.pem",
			InputFilename:  "QcStmtEtsiEsealValidCert02.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "QcStmtEtsiNoQcStatmentsCert22.pemm",
			InputFilename:  "QcStmtEtsiNoQcStatmentsCert22.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "QcStmtEtsiDuplicatedQcTypeStatementCert27.pem",
			InputFilename:  "QcStmtEtsiDuplicatedQcTypeStatementCert27.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiQcTypeEmtpyCert28.pem",
			InputFilename:  "QcStmtEtsiQcTypeEmtpyCert28.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_qcstatem_qctype_web", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
