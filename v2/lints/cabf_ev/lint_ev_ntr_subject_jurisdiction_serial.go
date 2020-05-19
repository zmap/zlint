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
CA/Browser Forum EV Guidelines v1.7, Appendix H

This lints considers the additional consistency requirements in the case of the NTR registration scheme.

NTR: The information carried in this field shall be the same as held in Subject Registration Number Field as specified
in 9.2.5 and the country code used in the Registration Scheme identifier shall match that of the subject’s jurisdiction
as specified in Section 9.2.4. Where the Subject Jurisdiction of Incorporation or Registration Field in 9.2.4 includes
more than the country code, the additional locality information shall be included as specified in sections 9.2.8
and/or 9.8.1.
************************************************/

type evNtrSubjectJurisdiction struct{}

func (l *evNtrSubjectJurisdiction) Initialize() error {
	return nil
}

func (l *evNtrSubjectJurisdiction) CheckApplies(c *x509.Certificate) bool {
	// check whether the cert is an EV cert and features either a subject:organizationIdentifier
	// or the CabfExtensionOrganizationIdentifier Extension with scheme ID "NTR". If at least one of the two is present,
	// the subject:SerialNumber in its role as subject registration number, and subject:jurisdiction... values are
	// checked against the organizationIdentifier values

	if !util.IsEV(c.PolicyIdentifiers) {
		return false
	}
	hasOrgId, orgId := GetOrgIdFromSubjOrExt(c)
	if !hasOrgId {
		return false
	}
	if orgId.Rsi == "NTR" {
		return true
	}
	return false
}

func (l *evNtrSubjectJurisdiction) Execute(c *x509.Certificate) *lint.LintResult {

	// see https://github.com/zmap/zlint/pull/336#discussion_r362983627
	// The correctness of this lint relies on lint_ev_orgidext_matches_subject.go also running.
	// If that did not run, this lint would no longer produce correct results, because the subject field
	// could be misencoded if the extension was present and encoded correctly.

	_, parsedOrgId := GetOrgIdFromSubjOrExt(c) // one must be present, otherwise lint would not be invoked

	if util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier) {
		errStr, _ := util.ParseCabfOrgIdExt(c)
		if errStr != "" {
			return &lint.LintResult{Status: lint.Error, Details: errStr}
		}
	}

	// perform checks against subject:jurisdiction fields in case of NTR:
	jurSop := util.GetSubjectElement(c.RawSubject, util.SubjectJurisdictionStateOrProvinceNameOID)
	if jurSop.ErrorString != "" {
		return &lint.LintResult{Status: lint.Error, Details: jurSop.ErrorString}
	}
	if jurSop.IsPresent && parsedOrgId.StateOrProvince == "" {
		return &lint.LintResult{Status: lint.Error, Details: "subject:jurisdictionStateOrProvince is present, but subject:organizationIdentifier does not feature stateOrProvince"}
	}
	if parsedOrgId.StateOrProvince != "" && !jurSop.IsPresent {
		return &lint.LintResult{Status: lint.Error, Details: "subject:organizationIdentifier features stateOrProvince but subject:jurisdictionStateOrProvince is not present"}
	}
	jurCn := util.GetSubjectElement(c.RawSubject, util.SubjectJurisdictionCountryNameOID)
	if jurCn.ErrorString != "" {
		return &lint.LintResult{Status: lint.Error, Details: jurCn.ErrorString}
	}
	if jurCn.IsPresent && (jurCn.Value != parsedOrgId.Country) {
		return &lint.LintResult{Status: lint.Error, Details: "subject:organizationIdentifier features different country code than subject:jurisdictionCountryName"}
	}
	jurLoc := util.GetSubjectElement(c.RawSubject, util.SubjectJurisdictionLocalityNameOID)
	if jurLoc.ErrorString != "" {
		return &lint.LintResult{Status: lint.Error, Details: jurLoc.ErrorString}
	}
	if c.Subject.SerialNumber != parsedOrgId.RegRef {
		return &lint.LintResult{Status: lint.Error, Details: "subject:SerialNumber does not contain the registration reference found in the subject:organizationIdentifier"}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ev_ntr_subject_jurisdiction_serial",
		Description:   "Checks that the consistency requirements about NTR regarding subject:jurisdiction...Name fields, subject:SerialNumber, and organization identifier are satisfied",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Appendix H",
		Source:        lint.CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evNtrSubjectJurisdiction{},
	})
}
