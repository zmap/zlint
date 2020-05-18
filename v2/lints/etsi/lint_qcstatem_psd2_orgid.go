package etsi

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

import (
	"regexp"
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemPsd2OrgId struct{}

func (l *qcStatemPsd2OrgId) Initialize() error {
	return nil
}

func (l *qcStatemPsd2OrgId) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	if !util.ParseQcStatem(util.GetExtFromCert(c, util.QcStateOid).Value, util.IdEtsiPsd2Statem).IsPresent() {
		return false
	}
	return true
}

func orgidHasLegalPersonPrefix(orgId string) bool {
	if utf8.RuneCountInString(orgId) < 3 {
		return false
	}
	re := regexp.MustCompile(`^(NTR|VAT|PSD|LEI|(.){2}:)`)
	return re.MatchString(orgId)
}

func (l *qcStatemPsd2OrgId) Execute(c *x509.Certificate) *lint.LintResult {

	orgId := util.GetSubjectOrgId(c.RawSubject)
	if len(orgId.ErrorString) != 0 {
		return &lint.LintResult{Status: lint.Error, Details: orgId.ErrorString}
	}
	if !orgId.IsPresent {
		return &lint.LintResult{Status: lint.Error, Details: "subject:organizationIdentifier field missing from a certificate for which it is mandatory"}
	}

	_, isPresent := util.IsQcStatemPresent(c, &util.IdQcsPkixQCSyntaxV2)
	if isPresent {
		qcs2Generic := util.ParseQcStatem(util.GetQcStatemExtValue(c), util.IdQcsPkixQCSyntaxV2)
		if qcs2Generic.GetErrorInfo() != "" {
			return &lint.LintResult{Status: lint.Error, Details: qcs2Generic.GetErrorInfo()}
		}
		qc2 := qcs2Generic.(util.DecodedQcS2)
		if qc2.Decoded.SemanticsId.Equal(util.IdEtsiQcsSemanticsIdLegal) {
			// TS 119 412 1: all the 5 prefixes allowed
			if !orgidHasLegalPersonPrefix(orgId.Value) {
				return &lint.LintResult{Status: lint.Error, Details: "found legal person syntax identifier, but subject:organizationIdentifier does not have one of the prefixes allowed in this syntax"}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_psd2_orgid",
		Description:   "For a PSD2 Certificate this lint checks that if the Legal Person Semantics Identifier is present, the subject:organizationIdentifier field has one of the allowed prefixes",
		Citation:      "ETSI TS 119 495 V1.2.1, Sec. 5.2.1",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2OrgId{},
	})
}
