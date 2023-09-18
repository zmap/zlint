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

package cabf_br

import (
	"regexp"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_fqdn_contains_non_ldh_label",
		Description:   "The Fully‐Qualified Domain Name or the FQDN portion of the Wildcard Domain Name contained in the entry MUST be composed entirely of LDH Labels",
		Citation:      "BR 7.1.4.2.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABFBRs_1_8_0_Date,
		Lint:          NewFQDNContainsNonLDHLabel,
	})
}

var nonLDHCharacterRegex = regexp.MustCompile(`[^a-zA-Z0-9\-]`)

type FQDNContainsNonLDHLabel struct{}

func NewFQDNContainsNonLDHLabel() lint.LintInterface {
	return &FQDNContainsNonLDHLabel{}
}

func (l *FQDNContainsNonLDHLabel) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func IsLDHLabel(label string) bool {
	return len(label) > 0 &&
		len(label) <= 63 &&
		!nonLDHCharacterRegex.MatchString(label) &&
		!strings.HasPrefix(label, "-") &&
		!strings.HasSuffix(label, "-")
}

func (l *FQDNContainsNonLDHLabel) Execute(c *x509.Certificate) *lint.LintResult {
	for _, dns := range c.DNSNames {
		fqdnPortion := util.RemovePrependedWildcard(dns)

		if !util.AllLabelsSatisfyPredicate(fqdnPortion, IsLDHLabel) {
			return &lint.LintResult{Status: lint.Error}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
