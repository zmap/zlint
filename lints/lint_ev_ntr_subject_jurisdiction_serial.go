/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evNtrSubjectJurisdiction struct{}

func (l *evNtrSubjectJurisdiction) Initialize() error {
	return nil
}

func GetOrgIdFromSubjOrExt(c *x509.Certificate) (bool, util.ParsedEvOrgId) {

	var result util.ParsedEvOrgId

	subjOrgId := util.GetSubjectOrgId(c.RawSubject)
	if util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier) {
		_, result = util.ParseCabfOrgIdExt(c) // check error string during Execute

	} else if subjOrgId.IsPresent {
		_, result = util.ParseCabfOrgId(subjOrgId.Value)
	} else {
		return false, result
	}
	return true, result
}
func (l *evNtrSubjectJurisdiction) CheckApplies(c *x509.Certificate) bool {
	// check whether the cert is an EV cert and features either a subject:organizationIdentifier or the CabfExtensionOrganizationIdentifier Extension with scheme ID "NTR". If at least one of the two is present, the subject:SerialNumber in its role as subject registration number, and subject:jurisdiction... values are checked against the organizationIdentifier values
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

func (l *evNtrSubjectJurisdiction) Execute(c *x509.Certificate) *LintResult {

	_, parsedOrgId := GetOrgIdFromSubjOrExt(c) // one must be present, otherwise lint would not be invoked

	if util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier) {
		errStr, _ := util.ParseCabfOrgIdExt(c)
		if errStr != "" {
			return &LintResult{Status: Error, Details: errStr}
		}
	}

	// perform checks against subject:jurisdiction fields in case of NTR:
	jurSop := util.GetSubjectElement(c.RawSubject, util.SubjectJurisdictionStateOrProvinceNameOID)
	if jurSop.ErrorString != "" {
		return &LintResult{Status: Error, Details: jurSop.ErrorString}
	}
	if jurSop.IsPresent && parsedOrgId.StateOrProvince == "" {
		return &LintResult{Status: Error, Details: "subject:jurisdictionStateOrProvince is present, but subject:organizationIdentifier does not feature stateOrProvince"}
	}
	if parsedOrgId.StateOrProvince != "" && !jurSop.IsPresent {
		return &LintResult{Status: Error, Details: "subject:organizationIdentifier features stateOrProvince but subject:jurisdictionStateOrProvince is not present"}
	}
	jurCn := util.GetSubjectElement(c.RawSubject, util.SubjectJurisdictionCountryNameOID)
	if jurCn.ErrorString != "" {
		return &LintResult{Status: Error, Details: jurCn.ErrorString}
	}
	if jurCn.IsPresent && (jurCn.Value != parsedOrgId.Country) {
		return &LintResult{Status: Error, Details: "subject:organizationIdentifier features different country code than subject:jurisdictionCountryName"}
	}
	jurLoc := util.GetSubjectElement(c.RawSubject, util.SubjectJurisdictionLocalityNameOID)
	if jurLoc.ErrorString != "" {
		return &LintResult{Status: Error, Details: jurLoc.ErrorString}
	}
	if c.Subject.SerialNumber != parsedOrgId.RegRef {
		return &LintResult{Status: Error, Details: "subject:SerialNumber does not contain the registration reference found in the subject:organizationIdentifier"}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_ntr_subject_jurisdiction_serial",
		Description:   "Checks that the consistency requirements about NTR regarding subject:jurisdiction...Name fields, subject:SerialNumber, and organization identifier are satisfied",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Appendix H",
		Source:        CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evNtrSubjectJurisdiction{},
	})
}
