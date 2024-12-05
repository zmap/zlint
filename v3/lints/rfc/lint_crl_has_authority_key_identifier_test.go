package rfc

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

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
		path string
		want lint.LintStatus
	}{
		{
			name: "crlWithMissingAuthKeyID",
			path: "crlWithMissingAuthKeyID.pem",
			want: lint.Error,
		},
		{
			name: "crlWithAuthKeyID",
			path: "crlWithAuthKeyID.pem",
			want: lint.Pass,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := test.TestRevocationListLint(t, "e_crl_has_authority_key_identifier", tc.path)
			if got.Status != tc.want {
				t.Errorf("Execute() = %v, want %v", got.Status, tc.want)
			}
		})
	}
}
