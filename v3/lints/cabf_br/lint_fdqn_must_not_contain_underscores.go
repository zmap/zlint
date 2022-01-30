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
	"fmt"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name: "e_fqdn_must_not_contain_underscores",
		Description: "Prior to April 1, 2019, certificates containing underscore characters (“_”) in domain labels in dNSName entries MAY be issued as follows: " +
			"• dNSName entries MAY include underscore characters such that replacing all underscore characters with hyphen characters (“-“) would result in a valid domain label, and; " +
			"• Underscore characters MUST NOT be placed in the left most domain label, and; " +
			"* Such certificates MUST NOT be valid for longer than 30 days.",
		Citation:      "BR 7.1.4.2.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABFBRs_1_6_2_Date,
		Lint:          NewDNSNameMustNotIncludeUnderscore,
	})
}

type DNSNameMustdNotIncludeUnderscore struct{}

func NewDNSNameMustNotIncludeUnderscore() lint.LintInterface {
	return &DNSNameMustdNotIncludeUnderscore{}
}

func (l *DNSNameMustdNotIncludeUnderscore) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *DNSNameMustdNotIncludeUnderscore) Execute(c *x509.Certificate) *lint.LintResult {
	for _, dns := range c.DNSNames {
		fqdnPortion := util.RemovePrependedWildcard(dns)
		labels := strings.Split(fqdnPortion, ".")
		for _, label := range labels {
			if strings.Contains(label, "_") {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("dNSName ('%s') MUST NOT contain an underscore character ('_')", dns)}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
