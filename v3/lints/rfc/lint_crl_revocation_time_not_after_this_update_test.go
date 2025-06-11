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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestRevocationTimeNotAfterThisUpdate(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		path string
		want lint.LintStatus
	}{
		{
			name: "crl revocation time before thisUpdate",
			path: "crl_revocation_time_before_this_update.pem",
			want: lint.Pass,
		},
		{
			name: "crl revocation time equals thisUpdate",
			path: "crl_revocation_time_equals_this_update.pem",
			want: lint.Pass,
		},
		{
			name: "crl revocation time after thisUpdate",
			path: "crl_revocation_time_after_this_update.pem",
			want: lint.Error,
		},
		{
			name: "CRL with no revoked certificates",
			path: "crl_no_revoked_certs.pem",
			want: lint.NA,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := test.TestRevocationListLint(t, "e_crl_revocation_time_after_this_update", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
