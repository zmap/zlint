package rfc

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

func TestDuplicateExtensions(t *testing.T) {
	testCases := []struct {
		name            string
		path            string
		expectedStatus  lint.LintStatus
		expectedDetails string
	}{
		{
			name:            "duplicate SAN extension",
			path:            "extSANDuplicated.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "The following extensions are duplicated: 2.5.29.17",
		},
		{
			name:            "multiple duplicate extensions",
			path:            "multDupeExts.pem",
			expectedStatus:  lint.Error,
			expectedDetails: "The following extensions are duplicated: 2.5.29.14, 2.5.29.35",
		},
		{
			name:           "no duplicate extensions",
			path:           "caBasicConstCrit.pem",
			expectedStatus: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := test.TestLint("e_ext_duplicate_extension", tc.path)
			if actual.Status != tc.expectedStatus {
				t.Errorf("%s: expected status %q got %q",
					tc.path, tc.expectedStatus, actual.Status)
			}
			if actual.Details != tc.expectedDetails {
				t.Errorf("%s: expected detail %q got %q",
					tc.path, tc.expectedDetails, actual.Details)
			}
		})
	}
}
