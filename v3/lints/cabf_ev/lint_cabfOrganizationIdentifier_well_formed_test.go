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

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestEvCabfOrganizationIdentifierWellFormed(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with the cabfOrganizationIdentifier extension with correct data",
			InputFilename:  "evWithCABFOrgIdExtValid.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with the cabfOrganizationIdentifier extension with correct data without the optional registrationStateOrProvince",
			InputFilename:  "evWithCABFOrgIdExtValidWithoutRegistrationStateOrProvince.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with the cabfOrganizationIdentifier extension and issued before effective date",
			InputFilename:  "evWithCABFOrgIdExtIssuedBeforeEffectiveDate.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with the cabfOrganizationIdentifier extension with registrationSchemeIdentifier BAT (unsupported)",
			InputFilename:  "evWithCABFOrgIdExtUnsupportedRSI.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with the cabfOrganizationIdentifier extension but registrationSchemeIdentifier is encoded as UTF8String",
			InputFilename:  "evWithCABFOrgIdExtWrongEncodingInRSI.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with the cabfOrganizationIdentifier extension but registrationCountry is too short (1 character)",
			InputFilename:  "evWithCABFOrgIdRegCountryTooShort.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "EV certificate, issued after 21 June 2019 (EV 1.7.0 effective), with the cabfOrganizationIdentifier extension but registrationStateOrProvince is too long (129 characters)",
			InputFilename:  "evWithCABFOrgIdRegSTOrLTooLong.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ev_cabfOrganizationIdentifier_well_formed", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
