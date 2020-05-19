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
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
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

func (l *evOrgIdEncoding) Execute(c *x509.Certificate) *lint.LintResult {

	rdnSequence := util.RawRDNSequence{}
	rest, err := asn1.Unmarshal(c.RawSubject, &rdnSequence)
	if err != nil {
		return &lint.LintResult{
			Status:  lint.Fatal,
			Details: "Failed to Unmarshal RawSubject into RawRDNSequence",
		}
	}
	if len(rest) > 0 {
		return &lint.LintResult{
			Status:  lint.Fatal,
			Details: "Trailing data after RawSubject RawRDNSequence",
		}
	}

	for _, attrTypeAndValueSet := range rdnSequence {
		for _, attrTypeAndValue := range attrTypeAndValueSet {
			if attrTypeAndValue.Type.Equal(util.OrganizationIdentifierOID) {
				if !(attrTypeAndValue.Value.Tag == asn1.TagPrintableString || attrTypeAndValue.Value.Tag == asn1.TagUTF8String) {
					return &lint.LintResult{Status: lint.Error, Details: "invalid string type in subject:organizationIdentifier"}
				}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ev_orgid_encoding",
		Description:   "The organizationIdentifier MUST be encoded as a PrintableString or UTF8String",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Sec. 9.2.8",
		Source:        lint.CABFEVGuidelines,
		EffectiveDate: util.CABAltRegNumEvDate,
		Lint:          &evOrgIdEncoding{},
	})
}
