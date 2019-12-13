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

type evOrgIdEncoding struct{}

func (l *evOrgIdEncoding) Initialize() error {
	return nil
}

func (l *evOrgIdEncoding) CheckApplies(c *x509.Certificate) bool {
	orgId := util.GetSubjectOrgId(c.RawSubject)
	if !util.IsEV(c.PolicyIdentifiers) || !orgId.IsPresent {
		return false
	}
	return true
}

func (l *evOrgIdEncoding) Execute(c *x509.Certificate) *LintResult {
	orgId := util.GetSubjectOrgId(c.RawSubject)
	var errStr string
	reencParams := []string{"printable", "utf8"}
	foundCorrectEncoding := false
	for _, rp := range reencParams {
		errStr = util.CheckAsn1ReencodingWithParams(orgId.Value, orgId.Asn1RawValue.FullBytes, "error with subject:organizationIdentifier ASN.1 string type", rp)
		if errStr == "" {
			foundCorrectEncoding = true
			break
		}
	}
	if !foundCorrectEncoding {
		return &LintResult{Status: Error, Details: "invalid string type in subject:organizationIdentifier"}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_orgid_encoding",
		Description:   "If the subject:organizationIdentifier field is present in an EV certificate, then this lint checks that the format of its encoding (i.e. the string type) is in conformance to the CAB/F EV Guidelines",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Sec. 9.2.8",
		Source:        CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evOrgIdEncoding{},
	})
}
