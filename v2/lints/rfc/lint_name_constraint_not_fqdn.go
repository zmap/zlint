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
	for _, subtreeString := range c.PermittedURIAddresses {
		if subtreeString.Data != "" {
			parsedURI, err := url.Parse(subtreeString.Data)
			if err != nil {
				return &lint.LintResult{Status: lint.Error}
			}
			host := parsedURI.Host
			if !util.IsFQDN(host) {
				return &lint.LintResult{Status: lint.Error}
			}
		}
	}
	for _, subtreeString := range c.ExcludedURIAddresses {
		if subtreeString.Data != "" {
			parsedURI, err := url.Parse(subtreeString.Data)
			if err != nil {
				return &lint.LintResult{Status: lint.Error}
			}
			host := parsedURI.Host
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
		Description:   " For URIs, the constraint MUST be specified as a fully qualified domain name ",
		Citation:      "RFC 5280: 4.2.1.10",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &nameConstraintNotFQDN{},
	})
}
