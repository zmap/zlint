package rfc

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

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type DNSNameLabelLengthTooLong struct{}

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_rfc_dnsname_label_too_long",
			Description:   "DNSName labels MUST be less than or equal to 63 characters",
			Citation:      "RFC 5280: 4.2.1.6, citing RFC 1035",
			Source:        lint.RFC5280,
			EffectiveDate: util.RFC5280Date,
		},
		Lint: NewDNSNameLabelLengthTooLong,
	})
}

func NewDNSNameLabelLengthTooLong() lint.LintInterface {
	return &DNSNameLabelLengthTooLong{}
}

func (l *DNSNameLabelLengthTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func labelLengthTooLong(domain string) bool {
	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if len(label) > 63 {
			return true
		}
	}
	return false
}

func (l *DNSNameLabelLengthTooLong) Execute(c *x509.Certificate) *lint.LintResult {
	for _, dns := range c.DNSNames {
		labelTooLong := labelLengthTooLong(dns)
		if labelTooLong {
			return &lint.LintResult{Status: lint.Error}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
