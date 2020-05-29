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

type consistentPolicyAndQCType struct{}

func (l *consistentPolicyAndQCType) Initialize() error {
	return nil
}

func (l *consistentPolicyAndQCType) CheckApplies(c *x509.Certificate) bool {
	// Checks if the QcStatement id-etsi-qcs-QcCompliance is present. If yes, then this lint is triggered.
	_, isEUQualified := util.IsQcStatementPresent(c, &util.IdEtsiQcsQcCompliance)
	return isEUQualified
}

func policyAndTypeAreConsistent(c *x509.Certificate, policy asn1.ObjectIdentifier, policyStr string, qcType asn1.ObjectIdentifier, typeStr string) string {

	// Check #1: Policy in question is present.
	// If the policy is missing, not enough information exists to determine if there is an inconsistency.
	if !util.HasCertPolicy(c, policy) {
		return ""
	}

	// Check #2: Consistency of QCStatement QcType (Id-etsi-qcs-QcType) and policy.
	_, isQcTypePresent := util.IsQcStatementPresent(c, &util.IdEtsiQcsQcType)

	if isQcTypePresent && !util.HasCertEtsiQcType(c, qcType) {
		return fmt.Sprintf("EU Qualified Certificate has policy %s but does not have the corresponding QcType '%s' (ETSI EN 319 412-4: \"Policy identifiers included in the certificate policies extension of EU Qualified Certificates shall be consistent with the EU Qualified Certificate Statements\")", policyStr, typeStr)
	}

	// Check #3: Consistency of QCStatement QcSSCD (id-etsi-qcs-QcSSCD) and policy (bidirectional).
	_, isQscdStatementPresent := util.IsQcStatementPresent(c, &util.IdEtsiQcsQcSSCD)

	qscdPolicyIsPresent := util.HasCertPolicy(c, util.IdEtsiPolicyQcpNaturalQscd) || util.HasCertPolicy(c, util.IdEtsiPolicyQcpLegalQscd)

	if !qscdPolicyIsPresent && isQscdStatementPresent {
		return "ETSI EN 319 411-2: GEN-6.6.1-04: The qcStatement for QSCD (esi4-qcStatement-4) shall not be included in certificates that are not issued according to [QCP-n-qscd] or [QCP-l-qscd] requirements"
	}
	if qscdPolicyIsPresent && !isQscdStatementPresent {
		return "Policy indicating QSCD is present, but etsi4-qcStatement-4 is missing (ETSI EN 319 412-5: Clause 5, ETSI EN 319 412-5, Clause 4.2.2)"
	}
	return ""
}

func (l *consistentPolicyAndQCType) Execute(c *x509.Certificate) *lint.LintResult {

	// Check if an EU qualified certificate contains at least one policy identifier
	if len(c.PolicyIdentifiers) == 0 {
		return &lint.LintResult{Status: lint.Error, Details: "The certificate shall include at least one of the following policy identifier [CHOICE] (ETSI EN 319 411-2: GEN-6.6.1-05)."}
	}

	// Attempt to determine the type based on the certificate policy, based on ETSI EN 319
	// 411-2, GEN-6.6.1-05 
	// CAs are allowed to use either the ETSI-defined policy OIDs or CA-specific OIDs,
	// similar to EV certificates. If using a CA-specific OID, it's not possible to
	// determine the claimed certificate type without maintaining a mapping of every
	// policy OID to ETSI type, similar to what is done for EV policies. Because that is
	// not implemented, this only checks to see if the CA is using one of the standard
	// ETSI policy OIDs as described in GEN-6.6.1-05. If not, the check is skipped
	// because not enough information is present.
	errStr := policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpWeb, "QCP-w", util.IdEtsiQcsQctWeb, "Web")
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpNatural, "QCP-n", util.IdEtsiQcsQctEsign, "eSign"))
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpNaturalQscd, "QCP-n-qscd", util.IdEtsiQcsQctEsign, "eSign"))
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpLegal, "QCP-l", util.IdEtsiQcsQctEseal, "eSeal"))
	util.AppendToStringSemicolonDelim(&errStr, policyAndTypeAreConsistent(c, util.IdEtsiPolicyQcpLegalQscd, "QCP-l-qscd", util.IdEtsiQcsQctEseal, "eSeal"))
	if errStr != "" {
		return &lint.LintResult{Status: lint.Error, Details: errStr}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_policy_and_qctype_consistent",
		Description:   "Checks if the policy specified in an EU qualified certificate is consistent with the QcType",
		Citation:      "ETSI EN 319 411-2: GEN-6.6.1-05, ETSI EN 319 412-4: Clause 4.3, ETSI EN 319 412-5: Clause 5",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiPSD2Date,
		Lint:          &consistentPolicyAndQCType{},
	})
}
