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

func TestEvSubjectOrganizationIdentifierWellFormed(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), without a subject:organizationIdentifier",
			InputFilename:  "EvAltRegNumCert52NoOrgId.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with an invalid subject:organizationIdentifier (=ABCDE-12345678)",
			InputFilename:  "EvAltRegNumCert53OrgIdInvalid.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with a valid subject:organizationIdentifier (=VATXG-028947476)",
			InputFilename:  "EvAltRegNumCert61Valid.pem",
			ExpectedResult: lint.Pass,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ev_subject_organization_identifier_well_formed", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
