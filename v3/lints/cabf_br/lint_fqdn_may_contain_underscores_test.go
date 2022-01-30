/*
* ZLint Copyright 2021 Regents of the University of Michigan
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

func TestFQDNMayContainUnderscores(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "No underscores",
			InputFilename:  "dNSNameNoUnderscoresButAllowed.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Underscores but allowable",
			InputFilename:  "dNSNameUnderscoreAllowedBefore2019.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "Underscore in left most label",
			InputFilename:  "dNSNameInLeftMostLabel.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Underscore but validity is too long",
			InputFilename:  "dNSNameWithUnderscoreValidityTooLong.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Not a valid label is _ replaced with -",
			InputFilename:  "dNSNameUnderscoreInvalidEvenIfReplaced.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Not effective",
			InputFilename:  "dNSNameWithUnderscore.pem",
			ExpectedResult: lint.NE,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_fqdn_may_contain_underscores", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
