package mozilla

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

/************************************************
https://www.mozilla.org/en-US/about/governance/policies/security-group/certs/policy/

Section 5.1.1 RSA

CAs MUST NOT use the id-RSASSA-PSS OID (1.2.840.113549.1.1.10) within a SubjectPublicKeyInfo to represent a RSA key.
************************************************/

import (
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/util"
)

type rsaPssInSPKI struct{}

func (l *pssInSPKI) Initialize() error {
	return nil
}

func (l *pssInSPKI) CheckApplies(c *x509.Certificate) bool {
	// always check, no certificate is allowed to contain the PSS OID in public key
	return true
}

func (l *pssInSPKI) Execute(c *x509.Certificate) *lint.LintResult {
	publicKeyOID, err := util.GetPublicKeyOID(c)

	if err != nil {
		return &lint.LintResult{Status: lint.Error, Details: "error reading public key OID"}
	}

	if publicKeyOID.Equal(util.OidRSASSAPSS) {
		return &lint.LintResult{Status: lint.Error, Details: "id-RSASSA-PSS OID found in public key"}
	}

	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_mp_pss_in_spki",
		Description:   "CAs MUST NOT use the id-RSASSA-PSS OID (1.2.840.113549.1.1.10) within a SubjectPublicKeyInfo to represent a RSA key.",
		Citation:      "Mozilla Root Store Policy / Section 5.1.1",
		Source:        lint.MozillaRootStorePolicy,
		EffectiveDate: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Lint:          &pssInSPKI{},
	})
}
