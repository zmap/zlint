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

/*
 * This lint contributed by Adriano Santoni <adriano.santoni@staff.aruba.it>
 * of ACTALIS S.p.A. (an ARUBA company). Last revised August 30, 2022.
 *
 * Checks that the Subject in a Server certificate meets the requirement set out
 * in the BRs 7.1.4.2.2, letter i) as modified by ballot SC47 (since BR 1.7.9)
 *
 * 		Certificate Field: subject:organizationalUnitName (OID: 2.5.4.11)
 * 		Required/Optional: Deprecated. Prohibited if the subject:organizationName
 * 		is absent or the certificate is issued on or after September 1, 2022.
 *
 * Since this lint is meant to be deployed after September 1, 2022, it just checks
 * that the Subject field does not contain any OU attributes, regardless of
 * subject:organizationName being absent or not. This lint is only applicable to
 * certificates issued or after September 1, 2022.
 *
 * This lint raises an error in case the MUST NOT above is not met.
 */

package cabf_br

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type serverCertSubjectContainsOU struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name: "e_server_cert_subject_contains_ou",
		Description: "The subject:organizationalUnitName field is prohibited if" +
			" the certificate is issued on or after September 1, 2022.",
		Citation:      "BRs: 7.1.4.2.2",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABFBRs_1_7_9_Date,
		Lint:          NewServerCertSubjectContainsOU,
	})
}

func NewServerCertSubjectContainsOU() lint.LintInterface {
	return &serverCertSubjectContainsOU{}
}

func (l *serverCertSubjectContainsOU) CheckApplies(c *x509.Certificate) bool {

	deprecBeginDate := time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC)

	issueDate := c.NotBefore

	return (issueDate == deprecBeginDate) || issueDate.After(deprecBeginDate)
}

func (l *serverCertSubjectContainsOU) Execute(c *x509.Certificate) *lint.LintResult {

	var subject pkix.RDNSequence

	_, err := asn1.Unmarshal(c.RawSubject, &subject)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal}
	}

	const oidOrgUnit = "2.5.4.11"

	for _, rdn := range subject {
		for _, ava := range rdn {
			if ava.Type.String() == oidOrgUnit {
				return &lint.LintResult{Status: lint.Error}
			}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
