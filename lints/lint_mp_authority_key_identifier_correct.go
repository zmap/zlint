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
Section 5.2 - Forbidden and Required Practices
CAs MUST NOT issue certificates that have:
- incorrect extensions (e.g., SSL certificates that exclude SSL usage, or authority key IDs
  that include both the key ID and the issuer’s issuer name and serial number); or
********************************************************************/

package lints

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type keyIdentifier struct {
	KeyIdentifier             asn1.RawValue `asn1:"optional,tag:0"`
	AuthorityCertIssuer       asn1.RawValue `asn1:"optional,tag:1"`
	AuthorityCertSerialNumber asn1.RawValue `asn1:"optional,tag:2"`
}

type authorityKeyIdentifierCorrect struct{}

func (l *authorityKeyIdentifierCorrect) Initialize() error {
	return nil
}

func (l *authorityKeyIdentifierCorrect) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.AuthkeyOID) {
		return false
	}
	return true
}

func (l *authorityKeyIdentifierCorrect) Execute(c *x509.Certificate) *LintResult {
	var keyID keyIdentifier

	ext := util.GetExtFromCert(c, util.AuthkeyOID)

	if ext == nil {
		return &LintResult{Status: Fatal}
	}
	if _, err := asn1.Unmarshal(ext.Value, &keyID); err != nil {
		return &LintResult{Status: Fatal}
	}

	hasKeyID := len(keyID.KeyIdentifier.Bytes) > 0
	hasCertIssuer := len(keyID.AuthorityCertIssuer.Bytes) > 0
	if hasKeyID && hasCertIssuer {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_mp_authority_key_identifier_correct",
		Description:   "CAs MUST NOT issue certificates that have authority key IDs that include both the key ID and the issuer's issuer name and serial number",
		Citation:      "Mozilla Root Store Policy / Section 5.2",
		Source:        MozillaRootStorePolicy,
		EffectiveDate: util.MozillaPolicy22Date,
		Lint:          &authorityKeyIdentifierCorrect{},
	})
}
