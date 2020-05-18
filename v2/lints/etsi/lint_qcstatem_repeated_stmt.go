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
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemRepeatedStmt struct{}

func (l *qcStatemRepeatedStmt) Initialize() error {
	return nil
}

func (l *qcStatemRepeatedStmt) CheckApplies(c *x509.Certificate) bool {
	return util.HasCertAnyEtsiQcStatement(c)
}

type anyContent struct {
	Raw asn1.RawContent
}

type qcStatementWithOptionalInfoField struct {
	Oid asn1.ObjectIdentifier
	Any asn1.RawValue `asn1:"optional"`
}

func isOidInList(oid asn1.ObjectIdentifier, oidList []*asn1.ObjectIdentifier) bool {
	for _, x := range oidList {
		if oid.Equal(*x) {
			return true
		}
	}
	return false
}

func (l *qcStatemRepeatedStmt) Execute(c *x509.Certificate) *lint.LintResult {
	extVal := util.GetQcStatemExtValue(c)
	var foundOidListSlice = make([]*asn1.ObjectIdentifier, 0)
	sl := make([]anyContent, 0)
	rest, err := asn1.Unmarshal(extVal, &sl)
	if err != nil {
		return &lint.LintResult{Status: lint.Error, Details: "error parsing outer SEQ of QcStatement Extension"}
	}
	if len(rest) != 0 {
		return &lint.LintResult{Status: lint.Error, Details: "QcStatements Extension: rest len of outer seq != 0"}

	}
	for _, raw := range sl {
		var statem qcStatementWithOptionalInfoField
		_, err = asn1.Unmarshal(raw.Raw, &statem)
		if err != nil {
			return &lint.LintResult{Status: lint.Error, Details: "error when parsing QcStatements: " + err.Error()}
		}
		oid := statem.Oid
		etsiQcStmtOidListSlice := util.EtsiQcStmtOidList[:]
		if isOidInList(oid, etsiQcStmtOidListSlice) && isOidInList(oid, foundOidListSlice) {
			oidStr := fmt.Sprintf("%v", oid)
			return &lint.LintResult{Status: lint.Error, Details: "ETSI QcStatement with OID " + oidStr + " appears twice in QcStatements"}
		}
		foundOidListSlice = append(foundOidListSlice, &oid)
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_repeated_stmt",
		Description:   "Checks that none of the ETSI QcStatements appear multiple times within the QcStatements.",
		Citation:      "(Implicit requirement)",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemRepeatedStmt{},
	})
}
