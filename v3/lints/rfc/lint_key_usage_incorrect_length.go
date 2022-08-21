package rfc

/*
 * ZLint Copyright 2022 Regents of the University of Michigan
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

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type keyUsageIncorrectLength struct{}

// "When present, conforming CAs SHOULD mark this extension as critical."

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_key_usage_incorrect_length",
		Description:   "The key usage is a bit string with exactly eight bits",
		Citation:      "RFC 5280: 4.2.1.3",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          NewKeyUsageIncorrectLength,
	})
}

func NewKeyUsageIncorrectLength() lint.LintInterface {
	return &keyUsageIncorrectLength{}
}

func (l *keyUsageIncorrectLength) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *keyUsageIncorrectLength) Execute(c *x509.Certificate) *lint.LintResult {
	keyUsage := util.GetExtFromCert(c, util.KeyUsageOID).Value
	// Tag: keyUsage[0]
	// Length: keyUsage[1]
	// Unused: keyUsage[2]
	// The actual key usage...
	content := keyUsage[3:]
	if len(content) != 1 {
		return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("key usage is %d bytes long (should be exactly one)", len(content))}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}
