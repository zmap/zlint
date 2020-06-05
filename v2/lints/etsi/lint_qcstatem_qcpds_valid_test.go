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

func TestEtsiQcPds(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "QcStmtEtsiNumberInLangCodeCert21.pem",
			InputFilename:  "QcStmtEtsiNumberInLangCodeCert21.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiMissingEnglishPdsCert04.pem",
			InputFilename:  "QcStmtEtsiMissingEnglishPdsCert04.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiTwoEnglPdsCert12.pem",
			InputFilename:  "QcStmtEtsiTwoEnglPdsCert12.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiWrongEncodingLangCodeCert07.pem",
			InputFilename:  "QcStmtEtsiWrongEncodingLangCodeCert07.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiWrongLangCodeCert05.pem",
			InputFilename:  "QcStmtEtsiWrongLangCodeCert05.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiLangCodeUpperCaseCert23.pem",
			InputFilename:  "QcStmtEtsiLangCodeUpperCaseCert23.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtEtsiWrongEncodingUrlCert08.pem",
			InputFilename:  "QcStmtEtsiWrongEncodingUrlCert08.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiTwoLangCodesCert17.pem",
			InputFilename:  "QcStmtEtsiTwoLangCodesCert17.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiValidCert03.pem",
			InputFilename:  "QcStmtEtsiValidCert03.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtEtsiValidCert11.pem",
			InputFilename:  "QcStmtEtsiValidCert11.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtEtsiValidAddLangCert13.pem",
			InputFilename:  "QcStmtEtsiValidAddLangCert13.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtEtsiEsealValidCert02.pem",
			InputFilename:  "QcStmtEtsiEsealValidCert02.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtEtsiNoQcStatmentsCert22.pem",
			InputFilename:  "QcStmtEtsiNoQcStatmentsCert22.pem",
			ExpectedResult: lint.NA,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_qcstatem_qcpds_valid", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
