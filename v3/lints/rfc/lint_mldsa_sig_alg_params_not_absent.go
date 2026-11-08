/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

package rfc

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type mldsaSigAlgParamsNotAbsent struct{}

/*******************************************************************************************************
RFC 9881, Section 2:
The contents of the parameters component for each algorithm MUST be absent.

RFC 9881, Section 3:
sa-ml-dsa-44 SIGNATURE-ALGORITHM ::= {
  IDENTIFIER id-ml-dsa-44
  PARAMS ARE absent
  ...
*******************************************************************************************************/

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_mldsa_sig_alg_params_not_absent",
			Description:   "ML-DSA: the parameters of the signature algorithm identifier MUST be absent",
			Citation:      "RFC 9881, Section 2",
			Source:        lint.RFC9881,
			EffectiveDate: util.RFC9881Date,
		},
		Lint: NewMldsaSigAlgParamsNotAbsent,
	})
}

func NewMldsaSigAlgParamsNotAbsent() lint.LintInterface {
	return &mldsaSigAlgParamsNotAbsent{}
}

func (l *mldsaSigAlgParamsNotAbsent) CheckApplies(c *x509.Certificate) bool {
	return util.IsMLDSAOID(c.SignatureAlgorithmOID)
}

func (l *mldsaSigAlgParamsNotAbsent) Execute(c *x509.Certificate) *lint.LintResult {
	tbsSignature, err := util.GetSignatureAlgorithmInTBSEncoded(c)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal, Details: err.Error()}
	}
	hasParameters, err := util.AlgorithmIdentifierHasParameters(tbsSignature)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal, Details: "tbsCertificate.signature: " + err.Error()}
	}
	if hasParameters {
		return &lint.LintResult{Status: lint.Error, Details: "tbsCertificate.signature algorithm identifier has parameters"}
	}

	signatureAlgorithm, err := util.GetSignatureAlgorithmEncoded(c)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal, Details: err.Error()}
	}
	hasParameters, err = util.AlgorithmIdentifierHasParameters(signatureAlgorithm)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal, Details: "signatureAlgorithm: " + err.Error()}
	}
	if hasParameters {
		return &lint.LintResult{Status: lint.Error, Details: "signatureAlgorithm algorithm identifier has parameters"}
	}

	return &lint.LintResult{Status: lint.Pass}
}
