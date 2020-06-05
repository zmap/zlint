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

package etsi

import (
	"encoding/asn1"
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
	"strings"
)

type qcStatemQcPdsValid struct{}

func (l *qcStatemQcPdsValid) Initialize() error {
	return nil
}

func (l *qcStatemQcPdsValid) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	return util.IsQCStatementPresent(c, util.IdEtsiQcsQcEuPDS.String())
}

func isInList(s string, list []string) bool {
	for _, i := range list {
		if i == s {
			return true
		}
	}
	return false
}

func checkPdsTags(c *x509.Certificate, statementOid string) string {
	errString := util.ErrorStringBuilder{Delimiter: "; "}

	pdsQCStatementIndex, err := util.IndexOfValue(c.QCStatements.StatementIDs, statementOid)
	if err != nil {
		errString.Append(fmt.Sprintf("QCStatement %v not present", statementOid))
	}

	ext := util.GetExtFromCert(c, util.QcStateOid)
	rawStatements := x509.QCStatementsASN{}
	if _, err := asn1.Unmarshal(ext.Value, &rawStatements.QCStatements); err != nil {
		return "QCStatement PdsLocations is malformed"
	}

	parsedPdsLocations := make([]asn1.RawValue, 0)
	rawPdsLocations := rawStatements.QCStatements[pdsQCStatementIndex].StatementInfo.FullBytes
	if _, err := asn1.Unmarshal(rawPdsLocations, &parsedPdsLocations); err != nil {
		return "QCStatement PdsLocations is malformed"
	}

	for index, rawPdsLocation := range parsedPdsLocations {
		parsedPdsLocation := make([]asn1.RawValue, 0)
		if _, err := asn1.Unmarshal(rawPdsLocation.FullBytes, &parsedPdsLocation); err != nil {
			return "QCStatement PdsLocations is malformed"
		}
		if len(parsedPdsLocation) != 2 {
			errString.Append(fmt.Sprintf("PdsLocation at index %d is malformed", index))
			continue
		}
		if parsedPdsLocation[0].Tag != asn1.TagIA5String {
			errString.Append(fmt.Sprintf("url attribute of PdsLocation at index %d has an incorrect asn.1 tag", index))
		}
		if parsedPdsLocation[1].Tag != asn1.TagPrintableString {
			errString.Append(fmt.Sprintf("language attribute of PdsLocation at index %d has an incorrect asn.1 tag", index))
		}
	}
	return errString.String()
}

func (l *qcStatemQcPdsValid) Execute(c *x509.Certificate) *lint.LintResult {
	errString := util.ErrorStringBuilder{Delimiter: "; "}

	if len(c.QCStatements.ParsedStatements.PDSLocations) != 1 {
		return &lint.LintResult{Status: lint.Error, Details: "invalid number of PdsLocations objects"}
	}

	//check whether the correct ASN.1 tags were used
	errString.Append(checkPdsTags(c, util.IdEtsiQcsQcEuPDS.String()))

	pds := c.QCStatements.ParsedStatements.PDSLocations[0]
	if len(pds.Locations) == 0 {
		errString.Append("PDS list is empty")
	}

	codeList := make([]string, 0)
	foundEn := false
	for i, loc := range pds.Locations {
		if len(loc.Language) != 2 {
			errString.Append(fmt.Sprintf("PDS location %d has a language code with an invalid length", i))
		}
		if strings.ToLower(loc.Language) == "en" {
			foundEn = true
		}
		if isInList(strings.ToLower(loc.Language), codeList) {
			errString.Append("country code '" + loc.Language + "' appears multiple times")
		}
		codeList = append(codeList, loc.Language)
	}

	if !foundEn {
		errString.Append("no english PDS present")
	}

	if errString.IsEmpty() {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Error, Details: errString.String()}
	}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_qcpds_valid",
		Description:   "Checks that a QC Statement of the type id-etsi-qcs-QcPDS has the correct form",
		Citation:      "ETSI EN 319 412 - 5 V2.2.1 (2017 - 11) / Section 4.3.4",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemQcPdsValid{},
	})
}
