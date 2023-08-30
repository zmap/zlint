package cabf_br

import (
	"strings"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

func TestCrlValidReasonCodes(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		path       string
		want       lint.LintStatus
		wantSubStr string
	}{
		{
			name:       "CRL with reason code 0",
			path:       "crlWithReasonCode0.pem",
			want:       lint.Error,
			wantSubStr: "The reason code CRL entry extension SHOULD be absent instead of using the unspecified",
		},
		{
			// This test case is significant since reason code 2 is not allowed by CABF
			name:       "CRL with reason code 2",
			path:       "crlWithReasonCode2.pem",
			want:       lint.Error,
			wantSubStr: "Reason code not included in BR: 7.2.2",
		},
		{
			name: "CRL with reason code 5",
			path: "crlWithReasonCode5.pem",
			want: lint.Pass,
		},
		{
			name:       "CRL with reason code 7",
			path:       "crlWithReasonCode7.pem",
			want:       lint.Error,
			wantSubStr: "Reason code not included in BR: 7.2.2",
		},
		{
			name: "CRL thisUpdate before enforcement",
			path: "crlThisUpdate20230505.pem",
			want: lint.NE,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotStatus := test.TestRevocationListLint(t, "e_cab_crl_has_valid_reason_code", tc.path)
			if tc.want != gotStatus.Status {
				t.Errorf("%s: expected %s, got %s", tc.path, tc.want, gotStatus.Status)
			}
			if !strings.Contains(gotStatus.Details, tc.wantSubStr) {
				t.Errorf("%s: expected %s, got %s", tc.path, tc.wantSubStr, gotStatus.Details)
			}
		})
	}
}
