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

type evOrgIdExtPresentRecomm struct{}

func (l *evOrgIdExtPresentRecomm) Initialize() error {
	return nil
}

func (l *evOrgIdExtPresentRecomm) CheckApplies(c *x509.Certificate) bool {
	orgId := util.GetSubjectOrgId(c.RawSubject)
	if !util.IsEV(c.PolicyIdentifiers) || !orgId.IsPresent {
		return false
	}
	return true
}

func (l *evOrgIdExtPresentRecomm) Execute(c *x509.Certificate) *LintResult {
	if !util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier) {
		return &LintResult{Status: Warn, Details: "subject:organizationIdentifier field is present in an EV certificate but the CA/Browser Forum Organization Identifier Field Extension is missing"}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ev_orgidext_present_recommended",
		Description:   "If the subject:organizationIdentifier field is present, checks that the CAB/F organization identifier extension is also present.",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Sec. 9.8.2",
		Source:        CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evOrgIdExtPresentRecomm{},
	})
}