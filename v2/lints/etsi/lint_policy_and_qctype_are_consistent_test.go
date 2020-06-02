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
			Name:           "Standard EV certificate.",
			InputFilename:  "evAllGood.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-eseal QcType. Policies contain the identifiers qcp-web and QCP-w-psd2",
			InputFilename:  "QcStmtPsd2Cert84InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-esign QcType. Policies contain the identifiers qcp-web and QCP-w-psd2",
			InputFilename:  "QcStmtPsd2Cert85InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-web QcType. Policies contain the identifiers qcp-web and QCP-w-psd2",
			InputFilename:  "QcStmtPsd2Cert86ValidWebPolicyConsistent.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-eseal QcType. Policies contain only the identifier qcp-natural.",
			InputFilename:  "QcStmtPsd2Cert87EuQcCertWoQcPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-eseal QcType. Policies contain only the identifier qcp-natural-qscd.",
			InputFilename:  "QcStmtPsd2Cert88EuQcCertWoQcPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-esign QcType. Policies contain only the identifier qcp-natural.",
			InputFilename:  "QcStmtPsd2Cert89ValidQcTypeEsign.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-esign QcType. Policies contain only the identifier qcp-natural-qscd.",
			InputFilename:  "QcStmtPsd2Cert90MissingQcSscdStmt.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-web QcType. Policies contain only the identifier qcp-natural.",
			InputFilename:  "QcStmtPsd2Cert91InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-web QcType. Policies contain only the identifier qcp-natural-qscd.",
			InputFilename:  "QcStmtPsd2Cert92InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-eseal QcType. Policies contain only the identifier qcp-legal",
			InputFilename:  "QcStmtPsd2Cert93ValidEsealPolicyConsistent.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-eseal QcType. Policies contain only the identifier qcp-legal-qscd",
			InputFilename:  "QcStmtPsd2Cert94MissingQcSscdStmt.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-esign QcType. Policies contain only the identifier qcp-legal",
			InputFilename:  "QcStmtPsd2Cert95InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-esign QcType. Policies contain only the identifier qcp-legal-qscd.",
			InputFilename:  "QcStmtPsd2Cert96InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-web QcType. Policies contain only the identifier qcp-legal.",
			InputFilename:  "QcStmtPsd2Cert97InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qct-web QcType. Policies contain only the identifier qcp-legal-qscd.",
			InputFilename:  "QcStmtPsd2Cert98InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance). Policies contain the identifiers qcp-web and QCP-w-psd2",
			InputFilename:  "QcStmtPsd2Cert102InconsistentPolicy.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qcs-QcSSCD set and id-etsi-qct-esign QcType. Policies contain only the identifier qcp-natural.",
			InputFilename:  "QcStmtPsd2Cert103InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qcs-QcSSCD set and id-etsi-qct-esign QcType. Policies contain only the identifier qcp-natural-qscd.",
			InputFilename:  "QcStmtPsd2Cert104ValidEsignQcSSCD.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qcs-QcSSCD set and id-etsi-qct-eseal QcType. Policies contain only the identifier qcp-legal.",
			InputFilename:  "QcStmtPsd2Cert105InconsistentPolicy.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EU qualified certificate (id-etsi-qcs-QcCompliance) with id-etsi-qcs-QcSSCD set and id-etsi-qct-eseal QcType. Policies contain only the identifier qcp-legal-qscd.",
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
