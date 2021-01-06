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

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type nameConstraintNotFQDN struct{}

func (l *nameConstraintNotFQDN) Initialize() error {
	return nil
}

func (l *nameConstraintNotFQDN) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintNotFQDN) Execute(c *x509.Certificate) *lint.LintResult {
	for _, subtreeString := range c.PermittedURIs {
		var host = subtreeString.Data
		if string(host[0]) == string(".") {
			host = "l" + host
		}
		if !util.IsFQDN(host) {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: fmt.Sprintf("certificate contained a name constraint that wasn't specified as a fully qualified domain name: %v", host),
			}
		}
	}
	for _, subtreeString := range c.ExcludedURIs {
		var host = subtreeString.Data
		if string(host[0]) == string(".") {
			host = "l" + host
		}
		if !util.IsFQDN(host) {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: fmt.Sprintf("certificate contained a name constraint that wasn't specified as a fully qualified domain name: %v", host),
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_name_constraint_not_fqdn",
		Description:   "For URIs, the constraint MUST be specified as a fully qualified domain name [...] When the constraint begins with a period, it MAY be expanded with one or more labels.",
		Citation:      "RFC 5280: 4.2.1.10",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &nameConstraintNotFQDN{},
	})
}
