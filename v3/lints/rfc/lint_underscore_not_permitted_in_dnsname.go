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

package rfc

import (
	"fmt"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_underscore_not_permitted_in_dnsname",
		Description:   "DNSName MUST NOT contain underscore characters",
		Citation:      "RFC5280: 4.1.2.6",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          func() lint.LintInterface { return &UnderscoreNotPermittedInDNSName{} },
	})
}

type UnderscoreNotPermittedInDNSName struct{}

func (l *UnderscoreNotPermittedInDNSName) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *UnderscoreNotPermittedInDNSName) Execute(c *x509.Certificate) *lint.LintResult {
	for _, dns := range c.DNSNames {
		if strings.Contains(dns, "_") {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: fmt.Sprintf("The DNS name '%s' contains an underscore (_) character", dns),
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
