/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

package etsi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

var psd2OrgIdFormatRegexShould = regexp.MustCompile(`^PSD([A-Z]{2})-([A-Z]{2,8})-(.+)$`)

type qcStatemPsd2OrgIdFormatShould struct{}

// ETSI TS 119 495 V1.5.1 (2021-04) onward (verified current as of
// V1.8.1, 2026-04), Section 5.2.1:
//
//	GEN-5.2.1-3: If an Authorization Number was issued by a Competent
//	Authority the subject organizationIdentifier attribute should
//	contain the Authorization Number encoded using the following
//	structure in the presented order: "PSD" as 3 character legal person
//	identity type reference; 2 character ISO 3166-1 country code
//	representing the Competent Authority country; hyphen-minus "-";
//	2-8 character Competent Authority identifier without country code
//	(A-Z uppercase only, no separator); and hyphen-minus "-"; and
//	identifier (authorization number as specified by the Competent
//	Authority; no restrictions on the characters used).
//
// This clause was "shall" (MUST) from V1.1.2 (2018-07) through V1.4.1
// (2019-11) — see e_qcstatem_psd2_orgid_format for that earlier era.
func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "w_qcstatem_psd2_orgid_format",
			Description:   "Warns if a PSD2 certificate's subject:organizationIdentifier, when it starts with 'PSD', does not match the recommended PSD<country>-<NCAid>-<PSPid> structure",
			Citation:      "ETSI TS 119 495 V1.5.1 (2021-04) onward, Section 5.2.1, GEN-5.2.1-3",
			Source:        lint.EtsiEsi,
			EffectiveDate: util.EtsiTs119495_V1_5_1_Date,
		},
		Lint: NewQcStatemPsd2OrgIdFormatShould,
	})
}

func NewQcStatemPsd2OrgIdFormatShould() lint.LintInterface {
	return &qcStatemPsd2OrgIdFormatShould{}
}

func (l *qcStatemPsd2OrgIdFormatShould) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	if !util.ParseQcStatem(util.GetExtFromCert(c, util.QcStateOid).Value, util.IdEtsiPsd2Statem).IsPresent() {
		return false
	}
	if len(c.Subject.OrganizationIDs) == 0 {
		return false
	}
	return strings.HasPrefix(c.Subject.OrganizationIDs[0], "PSD")
}

func (l *qcStatemPsd2OrgIdFormatShould) Execute(c *x509.Certificate) *lint.LintResult {
	orgId := c.Subject.OrganizationIDs[0]
	m := psd2OrgIdFormatRegexShould.FindStringSubmatch(orgId)
	if m == nil {
		return &lint.LintResult{Status: lint.Warn, Details: fmt.Sprintf(
			"subject:organizationIdentifier %q does not match the required PSD<country>-<NCAid>-<PSPid> structure", orgId)}
	}
	if !util.IsISOCountryCode(m[1]) {
		return &lint.LintResult{Status: lint.Warn, Details: fmt.Sprintf(
			"subject:organizationIdentifier %q has a country code that is not an assigned ISO 3166-1 country", orgId)}
	}
	return &lint.LintResult{Status: lint.Pass}
}
