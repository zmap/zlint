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
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestUTF8ReplacementCharInSubject(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Clean certificate",
			path: "utf8_replac_char_in_subj_clean.pem",
			want: lint.Pass,
		},
		{
			desc: "Certificate with dirty char in subject:stateOrProvinceName",
			path: "utf8_replac_char_in_subj_dirty1.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with dirty char in subject:localityName",
			path: "utf8_replac_char_in_subj_dirty2.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with dirty char in subject:givenName",
			path: "utf8_replac_char_in_subj_dirty3.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("e_utf8_replacement_char_in_subject", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
