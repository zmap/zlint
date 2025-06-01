/*
 * ZLint Copyright 2025 Regents of the University of Michigan
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

package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSC081ThirdDate47ServerCertValidityTooLong(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "NE - certificate is issued at 20260314 235959, just before the first date",
			InputFilename:  "justBeforeFirstMilestone.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20260315 000000, exactly on the first date and has the the full 200-day validity", //nolint:dupword
			InputFilename:  "exactlyOnFirstMilestoneExactly200days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20260315 000000, exactly on the first date and has a full 199-day validity",
			InputFilename:  "exactlyOnFirstMilestoneExactly199days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20260315 000000, exactly on the first date and has a full 200-day validity plus one second",
			InputFilename:  "exactlyOnFirstMilestoneLongerThan200days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20270314 235959, just before the second date and has a full 200-day validity",
			InputFilename:  "justBeforeSecondMilestoneExactly200days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20270315 000000, exactly on the the second date and has a full 100-day validity", //nolint:dupword
			InputFilename:  "exactlyOnSecondMilestoneExactly100days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20270315 000000, exactly on the second date and has a full 99-day validity",
			InputFilename:  "exactlyOnSecondMilestoneExactly99days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20270315 000000, exactly on the second date and has a full 100-day validity plus one second",
			InputFilename:  "exactlyOnSecondMilestoneLongerThan100days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NE - certificate is issued at 20290314 235959, just before the third date and has a full 100-day validity",
			InputFilename:  "justBeforeThirdMilestoneExactly100days.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Pass - certificate is issued at 20290315 000000, exactly on the third date and has a full 47-day validity",
			InputFilename:  "exactlyOnThirdMilestoneExactly47days.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Pass - certificate is issued at 20290315 000000, exactly on the third date and has a full 46-day validity",
			InputFilename:  "exactlyOnThirdMilestoneExactly46days.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "Error - certificate is issued at 20290315 000000, exactly on the third date and has a full 47-day validity plus one second",
			InputFilename:   "exactlyOnThirdMilestoneLongerThan47days.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "Certificate is issued on or after March 15, 2029 and has a validity of 48 days",
		},
		{
			Name:           "Pass - certificate is issued at 20320201 000000, considering the leap year and has a full 47-day validity",
			InputFilename:  "withinThirdMilestoneLeapYear.pem",
			ExpectedResult: lint.Pass,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_server_cert_valid_time_longer_than_47_days", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if tc.ExpectedResult == lint.Error && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %q, was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
