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

func TestPolicyAndQCTypeConsistent(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "evAllGood.pem",
			InputFilename:  "evAllGood.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "QcStmtPsd2Cert80EuQcCertWoAnyPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert80EuQcCertWoAnyPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert81EuQcCertWoAnyPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert81EuQcCertWoAnyPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert82EuQcCertWoAnyPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert82EuQcCertWoAnyPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert83EuQcCertWoAnyPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert83EuQcCertWoAnyPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert84InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert84InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert85InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert85InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert86ValidWebPolicyConsistent.pem",
			InputFilename:  "QcStmtPsd2Cert86ValidWebPolicyConsistent.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtPsd2Cert87EuQcCertWoQcPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert87EuQcCertWoQcPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert88EuQcCertWoQcPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert88EuQcCertWoQcPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert89ValidQcTypeEsign.pem",
			InputFilename:  "QcStmtPsd2Cert89ValidQcTypeEsign.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtPsd2Cert90MissingQcSscdStmt.pem",
			InputFilename:  "QcStmtPsd2Cert90MissingQcSscdStmt.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert91InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert91InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert92InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert92InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert93ValidEsealPolicyConsistent.pem",
			InputFilename:  "QcStmtPsd2Cert93ValidEsealPolicyConsistent.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtPsd2Cert94MissingQcSscdStmt.pem",
			InputFilename:  "QcStmtPsd2Cert94MissingQcSscdStmt.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert95InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert95InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert96InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert96InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert97InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert97InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert98InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert98InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with QCP-w-psd2 and qcp-web policies",
			InputFilename:  "QcStmtPsd2Cert102InconsistentPolicy.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with esi4-qcStatement-4 but without policy qcp-natural-qscd or qcp-legal-qscd",
			InputFilename:  "QcStmtPsd2Cert103InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert104ValidEsignQcSSCD.pem",
			InputFilename:  "QcStmtPsd2Cert104ValidEsignQcSSCD.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtPsd2Cert105InconsistentPolicy.pem",
			InputFilename:  "QcStmtPsd2Cert105InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtPsd2Cert106ValidEsealQcSSCD.pem",
			InputFilename:  "QcStmtPsd2Cert106ValidEsealQcSSCD.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_policy_and_qctype_consistent", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
