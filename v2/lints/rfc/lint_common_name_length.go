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

package rfc

/********************************************************************
--  specifications of Upper Bounds MUST be regarded as mandatory
--  from Annex B of ITU-T X.411 Reference Definition of MTS Parameter
--  Upper Bounds

ub-common-name INTEGER ::= 64

-- Note - upper bounds on string types, such as TeletexString, are
-- measured in characters.  Excepting PrintableString or IA5String, a
-- significantly greater number of octets will be required to hold
-- such a value.  As a minimum, 16 octets, or twice the specified
-- upper bound, whichever is the larger, should be allowed for
-- TeletexString.  For UTF8String or UniversalString at least four
-- times the upper bound should be allowed.

********************************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type commonNameLengthLimit struct{}

func (l *commonNameLengthLimit) Initialize() error {
	return nil
}

func (l *commonNameLengthLimit) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	if c.Issuer.CommonName != "" {
		return true
	}
	if c.Subject.CommonName != "" {
		return true
	}
	return false
}

func (l *commonNameLengthLimit) Execute(c *x509.Certificate) *lint.LintResult {
	// Add actual lint here
	if c.Issuer.CommonName != "" {
		if len(c.Issuer.CommonName) > 64 {
			return &lint.LintResult{Status: lint.Error}
		}
	}
	if c.Subject.CommonName != "" {
		if len(c.Subject.CommonName) > 64 {
			return &lint.LintResult{Status: lint.Error}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_common_name_length",
		Description:   "the common name upper limit is 64",
		Citation:      "RFC 5280: Appendix A",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &commonNameLengthLimit{},
	})
}
