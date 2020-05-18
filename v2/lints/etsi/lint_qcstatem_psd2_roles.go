package etsi

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (ส"License"); you may not
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

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemPsd2Roles struct{}

func (l *qcStatemPsd2Roles) Initialize() error {
	return nil
}

func (this *qcStatemPsd2Roles) getStatementOid() *asn1.ObjectIdentifier {
	return &util.IdEtsiPsd2Statem
}

func (l *qcStatemPsd2Roles) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	return util.ParseQcStatem(util.GetExtFromCert(c, util.QcStateOid).Value, *l.getStatementOid()).IsPresent()

}

func isValidPSD2Role(oid *asn1.ObjectIdentifier, name *string) bool {
	return ((oid.Equal(util.IdEtsiPsd2RolePspAs) && *name == "PSP_AS") ||
		(oid.Equal(util.IdEtsiPsd2RolePspPi) && *name == "PSP_PI") ||
		(oid.Equal(util.IdEtsiPsd2RolePspAi) && *name == "PSP_AI") ||
		(oid.Equal(util.IdEtsiPsd2RolePspIc) && *name == "PSP_IC"))
}

func (l *qcStatemPsd2Roles) Execute(c *x509.Certificate) *lint.LintResult {
	errString := ""
	ext := util.GetExtFromCert(c, util.QcStateOid)
	s := util.ParseQcStatem(ext.Value, *l.getStatementOid())
	errString += s.GetErrorInfo()
	if len(errString) == 0 {
		pds := s.(util.EtsiPsd2)
		for _, role := range pds.DecodedPsd2Statm.Roles {
			oid := role.RoleType
			if !isValidPSD2Role(&oid, &role.RoleOfPspName) {
				return &lint.LintResult{Status: lint.Error, Details: "invalid role in PSD2 QcStatement"}
			}
		}
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Error, Details: errString}
	}

}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_psd2_roles",
		Description:   "Verifies that the PSD2 QC Statement extension contains only allowed role OIDs",
		Citation:      "ETSI TS 119 495 V1.2.1, GEN-5.1-2",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2Roles{},
	})
}
