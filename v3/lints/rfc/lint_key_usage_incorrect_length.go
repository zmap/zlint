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
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type keyUsageIncorrectLength struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_key_usage_incorrect_length",
		Description:   "The key usage is a bit string with exactly nine possible flags",
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
	// Any combination of the nine bit flags is legal from perspective
	// of this lint (although requirements elsewhere may limit the combinations).
	//
	// As such, any value greater than 512 (2**9) is out of range of the possible
	// values for a key usage bit string.
	if big.NewInt(0).SetBytes(content).Int64() > 0b111111111 {
		return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("the key usage (%v) contains a value that is out of bounds of the range of possible KU values.", content)}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}
