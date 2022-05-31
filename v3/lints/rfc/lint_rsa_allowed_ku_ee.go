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

package rfc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type rsaAllowedKUEe struct{}

/************************************************
RFC 3279: 2.3.1  RSA Keys
  If the keyUsage extension is present in an end entity certificate
   which conveys an RSA public key, any combination of the following
   values MAY be present:

      digitalSignature;
      nonRepudiation;
      keyEncipherment; and
      dataEncipherment.
************************************************/

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_rsa_allowed_ku_ee",
		Description:   "Key usage values digitalSignature, nonRepudiation, keyEncipherment, and dataEncipherment may only be present in an end entity certificate with an RSA key",
		Citation:      "RFC 3279: 2.3.1",
		Source:        lint.RFC3279,
		EffectiveDate: util.RFC3279Date,
		Lint:          NewRsaAllowedKUEe,
	})
}

func NewRsaAllowedKUEe() lint.LintInterface {
	return &rsaAllowedKUEe{}
}

func (l *rsaAllowedKUEe) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.RSA && util.IsExtInCert(c, util.KeyUsageOID) && util.IsSubscriberCert(c)
}

func (l *rsaAllowedKUEe) Execute(c *x509.Certificate) *lint.LintResult {

	//KeyUsageDigitalSignature: allowed
	//KeyUsageContentCommitment: allowed
	//KeyUsageKeyEncipherment: allowed
	//KeyUsageDataEncipherment: allowed
	//KeyUsageKeyAgreement: not allowed
	//KeyUsageCertSign: not allowed
	//KeyUsageCRLSign: not allowed
	//KeyUsageEncipherOnly: not allowed
	//KeyUsageDecipherOnly: not allowed

	var invalidKUs []string

	if c.KeyUsage&x509.KeyUsageKeyAgreement != 0 {
		invalidKUs = append(invalidKUs, "keyAgreement")
	}

	if c.KeyUsage&x509.KeyUsageCertSign != 0 {
		invalidKUs = append(invalidKUs, "keyCertSign")
	}

	if c.KeyUsage&x509.KeyUsageCRLSign != 0 {
		invalidKUs = append(invalidKUs, "cRLSign")
	}

	if c.KeyUsage&x509.KeyUsageEncipherOnly != 0 {
		invalidKUs = append(invalidKUs, "encipherOnly")
	}

	if c.KeyUsage&x509.KeyUsageDecipherOnly != 0 {
		invalidKUs = append(invalidKUs, "decipherOnly")
	}

	if len(invalidKUs) > 0 {
		// Sort the invalid KUs to allow consistent ordering of Details messages for unit testing
		sort.Strings(invalidKUs)
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: fmt.Sprintf("Subscriber certificate with an RSA key contains invalid key usage(s): %s", strings.Join(invalidKUs, ", ")),
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
