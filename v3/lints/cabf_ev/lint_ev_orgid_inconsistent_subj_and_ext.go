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

/*
 * Contributed by Adriano Santoni <adriano.santoni@staff.aruba.it>
 * of ACTALIS S.p.A. (www.actalis.com).
 */

package cabf_ev

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"

	"errors"
	"regexp"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_ev_orgid_inconsistent_subj_and_ext",
			Description:   "Checks that the organizationIdentifier Subject attribute and the CABFOrganizationIdentifier extension are consistent",
			Citation:      "EVGs 9.2.8 and 9.8.2",
			Source:        lint.CABFEVGuidelines,
			EffectiveDate: util.CABFEV_Sec9_2_8_Date,
		},
		Lint: NewOrgIdInconsistentSubjAndExt,
	})
}

// According to EVGs 9.2.8
type OrganizationIdentifier struct {
	Scheme    string
	Country   string
	State     string
	Reference string
}

// This is according to the EVG (stricter than ETSI EN 319 412-1)
var OrgIdPattern = `^(?P<scheme>[A-Z]{3})(?P<country>[A-Z]{2})(?:\+(?P<state>[A-Z]{2}))?\-(?P<reference>.+)$`

func ParseOrgId(orgIdString string, orgId *OrganizationIdentifier) error {

	re := regexp.MustCompile(OrgIdPattern)

	if !re.MatchString(orgIdString) {
		return errors.New("Cannot parse organizationIdentifier: it is probably invalid")
	}

	names := re.SubexpNames()
	match := re.FindStringSubmatch(orgIdString)

	// Initialize a map to hold group names and values
	result := make(map[string]string)

	// Populate the map
	for i, name := range names {
		if i != 0 && name != "" { // Skip the whole match and unnamed groups
			result[name] = match[i]
		}
	}

	orgId.Scheme = result["scheme"]
	orgId.Country = result["country"]
	orgId.State = result["state"]
	orgId.Reference = result["reference"]

	return nil
}

type orgIdInconsistentSubjAndExt struct{}

func NewOrgIdInconsistentSubjAndExt() lint.LintInterface {
	return &orgIdInconsistentSubjAndExt{}
}

func (l *orgIdInconsistentSubjAndExt) CheckApplies(c *x509.Certificate) bool {
	// It is actually mandatory that, if orgId is present, cabfOrgId be present as well,
	// however this is already checked by another lint
	return util.IsEV(c.PolicyIdentifiers) && (len(c.Subject.OrganizationIDs) > 0) &&
		util.IsExtInCert(c, util.CabfExtensionOrganizationIdentifier)
}

func (l *orgIdInconsistentSubjAndExt) Execute(c *x509.Certificate) *lint.LintResult {
	// It should be safe to assume there is only one element in OrganizationIDs
	var orgId OrganizationIdentifier
	err := ParseOrgId(c.Subject.OrganizationIDs[0], &orgId)
	if err != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "the organizationIdentifier Subject attribute probably has an invalid value"}
	}

	if (c.CABFOrganizationIdentifier.Scheme != orgId.Scheme) ||
		(c.CABFOrganizationIdentifier.Country != orgId.Country) ||
		(c.CABFOrganizationIdentifier.State != orgId.State) ||
		(c.CABFOrganizationIdentifier.Reference != orgId.Reference) {

		return &lint.LintResult{
			Status:  lint.Error,
			Details: "CABFOrganizationIdentifier is NOT consistent with organizationIdentifier"}
	}

	return &lint.LintResult{Status: lint.Pass}
}
