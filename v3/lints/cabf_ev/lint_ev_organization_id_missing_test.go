package cabf_ev

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

func TestOrganizationIDMissing(t *testing.T) {
	var tests = map[string]lint.LintStatus{
		"evOrgIdExtMissing_NoOrgId.pem":                                   lint.NA,
		"evOrgIdExtMissing_CABFOrgIdExtMissingButBeforeEffectiveDate.pem": lint.NE,
		"evOrgIdExtMissing_ValidButBeforeEffectiveDate.pem":               lint.NE,
		"evOrgIdExtMissing_Invalid.pem":                                   lint.Error,
		"evOrgIdExtMissing_Valid.pem":                                     lint.Pass,
	}
	for file, want := range tests {
		t.Run(file, func(t *testing.T) {
			t.Parallel()
			got := test.TestLint("e_ev_organization_id_missing", file).Status
			if got != want {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
