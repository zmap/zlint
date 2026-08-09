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

var psd2OrgIdFormatRegexShall = regexp.MustCompile(`^PSD([A-Z]{2})-([A-Z]{2,8})-(.+)$`)

type qcStatemPsd2OrgIdFormatShall struct{}

// ETSI TS 119 495 V1.1.2 (2018-07) through V1.4.1 (2019-11), Section 5.2.1:
//
//	GEN-5.2.1-3: The organizationIdentifier attribute shall contain
//	information using the following structure in the presented order:
//	"PSD" as 3 character legal person identity type reference; 2 character
//	ISO 3166 country code representing the NCA country; hyphen-minus "-";
//	2-8 character NCA identifier (A-Z uppercase only, no separator); and
//	hyphen-minus "-"; and PSP identifier (authorization number as
//	specified by the NCA). There are no restrictions on the characters
//	used [for the PSP identifier].
//
// This clause was downgraded from "shall" to "should" in V1.5.1
// (2021-04) onward — this lint only covers the "shall" era. See
// w_qcstatem_psd2_orgid_format for the "should" era that follows it.
func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:            "e_qcstatem_psd2_orgid_format",
			Description:     "Checks that a PSD2 certificate's subject:organizationIdentifier, when it starts with 'PSD', matches the required PSD<country>-<NCAid>-<PSPid> structure",
			Citation:        "ETSI TS 119 495 V1.1.2 (2018-07) - V1.4.1 (2019-11), Section 5.2.1, GEN-5.2.1-3",
			Source:          lint.EtsiEsi,
			EffectiveDate:   util.EtsiTs119495_V1_1_2_Date,
			IneffectiveDate: util.EtsiTs119495_V1_5_1_Date,
		},
		Lint: NewQcStatemPsd2OrgIdFormatShall,
	})
}

func NewQcStatemPsd2OrgIdFormatShall() lint.LintInterface {
	return &qcStatemPsd2OrgIdFormatShall{}
}

func (l *qcStatemPsd2OrgIdFormatShall) CheckApplies(c *x509.Certificate) bool {
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

// Unlike the other PSD2 lints, this deliberately does not defer on
// GetErrorInfo(). The subject matter here is the Subject DN, which is
// parsed independently of the qcStatements extension; a garbled
// QCStatement neither makes OrganizationIDs unsafe to read nor makes the
// org-id format requirement inapplicable. (Note ParseQcStatem reports
// IsPresent() for any OID when the outer SEQUENCE fails to parse, so a
// "PSD"-prefixed org-id is what actually establishes PSD2 applicability
// here.)
func (l *qcStatemPsd2OrgIdFormatShall) Execute(c *x509.Certificate) *lint.LintResult {
	orgId := c.Subject.OrganizationIDs[0]
	m := psd2OrgIdFormatRegexShall.FindStringSubmatch(orgId)
	if m == nil {
		return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf(
			"subject:organizationIdentifier %q does not match the required PSD<country>-<NCAid>-<PSPid> structure", orgId)}
	}
	if !util.IsISOCountryCode(m[1]) {
		return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf(
			"subject:organizationIdentifier %q has a country code that is not an assigned ISO 3166-1 country", orgId)}
	}
	return &lint.LintResult{Status: lint.Pass}
}
