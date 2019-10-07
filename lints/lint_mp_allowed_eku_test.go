package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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
	"fmt"
	"testing"
)

func TestAllowedEKUs(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult LintStatus
	}{
		{
			Name:           "SubCA with no EKU",
			InputFilename:  "mpSubCAEKUDisallowed1.pem",
			ExpectedResult: Error,
		},
		{
			Name:           "SubCA with anyExtendedKeyUsage",
			InputFilename:  "mpSubCAEKUDisallowed2.pem",
			ExpectedResult: Error,
		},
		{
			Name:           "SubCA with serverAuth and emailProtection",
			InputFilename:  "mpSubCAEKUDisallowed3.pem",
			ExpectedResult: Error,
		},
		{
			Name:           "SubCA with serverAuth EKU",
			InputFilename:  "mpSubCAEKUAllowed.pem",
			ExpectedResult: Pass,
		},
		{
			Name:           "Cross-Certificate with no EKU",
			InputFilename:  "mpCrossCertNoEKU.pem",
			ExpectedResult: NA,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.InputFilename)
			result := Lints["e_mp_allowed_eku"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
