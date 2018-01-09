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

type subCAEKUNameConstraints struct{}

func (l *subCAEKUNameConstraints) Initialize() error {
	return nil
}

func (l *subCAEKUNameConstraints) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func includesNameConstraints(c *x509.Certificate) bool {
	if len(c.PermittedDNSNames) > 0 || len(c.ExcludedDNSNames) > 0 || len(c.PermittedIPAddresses) > 0 || len(c.ExcludedIPAddresses) > 0 || len(c.PermittedDirectoryNames) > 0 || len(c.ExcludedDirectoryNames) > 0 {
		return true
	} else {
		return false
	}
}

func (l *subCAEKUNameConstraints) Execute(c *x509.Certificate) *LintResult {
	for _, eku := range c.ExtKeyUsage {
		if eku == x509.ExtKeyUsageServerAuth {
			if includesNameConstraints(c) {
				return &LintResult{Status: Pass}
			} else {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: NA}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_eku_name_constraints",
		Description:   "Subordinate CA: If includes id-kp-serverAuth EKU, then it MUST include Name constraints w/ constraints on DNSName, IPAddress, and DirectoryName",
		Citation:      "BRs: 7.1.5",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCAEKUNameConstraints{},
	})
}
