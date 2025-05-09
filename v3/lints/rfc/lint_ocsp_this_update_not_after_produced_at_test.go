package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

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

func TestOCSPThisUpdateNotAfterProducedAt(t *testing.T) {
	cases := []struct {
		inputPath string
		want      lint.LintStatus
	}{
		{
			inputPath: "ocspThisUpdateNotAfterProducedAt",
			want:      lint.Pass,
		},
		{
			inputPath: "ocspThisUpdateAfterProducedAt",
			want:      lint.Error,
		},
	}
	for _, tc := range cases {
		t.Run(tc.inputPath, func(t *testing.T) {
			got := test.TestOCSPResponseLint(t, "e_this_update_not_after_produced_at", tc.inputPath).Status
			if tc.want != got {
				t.Errorf("%s: expected %s, got %s", tc.inputPath, tc.want, got)
			}
		})
	}
}
