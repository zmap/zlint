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

type crlMldsaSigAlgParamsNotAbsent struct{}

/*******************************************************************************************************
RFC 9881, Section 2:
The contents of the parameters component for each algorithm MUST be absent.

RFC 9881, Section 3:
Signatures are also used in the CRL list ASN.1; the representation below is equivalent to that in
[RFC5280]. In an X.509 CRL, a signature is encoded with an algorithm identifier in the
signatureAlgorithm attribute and a signatureValue attribute that contains the actual signature.
*******************************************************************************************************/

func init() {
	lint.RegisterRevocationListLint(&lint.RevocationListLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_crl_mldsa_sig_alg_params_not_absent",
			Description:   "ML-DSA: the parameters of the CRL signature algorithm identifier MUST be absent",
			Citation:      "RFC 9881, Section 2",
			Source:        lint.RFC9881,
			EffectiveDate: util.RFC9881Date,
		},
		Lint: NewCrlMldsaSigAlgParamsNotAbsent,
	})
}

func NewCrlMldsaSigAlgParamsNotAbsent() lint.RevocationListLintInterface {
	return &crlMldsaSigAlgParamsNotAbsent{}
}

func (l *crlMldsaSigAlgParamsNotAbsent) CheckApplies(r *x509.RevocationList) bool {
	signatureAlgorithm, err := util.GetSignatureAlgorithmInCertListEncoded(r)
	if err != nil {
		return false
	}
	oid, err := util.GetAlgorithmIdentifierOID(signatureAlgorithm)
	if err != nil {
		return false
	}
	return util.IsMLDSAOID(oid)
}

func (l *crlMldsaSigAlgParamsNotAbsent) Execute(r *x509.RevocationList) *lint.LintResult {
	tbsSignature, err := util.GetSignatureAlgorithmInTBSCertListEncoded(r)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal, Details: err.Error()}
	}
	hasParameters, err := util.AlgorithmIdentifierHasParameters(tbsSignature)
	if err != nil {
		return &lint.LintResult{Status: lint.Fatal, Details: "tbsCertList.signature: " + err.Error()}
	}
	if hasParameters {
		return &lint.LintResult{Status: lint.Error, Details: "tbsCertList.signature algorithm identifier has parameters"}
	}

	signatureAlgorithm, err := util.GetSignatureAlgorithmInCertListEncoded(r)
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
