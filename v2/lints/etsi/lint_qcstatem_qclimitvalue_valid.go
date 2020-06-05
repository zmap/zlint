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
	"unicode"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemQcLimitValueValid struct{}

func (this *qcStatemQcLimitValueValid) getStatementOid() *asn1.ObjectIdentifier {
	return &util.IdEtsiQcsQcLimitValue
}

func (l *qcStatemQcLimitValueValid) Initialize() error {
	return nil
}

func (l *qcStatemQcLimitValueValid) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	return util.IsQCStatementPresent(c, util.IdEtsiQcsQcLimitValue.String())
}

func isOnlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (l *qcStatemQcLimitValueValid) Execute(c *x509.Certificate) *lint.LintResult {
	errString := util.ErrorStringBuilder{Delimiter: "; "}

	if len(c.QCStatements.ParsedStatements.Limit) != 1 {
		return &lint.LintResult{Status: lint.Error, Details: "invalid number of MonetaryValue objects"}
	}
	qcLv := c.QCStatements.ParsedStatements.Limit[0]

	if qcLv.Amount < 0 {
		errString.Append("amount is negative")
	}

	// Check whether the alphabetic currency code was set
	if len(qcLv.Currency) > 0 {
		// Check whether alphabetic code is set correctly
		if len(qcLv.Currency) != 3 {
			errString.Append("invalid string length of currency code")
		}
		if !isOnlyLetters(qcLv.Currency) {
			errString.Append("currency code string contains not only letters")
		}
	} else { // if the alphabetic code is not set, check numeric code
		if qcLv.CurrencyNumber < 1 || qcLv.CurrencyNumber > 999 {
			errString.Append("numeric currency code is out of range")
		}
	}

	if errString.IsEmpty() {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Error, Details: errString.String()}
	}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_qclimitvalue_valid",
		Description:   "Checks that a QC Statement of the type id-etsi-qcs-QcLimitValue has the correct form",
		Citation:      "ETSI EN 319 412 - 5 V2.2.1 (2017 - 11) / Section 4.3.2",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemQcLimitValueValid{},
	})
}
