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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

/************************************************
CA/Browser Forum EV Guidelines v1.7, Sections 9.2.8 and 9.8.2

The content of subject:jurisdiction and cabfOrganizationIdentifier extension should follow a specific format.

Examples:
NTRGB-12345678 (NTR scheme, Great Britain, Unique Identifier at Country level is 12345678)
NTRUS+CA-12345678 (NTR Scheme, United States - California, Unique identifier at State level is 12345678)
VATDE-123456789 (VAT Scheme, Germany, Unique Identifier at Country Level is 12345678)
PSDBE-NBB-1234.567.890 (PSD Scheme, Belgium, NCA's identifier is NBB, Subject Unique Identifier as-signed by the NCA is 1234.567.890)
************************************************/

type evOrgIdWellFormed struct{}

func GetOrgIdFromSubjOrExt(c *x509.Certificate) (bool, util.ParsedEvOrgId) {

	var result util.ParsedEvOrgId

	subjOrgId := util.GetSubjectOrgId(c.RawSubject)
	if util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier) {
		_, result = util.ParseCabfOrgIdExt(c) // check error string during Execute

	} else if subjOrgId.IsPresent {
		_, result = util.ParseCabfOrgId(subjOrgId.Value, false)
	} else {
		return false, result
	}
	return true, result
}

func (l *evOrgIdWellFormed) Initialize() error {
	return nil
}

func (l *evOrgIdWellFormed) CheckApplies(c *x509.Certificate) bool {
	// check whether the cert is an EV cert and features either a subject:organizationIdentifier
	// or the CabfExtensionOrganizationIdentifier Extension.

	if !util.IsEV(c.PolicyIdentifiers) {
		return false
	}
	hasOrgId, _ := GetOrgIdFromSubjOrExt(c)
	if !hasOrgId {
		return false
	} else {
		return true
	}
}

func (l *evOrgIdWellFormed) Execute(c *x509.Certificate) *lint.LintResult {

	// see https://github.com/zmap/zlint/pull/336#discussion_r362983627
	// The correctness of this lint relies on lint_ev_orgidext_matches_subject.go also running.
	// If that did not run, this lint would no longer produce correct results, because the subject field
	// could be misencoded if the extension was present and encoded correctly.

	subjOrgId := util.GetSubjectOrgId(c.RawSubject)

	if subjOrgId.IsPresent {
		var errStr string
		errStr, _ = util.ParseCabfOrgId(subjOrgId.Value, false)
		if errStr != "" {
			return &lint.LintResult{Status: lint.Error, Details: errStr}
		}
	}

	if util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier) {
		errStr, _ := util.ParseCabfOrgIdExt(c)
		if errStr != "" {
			return &lint.LintResult{Status: lint.Error, Details: errStr}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ev_orgid_well_formed",
		Description:   "Checks that the content of subject:jurisdiction and cabfOrganizationIdentifier extension are well-formed and compliant to the specified format.",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Sections 9.2.8 and 9.8.2",
		Source:        lint.CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evOrgIdWellFormed{},
	})
}
