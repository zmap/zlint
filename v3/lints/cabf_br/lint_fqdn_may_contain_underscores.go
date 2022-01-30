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
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name: "w_fqdn_may_contain_underscores",
		Description: "Prior to April 1, 2019, certificates containing underscore characters (“_”) in domain labels in dNSName entries MAY be issued as follows: " +
			"• dNSName entries MAY include underscore characters such that replacing all underscore characters with hyphen characters (“-“) would result in a valid domain label, and; " +
			"• Underscore characters MUST NOT be placed in the left most domain label, and; " +
			"* Such certificates MUST NOT be valid for longer than 30 days.",
		Citation:        "BR 7.1.4.2.1",
		Source:          lint.CABFBaselineRequirements,
		EffectiveDate:   util.ZeroDate,
		IneffectiveDate: util.CABFBRs_1_6_2_Date,
		Lint:            NewDNSNameMayIncludeUnderscore,
	})
}

type DNSNameShouldNotIncludeUnderscore struct{}

func NewDNSNameMayIncludeUnderscore() lint.LintInterface {
	return &DNSNameShouldNotIncludeUnderscore{}
}

func (l *DNSNameShouldNotIncludeUnderscore) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *DNSNameShouldNotIncludeUnderscore) Execute(c *x509.Certificate) *lint.LintResult {
	validLongerThanThirtyDays := c.NotAfter.Sub(c.NotBefore) > time.Hour*24*30
	for _, dns := range c.DNSNames {
		fqdnPortion := util.RemovePrependedWildcard(dns)
		labels := strings.Split(fqdnPortion, ".")
		if len(labels) > 0 {
			leftMostLabel := labels[0]
			if strings.Contains(leftMostLabel, "_") {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("The left most label of '%s' MUST NOT contain an underscore character ('_')", dns)}
			}
		}
		for _, label := range labels {
			containsUnderscore := strings.Contains(label, "_")
			wouldBeValidIfHyphenInstead := util.IsLDHLabel(strings.ReplaceAll(label, "_", "-"))
			if containsUnderscore && wouldBeValidIfHyphenInstead && !validLongerThanThirtyDays {
				// Fulfills all clauses, so simply warn that this is deprecated.
				return &lint.LintResult{Status: lint.Warn, Details: fmt.Sprintf("%s contains an underscore "+
					"character ('_'). This character MUST NOT appear within the FQDN after April 1, 2019. For more "+
					"information, please see Ballot SC12: Sunset of Underscores in dNSNames "+
					"(https://cabforum.org/2018/11/12/ballot-sc-12-sunset-of-underscores-in-dnsnames/)", dns)}
			} else if containsUnderscore && !wouldBeValidIfHyphenInstead {
				// Fails the first clause about replacing the _ with - which must result in a valid LDH.
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("%s contains an underscore "+
					"character ('_') that, when replaces by a hyphen, does not result in a valid LDH label. This "+
					"character MUST NOT appear within the FQDN after April 1, 2019. For more information, please see "+
					"Ballot SC12: Sunset of Underscores in dNSNames "+
					"(https://cabforum.org/2018/11/12/ballot-sc-12-sunset-of-underscores-in-dnsnames/)", dns)}
			} else if containsUnderscore && wouldBeValidIfHyphenInstead && validLongerThanThirtyDays {
				// Fails the third clause regarding validity dates being longer than 30 days.
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("%s contains an underscore "+
					"character ('_') and the certificate is valid for more than 30 days. This character MUST NOT appear "+
					"within the FQDN after April 1, 2019. For more information, please see Ballot SC12: Sunset of "+
					"Underscores in dNSNames "+
					"(https://cabforum.org/2018/11/12/ballot-sc-12-sunset-of-underscores-in-dnsnames/)", dns)}
			} else {
				// Label passes lint.
				continue
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
