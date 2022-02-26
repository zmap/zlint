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
	"fmt"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

/*
RFC 5280
4.2.1.6.  Subject Alternative Name
...
When the subjectAltName extension contains a domain name system
label, the domain name MUST be stored in the dNSName (an IA5String).
The name MUST be in the "preferred name syntax", as specified by
Section 3.5 of [RFC1034] and as modified by Section 2.1 of
[RFC1123].
...
RFC 1034
3.5. Preferred name syntax
...
<domain> ::= <subdomain> | " "
<subdomain> ::= <label> | <subdomain> "." <label>
<label> ::= <letter> [ [ <ldh-str> ] <let-dig> ]
<ldh-str> ::= <let-dig-hyp> | <let-dig-hyp> <ldh-str>
<let-dig-hyp> ::= <let-dig> | "-"
<let-dig> ::= <letter> | <digit>
<letter> ::= any one of the 52 alphabetic characters A through Z in upper case and a through z in lower case
<digit> ::= any one of the ten digits 0 through 9
...
RFC 1123
2.1  Host Names and Numbers
... One aspect of host name syntax is hereby changed: the
restriction on the first character is relaxed to allow either a
letter or a digit. ...
*/

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_underscore_not_permitted_in_dnsname",
		Description:   "DNSName MUST NOT contain underscore characters",
		Citation:      "RFC5280: 4.1.2.6",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          func() lint.LintInterface { return &UnderscoreNotPermittedInDNSName{} },
	})
}

type UnderscoreNotPermittedInDNSName struct{}

func (l *UnderscoreNotPermittedInDNSName) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *UnderscoreNotPermittedInDNSName) Execute(c *x509.Certificate) *lint.LintResult {
	for _, dns := range c.DNSNames {
		if strings.Contains(dns, "_") {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: fmt.Sprintf("The DNS name '%s' contains an underscore (_) character", dns),
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
