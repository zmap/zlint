/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"regexp"
	"unicode/utf8"
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
	if re.MatchString(orgId) {
		return true
	}
	return false

}

func (l *qcStatemPsd2OrgId) Execute(c *x509.Certificate) *LintResult {

	orgId := util.GetSubjectOrgId(c.RawSubject)
	if len(orgId.ErrorString) != 0 {
		return &LintResult{Status: Error, Details: orgId.ErrorString}
	}
	if !orgId.IsPresent {
		return &LintResult{Status: Error, Details: "subject:organizationIdentifier field missing from a certificate for which it is mandatory"}
	}

	_, isPresent := util.IsQcStatemPresent(c, &util.IdQcsPkixQCSyntaxV2)
	if isPresent {
		qcs2Generic := util.ParseQcStatem(util.GetQcStatemExtValue(c), util.IdQcsPkixQCSyntaxV2)
		if qcs2Generic.GetErrorInfo() != "" {
			return &LintResult{Status: Error, Details: qcs2Generic.GetErrorInfo()}
		}
		qc2 := qcs2Generic.(util.DecodedQcS2)
		if qc2.Decoded.SemanticsId.Equal(util.IdEtsiQcsSemanticsIdLegal) {
			// TS 119 412 1: all the 5 prefixes allowed
			if !orgidHasLegalPersonPrefix(orgId.Value) {
				return &LintResult{Status: Error, Details: "found legal person syntax identifier, but subject:organizationIdentifier does not have one of the prefixes allowed in this syntax"}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_qcstatem_psd2_orgid",
		Description:   "For a PSD2 Certificate this lint checks that if the Legal Person Semantics Identfiier is present, the subject:organizationIdentifier field has one of the allowed prefixes",
		Citation:      "ETSI TS 119 495 V1.2.1, Sec. 5.2.1",
		Source:        EtsiTs_119_495_EsiPsd,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2OrgId{},
	})
}
