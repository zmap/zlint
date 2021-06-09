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
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/************************************************
CA/Browser Forum EV Guidelines v1.7.0, Section 9.2.8

The Registration Scheme MUST be identified using the using the following structure in the presented order:

- 3 character Registration Scheme identifier;
- 2 character ISO 3166 country code for the nation in which the Registration Scheme is operated, or if the scheme is operated globally ISO 3166 code "XG" shall be used;
- For the NTR Registration Scheme identifier, if required under Section 9.2.4, a 2 character ISO 3166-2 identifier for the subdivision (state or province) of the nation in which the Registration Scheme is operated, preceded by plus "+" (0x2B (ASCII), U+002B (UTF-8));
- a hyphen-minus "-" (0x2D (ASCII), U+002D (UTF-8));
- Registration Reference allocated in accordance with the identified Registration Scheme

Examples:
NTRGB-12345678 (NTR scheme, Great Britain, Unique Identifier at Country level is 12345678)
NTRUS+CA-12345678 (NTR Scheme, United States - California, Unique identifier at State level is 12345678)
VATDE-123456789 (VAT Scheme, Germany, Unique Identifier at Country Level is 12345678)
PSDBE-NBB-1234.567.890 (PSD Scheme, Belgium, NCA's identifier is NBB, Subject Unique Identifier as-signed by the NCA is 1234.567.890)
************************************************/

type evSubjectOrganizationIdentifierWellFormed struct{}

func (l *evSubjectOrganizationIdentifierWellFormed) Initialize() error {
	return nil
}

func (l *evSubjectOrganizationIdentifierWellFormed) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers) && util.TypeInName(&c.Subject, util.OrganizationIdentifierOID)
}

func (l *evSubjectOrganizationIdentifierWellFormed) Execute(c *x509.Certificate) *lint.LintResult {

	var seq pkix.RDNSequence

	if _, err := asn1.Unmarshal(c.RawSubject, &seq); err != nil {
		return &lint.LintResult{Status: lint.Fatal}
	}

	for _, rdn := range seq {
		for _, atv := range rdn {
			if atv.Type.Equal(util.OrganizationIdentifierOID) {

				value, _ := atv.Value.(string)
				errorMessage, _ := util.ParseOrganizationIdentifier(value, false)

				if len(errorMessage) == 0 {
					return &lint.LintResult{Status: lint.Pass}
				} else {
					return &lint.LintResult{Status: lint.Error, Details: errorMessage}
				}
			}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ev_subject_organization_identifier_well_formed",
		Description:   "Checks that the content of subject:organizationIdentifier is well-formed and compliant to the specified format. The Registration Scheme MUST be identified using the using the following structure in the presented order.",
		Citation:      "CA/Browser Forum EV Guidelines v1.7.0, Section 9.2.8",
		Source:        lint.CABFEVGuidelines,
		EffectiveDate: util.CABFEV_1_7_0_Date,
		Lint:          &evSubjectOrganizationIdentifierWellFormed{},
	})
}
