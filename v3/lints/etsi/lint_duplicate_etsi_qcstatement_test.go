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

package etsi

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestDuplicateQCStatement(t *testing.T) {

	testCases := []struct {
		name string
		path string
		want lint.LintStatus
	}{
		{
			name: "Certificate without the QcStatements extension",
			path: "qcs0_dupx_effx.pem",
			want: lint.NA,
		},
		{
			name: "Certificate with valid QcStatements extension",
			path: "qcs1_dup0_eff1.pem",
			want: lint.Pass,
		},
		{
			name: "Certificate with invalid QcStatements, issued before the Effective Date",
			path: "qcs1_dup1_eff0.pem",
			want: lint.NE,
		},
		{
			name: "Certificate with invalid QcStatements, issued after the Effective Date",
			path: "qcs1_dup1_eff1.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := test.TestLint("e_duplicate_etsi_qcstatement", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
