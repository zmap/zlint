package etsi

/*
 * ZLint Copyright 2025 Regents of the University of Michigan
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
	"fmt"
	"regexp"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type qcStatemNationalScheme struct{}

/************************************************************************
ETSI EN 319 412-2 V2.2.1 (2020-07)
https://www.etsi.org/deliver/etsi_en/319400_319499/31941202/02.02.01_60/en_31941202v020201p.pdf
4.2.4 Certificates may include one or more semantics identifiers as specified in ETSI
EN 319 412-1 [i.4], clause 5 which defines the semantics for the organizationIdentifier attribute.

Certificates may include one or more semantics identifiers as specified in ETSI EN 319 412-1 [i.4],
clause 5 which define the semantics for the serialNumber attribute.

ETSI EN 319 412-1 V1.4.1 (2020-06)
https://www.etsi.org/deliver/etsi_en/319400_319499/31941201/01.04.01_60/en_31941201v010401p.pdf
5.1.3 Natural person semantics identifier
The semantics of id-etsi-qcs-SemanticsId-Natural shall be as follows.
When the natural person semantics identifier is included, any present serialNumber attribute in the subject field shall
contain information using the following structure in the presented order:
The three initial characters shall have one of the following defined values:
1) "PAS" for identification based on passport number.
2) "IDC" for identification based on national identity card number.
3) "PNO" for identification based on (national) personal number (national civic registration number).
4) "TAX" for identification based on a personal tax reference number issued by a national tax authority. This
value is deprecated. The value "TIN" should be used instead.
5) "TIN" Tax Identification Number according to the European Commission – Tax and Customs Union
(https://ec.europa.eu/taxation_customs/tin/tinByCountry.html).
6) Two characters according to local definition within the specified country and name registration authority,
identifying a national scheme that is considered appropriate for national and European level, followed by the
character ":" (colon).

5.1.4 Legal person semantics identifier
The semantics of id-etsi-qcs-SemanticsId-Legal shall be as follows.
When the legal person semantics identifier is included, any present organizationIdentifier attribute in the subject
field shall contain information using the following structure in the presented order:
• 3 character legal person identity type reference;
• 2 character ISO 3166 [2] country code;
• hyphen-minus "-" (0x2D (ASCII), U+002D (UTF-8)); and
• identifier (according to country and identity type reference).
The three initial characters shall have one of the following defined values:
1) "VAT" for identification based on a national value added tax identification number.
2) "NTR" for identification based on an identifier from a national trade register.
3) "PSD" for identification based on national authorization number of a payment service provider under
Payments Services Directive (EU) 2015/2366 [i.13]. This shall use the extended structure as defined in ETSI
TS 119 495 [3], clause 5.2.1.
4) "LEI" for a global Legal Entity Identifier as specified in ISO 17442 [4]. The 2 character ISO 3166 [2] country
code shall be set to 'XG'.
5) Two characters according to local definition within the specified country and name registration authority,
identifying a national scheme that is considered appropriate for national and European level, followed by the
character ":" (colon)
*************************************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_qcstatem_correct_national_scheme",
			Description:   "This lint checks that the national scheme is well-formed when used in the serialNumber or organizationIdentifier attribute in the subject field, provided that either the natural person semantics identifier or the legal person semantics identifier is present.",
			Citation:      "ETSI EN 319 412-1 V1.4.1, 5.1.3 Natural person semantics identifier and 5.1.4 Legal person semantics identifier",
			Source:        lint.EtsiEsi,
			EffectiveDate: util.ETSI_EN_319_412_1_V1_4_1_DATE,
		},
		Lint: NewQcStatemNationalScheme,
	})
}

func NewQcStatemNationalScheme() lint.LintInterface {
	return &qcStatemNationalScheme{}
}

func (l *qcStatemNationalScheme) CheckApplies(c *x509.Certificate) bool {
	_, isPresent := util.IsQcStatemPresent(c, &util.IdQcsPkixQCSyntaxV2)

	if !isPresent {
		return false
	}

	qcs2Generic := util.ParseQcStatem(util.GetQcStatemExtValue(c), util.IdQcsPkixQCSyntaxV2)

	qcs2 := qcs2Generic.(util.DecodedQcS2)
	semanticsId := qcs2.Decoded.SemanticsId
	re := regexp.MustCompile(`^.{2}:`)

	if semanticsId.Equal(util.IdEtsiQcsSemanticsIdNatural) {
		serialNumber := c.Subject.SerialNumber
		return re.MatchString(serialNumber)
	}

	if semanticsId.Equal(util.IdEtsiQcsSemanticsIdLegal) {
		for _, orgId := range c.Subject.OrganizationIDs {
			if re.MatchString(orgId) {
				return true
			}
		}
	}
	return false
}

func (l *qcStatemNationalScheme) Execute(c *x509.Certificate) *lint.LintResult {

	qcs2Generic := util.ParseQcStatem(util.GetQcStatemExtValue(c), util.IdQcsPkixQCSyntaxV2)

	qcs2 := qcs2Generic.(util.DecodedQcS2)
	semanticsId := qcs2.Decoded.SemanticsId

	if semanticsId.Equal(util.IdEtsiQcsSemanticsIdNatural) {
		serialNumber := c.Subject.SerialNumber
		if !util.CheckNationalScheme(serialNumber) {
			return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("invalid format of subject:serialNumber %s for national scheme", serialNumber)}
		}
	}

	if semanticsId.Equal(util.IdEtsiQcsSemanticsIdLegal) {
		re := regexp.MustCompile(`^.{2}:`)
		for _, orgId := range c.Subject.OrganizationIDs {
			if re.MatchString(orgId) && !util.CheckNationalScheme(orgId) {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("invalid format of subject:organizationIdentifier %s for national scheme", orgId)}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
