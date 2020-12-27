package rfc

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
RFC 5280: 4.1.2.2.  Serial Number
   The serial number MUST be a positive integer assigned by the CA to each
   certificate. It MUST be unique for each certificate issued by a given CA
   (i.e., the issuer name and serial number identify a unique certificate).
   CAs MUST force the serialNumber to be a non-negative integer.

   Given the uniqueness requirements above, serial numbers can be expected to
   contain long integers.  Certificate users MUST be able to handle serialNumber
   values up to 20 octets.  Conforming CAs MUST NOT use serialNumber values longer
   than 20 octets.

   Note: Non-conforming CAs may issue certificates with serial numbers that are
   negative or zero.  Certificate users SHOULD be prepared togracefully handle
   such certificates.
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type serialNumberTooLong struct{}

func (l *serialNumberTooLong) Initialize() error {
	return nil
}

func (l *serialNumberTooLong) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *serialNumberTooLong) Execute(c *x509.Certificate) *lint.LintResult {
	positive := c.SerialNumber.Sign() != -1
	length := c.SerialNumber.BitLen()
	if positive && length > 159 {
		// https://github.com/zmap/zlint/issues/502#event-4137076637
		//
		// The maximum number of octets is 20 (160 bits), however there is a
		// a complication when the serial number is exactly 160 bits long,
		// and the MSB is 1, wherein implementations must prefix the serial
		// with 0x00 in order to clearly signify the sign, thus putting the
		// encoding past the octet limit.
		//
		// Since big.Int returns the minimum bit length required to represent
		// the number, the MSB is always 1. Thus, if the bit length is 160 or
		// higher then the serial number will overflow our 20 octet limit.
		details := ""
		if length == 160 {
			details = "The certificate's serial number is " +
				"exactly 20 octets long. Once this encodes to DER, implementations " +
				"will prefix it with a 0x00 byte in order to maintain a positive sign, " +
				"thus putting it over the 20 octet limit (after encoding)."
		}
		return &lint.LintResult{Status: lint.Error, Details: details}
	} else if !positive && c.SerialNumber.BitLen() > 160 {
		// Negative numbers are invalid, however it is still worthwhile
		// to apply the lint.
		return &lint.LintResult{Status: lint.Error}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_serial_number_longer_than_20_octets",
		Description:   "Certificates must not have a serial number longer than 20 octets",
		Citation:      "RFC 5280: 4.1.2.2",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC3280Date,
		Lint:          &serialNumberTooLong{},
	})
}
