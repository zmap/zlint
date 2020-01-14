/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

/********************************************************************
Section 5.3 - Intermediate Certificates
Intermediate certificates created after January 1, 2019, with the exception
of cross-certificates that share a private key with a corresponding root
certificate: MUST contain an EKU extension; and, MUST NOT include the
anyExtendedKeyUsage KeyPurposeId; and, * MUST NOT include both the
id-kp-serverAuth and id-kp-emailProtection KeyPurposeIds in the same
certificate.
Note that the lint cannot distinguish cross-certificates from other
intermediates.
********************************************************************/

package lints

import (
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type allowedEKU struct{}

func (l *allowedEKU) Initialize() error {
	return nil
}

func (l *allowedEKU) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && !util.IsInMozillaRootStore(c)
}

func (l *allowedEKU) Execute(c *x509.Certificate) *LintResult {
	if len(c.ExtKeyUsage) == 0 || util.HasEKU(c, x509.ExtKeyUsageAny) ||
		(util.HasEKU(c, x509.ExtKeyUsageEmailProtection) && util.HasEKU(c, x509.ExtKeyUsageServerAuth)) {
		return &LintResult{Status: Error}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_mp_allowed_eku",
		Description:   "Separation of id-kp-serverAuth and id-kp-emailProtection KeyPurposeIds",
		Citation:      "Mozilla Root Store Policy / Section 5.3",
		Source:        MozillaRootStorePolicy,
		EffectiveDate: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Lint:          &allowedEKU{},
	})
}
