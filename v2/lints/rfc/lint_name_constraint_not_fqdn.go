/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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
	"net/url"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type nameConstraintNotFQDN struct{}

func (l *nameConstraintNotFQDN) Initialize() error {
	return nil
}

func (l *nameConstraintNotFQDN) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintNotFQDN) Execute(c *x509.Certificate) *lint.LintResult {
	for _, uri := range c.PermittedURIAddresses {
		if uri != "" {
			parsedUrl, err := url.Parse(uri)
			if err != nil {
				return &lint.LintResult{Status: lint.Error}
			}
			host := parsedUrl.Host
			if !util.IsFQDN(host) {
				return &lint.LintResult{Status: lint.Error}
			}
		}
	}
	for _, uri := range c.ExcludedURIAddresses {
		if uri != "" {
			parsedUrl, err := url.Parse(uri)
			if err != nil {
				return &lint.LintResult{Status: lint.Error}
			}
			host := parsedUrl.Host
			if !util.IsFQDN(host) {
				return &lint.LintResult{Status: lint.Error}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_name_constraint_not_fqdn",
		Description:   "Fill this in...",
		Citation:      "Fill this in...",
		Source:        UnknownLintSource,
		EffectiveDate: "Change this...",
		Lint:          &nameConstraintNotFQDN{},
	})
}
