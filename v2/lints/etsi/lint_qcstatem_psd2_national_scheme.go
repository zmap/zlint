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
	"encoding/asn1"
	"regexp"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemPsd2NationalScheme struct{}

func (l *qcStatemPsd2NationalScheme) Initialize() error {
	return nil
}

func (l *qcStatemPsd2NationalScheme) CheckApplies(c *x509.Certificate) bool {
	_, isPresent := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	if !isPresent {
		return false
	}

	orgId := util.GetSubjectOrgId(c.RawSubject)
	re := regexp.MustCompile(`^.{2}:`)
	return re.MatchString(orgId.Value)
}

func (l *qcStatemPsd2NationalScheme) Execute(c *x509.Certificate) *lint.LintResult {

	orgId := util.GetSubjectOrgId(c.RawSubject)
	if !orgId.IsPresent {
		return &lint.LintResult{Status: lint.Error, Details: "missing mandatory subject:OrganizationIdentifier"}
	}
	if orgId.ErrorString != "" {
		return &lint.LintResult{Status: lint.Error, Details: orgId.ErrorString}
	}
	if !util.CheckNationalScheme(orgId.Value) {
		return &lint.LintResult{Status: lint.Error, Details: "invalid format of subject:organizationIdentifier for national scheme"}
	}
	errStr, isPresent := util.IsQcStatemPresent(c, &util.IdQcsPkixQCSyntaxV2)
	if errStr != "" {
		return &lint.LintResult{Status: lint.Error, Details: "error parsing IdQcsPkixQCSyntaxV2 Qc Statement"}
	}

	if !isPresent {
		return &lint.LintResult{Status: lint.Error, Details: "national scheme requires URI in IdQcsPkixQCSyntaxV2 Qc Statement, but this Qc Statement is not present"}
	}
	qcs2Generic := util.ParseQcStatem(util.GetQcStatemExtValue(c), util.IdQcsPkixQCSyntaxV2)
	if qcs2Generic.GetErrorInfo() != "" {
		return &lint.LintResult{Status: lint.Error, Details: qcs2Generic.GetErrorInfo()}
	}
	qcs2 := qcs2Generic.(util.DecodedQcS2)
	for _, x := range qcs2.Decoded.NameRegAuthorities {
		if len(x.FullBytes) < 3 { // have at least tag, length, value one byte each
			continue
		}
		if x.FullBytes[0] != 0x86 {
			continue
		}
		var decodedUri string //
		rest, err := asn1.UnmarshalWithParams(x.FullBytes, &decodedUri, "tag:6")
		if err != nil {
			return &lint.LintResult{Status: lint.Error, Details: err.Error()}
		}
		if len(rest) != 0 {
			return &lint.LintResult{Status: lint.Error, Details: "Trailing bytes after URI"}
		}
		return &lint.LintResult{Status: lint.Pass}

	}

	return &lint.LintResult{Status: lint.Error, Details: "did not find URI element within IdQcsPkixQCSyntaxV2 Qc Statement, which is mandatory for the national scheme format of the subject:organizationIdentifier"}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_psd2_national_scheme",
		Description:   "This lint applies if in a PSD2 certificate (i.e. featuring the PSD2 QcStatement) the subject:organizationIdentifier has a prefix of the form: 2 arbitrary initial characters followed by a colon. In this case it checks that the remainder of the string also fulfills the national scheme syntax.",
		Citation:      "ETSI TS 119 495, '5.2.1 PSD2 Authorization Number or other recognized identifier'",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2NationalScheme{},
	})
}
