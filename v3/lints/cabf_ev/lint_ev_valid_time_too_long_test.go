package cabf_ev

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

func TestEvValidTooLong(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "EV certificate valid for > 27 months",
			InputFilename:  "evValidTooLong.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EV certificate issued before Ballot 193 valid for 27 months",
			InputFilename:  "evValidNotTooLong.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EV certificate issued after Ballot 193, valid for 825 days, which is >27 months",
			InputFilename:  "evValidNotTooLong825Days.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ev_valid_time_too_long", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
