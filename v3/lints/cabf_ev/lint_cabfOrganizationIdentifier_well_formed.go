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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/************************************************
CA/Browser Forum EV Guidelines v1.7.0, Sections 9.2.8 and 9.8.2

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

type evCabfOrganizationIdentifierWellFormed struct{}

func (l *evCabfOrganizationIdentifierWellFormed) Initialize() error {
	return nil
}

func (l *evCabfOrganizationIdentifierWellFormed) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers) && util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier)
}

func (l *evCabfOrganizationIdentifierWellFormed) Execute(c *x509.Certificate) *lint.LintResult {

	parseErrorMessage, parsedExtension := util.ParseCabfOrgIdExt(c)

	if parseErrorMessage != "" {
		return &lint.LintResult{Status: lint.Error, Details: parseErrorMessage}
	}

	checkErrorMessage := util.CheckParsedEvOrgId(parsedExtension)

	if checkErrorMessage != "" {
		return &lint.LintResult{Status: lint.Error, Details: checkErrorMessage}
	}

	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ev_cabfOrganizationIdentifier_well_formed",
		Description:   "Checks that the content of the cabfOrganizationIdentifier extension is well-formed and compliant to the specified format and that it is encoded according to the defined ASN.1 module.",
		Citation:      "CA/Browser Forum EV Guidelines v1.7.0, Section 9.8.2",
		Source:        lint.CABFEVGuidelines,
		EffectiveDate: util.CABFEV_1_7_0_Date,
		Lint:          &evCabfOrganizationIdentifierWellFormed{},
	})
}
