package rfc

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

func TestUnrecommendedQualifier(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Certificate with certificate policies extension and without the anyPolicy policyIdentifier present",
			InputFilename:  "withoutAnyPolicy.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Certificate without certificate policies extension",
			InputFilename:  "CNWithoutSANSeptember2021.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Certificate with certificate policies extension, with anyPolicy policyIdentifier present, without policyQualifiers",
			InputFilename:  "withAnyPolicyAndNoPolicyQualifiers.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Certificate with certificate policies extension, with anyPolicy policyIdentifier present and a CPS qualifier present",
			InputFilename:  "withAnyPolicyAndCPSQualifier.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Certificate with certificate policies extension, with anyPolicy policyIdentifier present and a UserNotice qualifier present",
			InputFilename:  "withAnyPolicyAndUserNoticeQualifier.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Certificate with certificate policies extension, with anyPolicy policyIdentifier present and neither CPS nor UserNotice qualifier present",
			InputFilename:  "withAnyPolicyWithoutCPSOrUserNoticeQualifier.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Certificate with certificate policies extension and many combinations of policies and qualifiers",
			InputFilename:  "withValidPoliciesRegardingAnyPolicy.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ext_cert_policy_disallowed_any_policy_qualifier", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
