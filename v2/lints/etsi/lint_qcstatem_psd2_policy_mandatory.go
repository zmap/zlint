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

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemPsd2PolicyMandatory struct{}

func (l *qcStatemPsd2PolicyMandatory) Initialize() error {
	return nil
}

func (l *qcStatemPsd2PolicyMandatory) CheckApplies(c *x509.Certificate) bool {
	isPresent := util.HasCertAnyEtsiQcStatement(c)
	return isPresent
}

func policyAndTypeAreConsistent(c *x509.Certificate, policy asn1.ObjectIdentifier, policyStr string, qcType asn1.ObjectIdentifier, typeStr string, mustHaveQcStmtFour bool) string {

	if util.HasCertPolicy(c, policy) && !util.HasCertEtsiQcType(c, qcType) {
		return "EU Qualified Certificate has policy " + policyStr + " but doesn't have the corresponding QcType '" + typeStr + "' (ETSI EN 319 412-4: \"Policy identifiers included in the certificate policies extension of EU Qualified Certificates shall be consistent with the EU Qualified Certificate Statements\")"
	}
	_, isQscdStmtPresent := util.IsQcStatemPresent(c, &util.IdEtsiQcsQcSSCD)
	if util.HasCertPolicy(c, policy) && mustHaveQcStmtFour && !isQscdStmtPresent {
		return "Policy indicating QSCD is present, but etsi4-qcStatement-4 is missing (ETSI EN 319 412-5: Clause 5, ETSI EN 319 412-5, Clause 4.2.2)"
	}
	if util.HasCertPolicy(c, policy) && !mustHaveQcStmtFour && isQscdStmtPresent {
		return "ETSI EN 319 411-2: GEN-6.6.1-04:  The qcStatement for QSCD (esi4-qcStatement-4) shall not be included in certificates that are not issued according to [QCP-n-qscd] or [QCP-l-qscd] requirements"
	}
	return ""
}

func (l *qcStatemPsd2PolicyMandatory) Execute(c *x509.Certificate) *lint.LintResult {
	errStr := policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpWeb, "QCP-w", util.IdEtsiQcsQctWeb, "Web", false)
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpNatural, "QCP-n", util.IdEtsiQcsQctEsign, "eSign", false))
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpNaturalQscd, "QCP-n-qscd", util.IdEtsiQcsQctEsign, "eSign", true))
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpLegal, "QCP-l", util.IdEtsiQcsQctEseal, "eSeal", false))
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpLegalQscd, "QCP-l-qscd", util.IdEtsiQcsQctEseal, "eSeal", true))
	if errStr != "" {
		return &lint.LintResult{Status: lint.Error, Details: errStr}
	}
	_, psd2Present := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	_, QcTypePresent := util.IsQcStatemPresent(c, &util.IdEtsiQcsQcType)
	if psd2Present && !QcTypePresent {
		return &lint.LintResult{Status: lint.Error, Details: "EU Qualified Certificate has PSD2 QcStatement but not the mandatory QcType Statement (ETSI EN 319 412-5: \"When the certificate is issued in accordance with Annex III or Annex IV of Regulation (EU) No 910/2014 [i.8], this statement shall be present\")."}
	}
	if util.HasCertAnyEtsiQcStatement(c) && len(c.PolicyIdentifiers) == 0 {
		return &lint.LintResult{Status: lint.Error, Details: "EU Qualified Certificate must at least contain one policy identifier (ETSI EN 319 411-2: GEN-6.6.1-05)"}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_qcstatem_psd2_policy_mandatory",
		Description:   "Tests various requirements regarding policy identifiers of EU Qualified Certificates",
		Citation:      "ETSI EN 319 411-2: GEN-6.6.1-05, ETSI EN 319 412-4: Clause 4.3, ETSI EN 319 412-5: Clause 5",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &qcStatemPsd2PolicyMandatory{},
	})
}
