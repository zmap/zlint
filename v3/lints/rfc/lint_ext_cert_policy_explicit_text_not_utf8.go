package rfc

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type explicitTextUtf8 struct{}

/*******************************************************************
https://tools.ietf.org/html/rfc6818#section-3

An explicitText field includes the textual statement directly in
the certificate.  The explicitText field is a string with a
maximum size of 200 characters.  Conforming CAs SHOULD use the
UTF8String encoding for explicitText.  VisibleString or BMPString
are acceptable but less preferred alternatives.  Conforming CAs
MUST NOT encode explicitText as IA5String.  The explicitText string
SHOULD NOT include any control characters (e.g., U+0000 to U+001F
and U+007F to U+009F).  When the UTF8String or BMPString encoding
is used, all character sequences SHOULD be normalized according
to Unicode normalization form C (NFC) [NFC].
*******************************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{LintMetadata: lint.LintMetadata{Name: "w_ext_cert_policy_explicit_text_not_utf8",
		Description:   "Compliant certificates should use the utf8string encoding for explicitText",
		Citation:      "RFC 6818: 3",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC6818Date}, Lint: NewExplicitTextUtf8})

}

func NewExplicitTextUtf8() lint.LintInterface {
	return &explicitTextUtf8{}
}

func (l *explicitTextUtf8) CheckApplies(c *x509.Certificate) bool {
	for _, text := range c.ExplicitTexts {
		if text != nil {
			return true
		}
	}
	return false
}

func (l *explicitTextUtf8) Execute(c *x509.Certificate) *lint.LintResult {
	for _, firstLvl := range c.ExplicitTexts {
		for _, text := range firstLvl {
			if text.Tag != 12 {
				return &lint.LintResult{Status: lint.Warn}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
