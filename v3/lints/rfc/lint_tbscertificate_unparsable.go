package rfc

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_lint_tbs_certificate_ex",
		Description:   "TBSCertificate could not be parsed",
		Citation:      "RFC 5280: 4.1.2",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          NewTBSCertificateUnparsable,
	})
}

func NewTBSCertificateUnparsable() lint.LintInterface {
	return &TBSCertificateUnparsable{}
}

type TBSCertificateUnparsable struct{}

func (l *TBSCertificateUnparsable) CheckApplies(c *x509.Certificate) bool {
	return false
}

func (l *TBSCertificateUnparsable) Execute(c *x509.Certificate) *lint.LintResult {
	return &lint.LintResult{Status: lint.Pass}
}
