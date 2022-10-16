/*
 * ZLint Copyright 2022 Regents of the University of Michigan
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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSuperfluousKuEncoding(t *testing.T) {
	testCases := []struct {
		name           string
		filepath       string
		expectedStatus lint.LintStatus
	}{
		{
			name:           "Known Trustwave P256 with trailing zero byte in KU",
			filepath:       "trustwaveP256CASuperfluousBytesOnKU.pem",
			expectedStatus: lint.Error,
		},
		{
			name:           "Known Trustwave P256 with trailing zero byte in KU",
			filepath:       "trustwaveP384CASuperfluousBytesOnKU.pem",
			expectedStatus: lint.Error,
		},
		{
			name:           "A cert with CertSign | CRLSign and no trailing zery byte",
			filepath:       "keyUsageWithoutTrailingZeroes.pem",
			expectedStatus: lint.Pass,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := test.TestLint("e_superfluous_ku_encoding", tc.filepath)
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v", tc.expectedStatus, result.Status)
			}
		})
	}
}
