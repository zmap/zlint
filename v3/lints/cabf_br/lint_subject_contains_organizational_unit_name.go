package cabf_br

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
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type SubjectContainsOrganizationalUnitName struct{}

/************************************************
BRs: 7.1.4.2.2
Certificate Field: subject:organizationalUnitName (OID: 2.5.4.11)
Required/Optional: Deprecated. Prohibited if the
subject:organizationName is absent or the certificate is issued on or after
September 1, 2022.
This lint check the second requirement, i.e.: Prohibited if the certificate is issued on or after September 1, 2022.
************************************************/

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_subject_contains_organizational_unit_name",
		Description:   "If the certificate is issued on or after September 1, 2022, an organizational unit name MUST NOT be included in subject",
		Citation:      "BRs: 7.1.4.2.2",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABFBRs_1_7_9_Date,
		Lint:          NewSubjectContainsOrganizationalUnitName,
	})
}

func NewSubjectContainsOrganizationalUnitName() lint.LintInterface {
	return &SubjectContainsOrganizationalUnitName{}
}

var subjectOUSunsetDate = time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC)

func (l *SubjectContainsOrganizationalUnitName) CheckApplies(cert *x509.Certificate) bool {
	return !cert.NotBefore.Before(subjectOUSunsetDate)
}

func (l *SubjectContainsOrganizationalUnitName) Execute(cert *x509.Certificate) *lint.LintResult {

	if util.TypeInName(&cert.Subject, util.OrganizationalUnitNameOID) {
		return &lint.LintResult{Status: lint.Error, Details: "subject:organizationalUnitName is prohibited for certificates issued on or after September 1, 2022"}
	}

	return &lint.LintResult{Status: lint.Pass}
}
