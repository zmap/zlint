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
	"fmt"
	"unicode"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemQcPdsLangCase struct{}

func (l *qcStatemQcPdsLangCase) Initialize() error {
	return nil
}

func (l *qcStatemQcPdsLangCase) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	return util.IsQCStatementPresent(c, util.IdEtsiQcsQcEuPDS.String())
}

func isOnlyLowerCaseLetters(s string) bool {
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func (l *qcStatemQcPdsLangCase) Execute(c *x509.Certificate) *lint.LintResult {
	warnString := util.ErrorStringBuilder{Delimiter: "; "}

	if len(c.QCStatements.ParsedStatements.PDSLocations) != 1 {
		return &lint.LintResult{Status: lint.Error, Details: "invalid number of PdsLocations objects"}
	}
	pds := c.QCStatements.ParsedStatements.PDSLocations[0]

	for i, loc := range pds.Locations {
		if !isOnlyLowerCaseLetters(loc.Language) {
			warnString.Append(fmt.Sprintf("PDS location %d has a language code containing invalid letters", i))
		}
	}

	if warnString.IsEmpty() {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Warn, Details: warnString.String()}
	}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_qcstatem_qcpds_lang_case",
		Description:   "Checks that a QC Statement of the type id-etsi-qcs-QcPDS features a language code comprised of only lower case letters",
		Citation:      "ETSI EN 319 412 - 5 V2.2.1 (2017 - 11) / Section 4.3.4",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemQcPdsLangCase{},
	})
}
