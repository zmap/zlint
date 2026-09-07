/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

package community

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"

	"regexp"
)

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "w_suspicious_chars_in_address",
			Description:   "Checks for suspicious characters in Subject locality and/or stateOrProvince",
			Citation:      "Do not know what to say here...",
			Source:        lint.Community,
			EffectiveDate: util.ZeroDate,
		},
		Lint: NewSuspiciousCharsInAddress,
	})
}

type SuspiciousCharsInAddress struct{}

func NewSuspiciousCharsInAddress() lint.LintInterface {
	return &SuspiciousCharsInAddress{}
}

func (l *SuspiciousCharsInAddress) CheckApplies(c *x509.Certificate) bool {
	return len(c.Subject.Locality) > 0 || len(c.Subject.Province) > 0
}

var suspiciousRe = regexp.MustCompile(`[,:="\(\)]`)

func (l *SuspiciousCharsInAddress) Execute(c *x509.Certificate) *lint.LintResult {

	var locSuspicious bool
	var stateSuspicious bool
	var errMsg string

	locSuspicious = len(c.Subject.Locality) > 0 && suspiciousRe.MatchString(c.Subject.Locality[0])

	stateSuspicious = len(c.Subject.Province) > 0 && suspiciousRe.MatchString(c.Subject.Province[0])

	if locSuspicious && stateSuspicious {
		errMsg = "Subject::localityName and Subject::stateOrProvinceName contain suspicious characters"
	} else if locSuspicious {
		errMsg = "Subject::localityName contains suspicious character(s)"
	} else {
		errMsg = "Subject::stateOrProvinceName contains suspicious character(s)."
	}

	if locSuspicious || stateSuspicious {
		return &lint.LintResult{
			Status: lint.Warn, Details: errMsg,
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
