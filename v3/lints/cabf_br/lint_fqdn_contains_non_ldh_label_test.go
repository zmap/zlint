package cabf_br

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

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestFQDNContainsNonLDHLabel(t *testing.T) {
	testCases := []struct {
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			InputFilename:  `dnsNameValidWildcard.pem`,
			ExpectedResult: lint.Pass,
		},
		{
			InputFilename:  `dnsNameLabel63CharactersValid.pem`,
			ExpectedResult: lint.Pass,
		},
		{
			InputFilename:  `dnsNameNonLDHEmptyLabel.pem`,
			ExpectedResult: lint.Error,
		},
		{
			InputFilename:  `dnsNameNonLDHInvalidCharacter.pem`,
			ExpectedResult: lint.Error,
		},
		{
			InputFilename:  `dnsNameNonLDHTooLongLabel.pem`,
			ExpectedResult: lint.Error,
		},
		{
			InputFilename:  `dnsNameNonLDHStartsWithHyphen.pem`,
			ExpectedResult: lint.Error,
		},
		{
			InputFilename:  `dnsNameNonLDHEndsWithHyphen.pem`,
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.InputFilename, func(t *testing.T) {
			result := test.TestLint("e_fqdn_contains_non_ldh_label", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
