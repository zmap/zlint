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

package community

import (
	"strings"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestUniqueRevokedCertificate(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		path       string
		want       lint.LintStatus
		wantSubStr string
	}{
		{
			name: "CRL with duplicates in revoked certificates list",
			path: "crlWithDuplicatesInRevokedCertificateList.pem",
			want: lint.Warn,
		},
		{
			name: "CRL with no duplicates in revoked certificates list",
			path: "crlWithNoDuplicatesInRevokedCertificateList.pem",
			want: lint.Pass,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotStatus := test.TestRevocationListLint(t, "e_crl_unique_revoked_certificate", tc.path)
			if tc.want != gotStatus.Status {
				t.Errorf("%s: expected %s, got %s", tc.path, tc.want, gotStatus.Status)
			}
			if !strings.Contains(gotStatus.Details, tc.wantSubStr) {
				t.Errorf("%s: expected %s, got %s", tc.path, tc.wantSubStr, gotStatus.Details)
			}
		})
	}
}
