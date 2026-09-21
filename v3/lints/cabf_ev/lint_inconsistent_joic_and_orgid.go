/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

package cabf_ev

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"

	"errors"
	"unicode/utf8"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_inconsistent_joic_and_orgid",
			Description:   "Checks that organizationIdentifier and jurisdictionCountry are consistent",
			Citation:      "EVGL, Appendix H",
			Source:        lint.CABFEVGuidelines,
			EffectiveDate: util.SC17EffectiveDate,
		},
		Lint: NewInconsistentJoiCountryAndOrgID,
	})
}

type InconsistentJoiCountryAndOrgID struct{}

func NewInconsistentJoiCountryAndOrgID() lint.LintInterface {
	return &InconsistentJoiCountryAndOrgID{}
}

func (l *InconsistentJoiCountryAndOrgID) CheckApplies(c *x509.Certificate) bool {
	return len(c.Subject.JurisdictionCountry) > 0 && len(c.Subject.OrganizationIDs) > 0
}

type orgIdPartsType struct {
	scheme  string
	country string
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

func getOrgIdParts(orgId string) (*orgIdPartsType, error) {
	if len(orgId) < 5 || !isASCII(orgId[:5]) {
		return nil, errors.New("The Subject::organizationIdentifier attribute has an invalid value")
	}
	scheme := orgId[:3]
	country := orgId[3:5]

	return &orgIdPartsType{scheme, country}, nil
}

func (l *InconsistentJoiCountryAndOrgID) Execute(c *x509.Certificate) *lint.LintResult {

	// Let's assume there is just one OrganizationID
	// If not so, it's not this lint's business to raise an alarm
	orgId := c.Subject.OrganizationIDs[0]

	orgIdParts, err := getOrgIdParts(orgId)
	if err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		}
	}

	if orgIdParts.scheme != "NTR" {
		return &lint.LintResult{Status: lint.Pass}
	}

	// It should be safe to assume there is just one instance of joiCountry
	joiCountry := c.Subject.JurisdictionCountry[0]

	if orgIdParts.country != joiCountry {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "Subject::organizationIdentifier and jurisdictionCountry are inconsistent",
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
