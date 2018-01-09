/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCaMustNotContainAnyPolicy struct{}

func (l *subCaMustNotContainAnyPolicy) Initialize() error {
	return nil
}

func (l *subCaMustNotContainAnyPolicy) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c)
}

func (l *subCaMustNotContainAnyPolicy) Execute(c *x509.Certificate) *LintResult {
	for _, policy := range c.PolicyIdentifiers {
		if policy.Equal(util.AnyPolicyOID) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_must_not_contain_any_policy",
		Description:   "Subordinate CA: MUST NOT contain the anyPolicy identifier (2.5.29.32.0)",
		Citation:      "BRs: 7.1.6.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCaMustNotContainAnyPolicy{},
	})
}
