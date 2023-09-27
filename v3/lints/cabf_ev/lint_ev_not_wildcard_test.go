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

package cabf_ev

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSmoke(t *testing.T) {
	var tests = map[string]lint.LintStatus{
		"evWildcard.pem":                lint.Error,
		"evSubscriberNotWildCard.pem":   lint.Pass,
		"evSubscriberWildcardOnion.pem": lint.Pass,
	}
	for file, want := range tests {
		f := file
		w := want
		t.Run(f, func(t *testing.T) {
			t.Parallel()
			got := test.TestLint("e_ev_not_wildcard", f).Status
			if got != w {
				t.Errorf("want %s, got %s", w, got)
			}
		})
	}
}
