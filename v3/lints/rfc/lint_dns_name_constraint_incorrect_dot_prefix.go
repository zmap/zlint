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

package rfc

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	description := `DNS name restrictions are expressed as host.example.com. Any DNS
name that can be constructed by simply adding zero or more labels to
the left-hand side of the name satisfies the name constraint. For
example, www.host.example.com would satisfy the constraint but
host1.example.com would not.

As per RFC 1034 section 3.5 (https://tools.ietf.org/html/rfc1034#section-3.5)
a label MUST begin with a letter.

Note that the error of prefixing a DNS name with a "." may often be caused
by confusing a DNS name with a URI (which DOES allow for the "." prefix).`
	lint.RegisterLint(&lint.Lint{
		Name:          "e_dns_name_constraint_incorrect_dot_prefix",
		Description:   description,
		Citation:      "RFC 5280: 4.2.1.10",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &dNSNameConstraintDotPrefix{},
	})
}

type dNSNameConstraintDotPrefix struct{}

func (l *dNSNameConstraintDotPrefix) Initialize() error {
	return nil
}

func (l *dNSNameConstraintDotPrefix) CheckApplies(c *x509.Certificate) bool {
	return len(c.ExcludedDNSNames)+len(c.PermittedDNSNames) != 0
}

func (l *dNSNameConstraintDotPrefix) Execute(c *x509.Certificate) *lint.LintResult {
	var failures []string
	for _, name := range c.PermittedDNSNames {
		if strings.HasPrefix(name.Data, ".") {
			failures = append(failures, name.Data)
		}
	}
	for _, name := range c.ExcludedDNSNames {
		if strings.HasPrefix(name.Data, ".") {
			failures = append(failures, name.Data)
		}
	}
	if len(failures) == 0 {
		return &lint.LintResult{Status: lint.Pass}
	}
	return &lint.LintResult{
		Status:  lint.Error,
		Details: strings.Join(failures, ", "),
	}
}
