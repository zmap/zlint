package community

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

func TestPubSuffix(t *testing.T) {
	testCases := []struct {
		path            string
		expectedStatus  lint.LintStatus
		expectedDetails string
	}{
		{
			path:            "SANBareSuffix.pem",
			expectedStatus:  lint.Notice,
			expectedDetails: "1 DNS name(s) are bare public suffixes: co.uk",
		},
		{
			path:            "multiEmptyPubSuffix.pem",
			expectedStatus:  lint.Notice,
			expectedDetails: "2 DNS name(s) are bare public suffixes: co.uk, ca",
		},
		{
			path:           "newlinesInTLD.pem",
			expectedStatus: lint.Pass,
		},
		{
			path:           "sanPrivatePublicSuffix.pem",
			expectedStatus: lint.Pass,
		},
		{
			path:           "SANGoodSuffix.pem",
			expectedStatus: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			result := test.TestLint("n_san_iana_pub_suffix_empty", tc.path)
			if result.Status != tc.expectedStatus {
				t.Errorf("expected status %v was %v", tc.expectedStatus, result.Status)
			}
			if result.Details != tc.expectedDetails {
				t.Errorf("expected details %v was %v", tc.expectedDetails, result.Details)
			}
		})
	}
}
