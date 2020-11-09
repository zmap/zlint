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
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type evOrgIdExtMatchesSubject struct{}

func (l *evOrgIdExtMatchesSubject) Initialize() error {
	return nil
}

func (l *evOrgIdExtMatchesSubject) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier)
}

func (l *evOrgIdExtMatchesSubject) Execute(c *x509.Certificate) *lint.LintResult {
	orgId := util.GetSubjectOrgId(c.RawSubject)

	if orgId.ErrorString != "" {
		return &lint.LintResult{Status: lint.Error, Details: orgId.ErrorString}
	}

	errStr, parsedExt := util.ParseCabfOrgIdExt(c)
	if errStr != "" {
		return &lint.LintResult{Status: lint.Error, Details: errStr}
	}
	if !orgId.IsPresent {
		return &lint.LintResult{Status: lint.Pass}
	}
	_, parsedOrgId := util.ParseCabfOrgId(orgId.Value, false)
	// no need to check an error parsing the subject:organizationIdentifier here, this is done in other lints dealing with that field explicitly.
	// StateOrProvince doesn't have to match literally, only semantically
	if parsedExt.Rsi != parsedOrgId.Rsi || parsedExt.Country != parsedOrgId.Country || parsedExt.RegRef != parsedOrgId.RegRef {
		return &lint.LintResult{Status: lint.Error, Details: "values in subject:organizationIdentifier and CAB/F organizationIdentifier Extension do not match"}
	}

	// if a state or province is present in subject:organizationIdentifier
	// then the corresponding field in the extension must not be blank
	if parsedOrgId.StateOrProvince != "" {
		isStateOrProvBlank := strings.TrimSpace(parsedExt.StateOrProvince) == ""
		if isStateOrProvBlank {
			return &lint.LintResult{Status: lint.Error, Details: "Empty registrationStateOrProvince found but present in subject:organizationIdentifier"}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ev_orgidext_matches_subject",
		Description:   "The contents of the cabfOrganizationIdentifier extension must match the entries of the corresponding field in the subject DN.",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Sec. 9.2.8, 9.8.2",
		Source:        lint.CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evOrgIdExtMatchesSubject{},
	})
}
