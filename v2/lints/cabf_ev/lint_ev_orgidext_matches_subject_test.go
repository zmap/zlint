package cabf_ev

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestEvAltRegNumOrgIdExtMatchesSubject(t *testing.T) {
	m := map[string]lint.LintStatus{
		"EvAltRegNumCert52NoOrgId.pem":                     lint.NA,
		"EvAltRegNumCert56JurContryNotMatching.pem":        lint.NA,
		"EvAltRegNumCert67ValidNtrWithOrgIdExt.pem":        lint.Pass,
		"EvAltRegNumCert68OrgIdExtWrongSchemeId.pem":       lint.Error,
		"EvAltRegNumCert69OrgIdExtWrongCountry.pem":        lint.Error,
		"EvAltRegNumCert70ValidOrgIdExtSopNotMatching.pem": lint.Pass,
		"EvAltRegNumCert71OrgIdExtWrongSchemeId.pem":       lint.Error,
		"EvAltRegNumCert72OrgIdExtRegRefZeroLen.pem":       lint.Error,
		"EvAltRegNumCert73OrgIdExtWrongRegRef.pem":         lint.Error,
		"EvAltRegNumCert75ValidOrgIdExt.pem":               lint.Pass,
		"EvAltRegNumCert76OrgIdExtWrongEncoding.pem":       lint.Error,
		"EvAltRegNumCert77OrgIdExtWrongEncoding.pem":       lint.Error,
		"EvAltRegNumCert78OrgIdExtWrongEncoding.pem":       lint.Error,
		"EvAltRegNumCert79OrgIdExtWrongEncoding.pem":       lint.Error,
		"doubleOI.pem":                         lint.Error,
		"emptyRegistrationStateOrProvince.pem": lint.Error,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_ev_orgidext_matches_subject", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
