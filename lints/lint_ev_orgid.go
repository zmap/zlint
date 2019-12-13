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

type evOrgId struct{}

func (l *evOrgId) Initialize() error {
	return nil
}

func (l *evOrgId) CheckApplies(c *x509.Certificate) bool {
	_, isPsd2 := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	orgId := util.GetSubjectOrgId(c.RawSubject)
	if isPsd2 || !util.IsEV(c.PolicyIdentifiers) || !orgId.IsPresent {
		// PSD2 certificates follow different requirements
		return false
	}
	return true
}

func (l *evOrgId) Execute(c *x509.Certificate) *LintResult {
	orgId := util.GetSubjectOrgId(c.RawSubject)

	errStr, _ := util.ParseCabfOrgId(orgId.Value)
	if errStr != "" {
		return &LintResult{Status: Error, Details: errStr}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_orgid",
		Description:   "If the subject:organizationIdentifier field is present in an EV certificate, then this lint checks that the format of its contents is in conformance to the CAB/F EV Guidelines",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Sec. 9.2.8",
		Source:        CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evOrgId{},
	})
}
