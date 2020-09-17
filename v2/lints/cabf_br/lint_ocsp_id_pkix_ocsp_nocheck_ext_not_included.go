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

package cabf_br

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type OCSPIDPKIXOCSPNocheckExtNotIncluded struct{}

func (l *OCSPIDPKIXOCSPNocheckExtNotIncluded) Initialize() error {
	return nil
}

func (l *OCSPIDPKIXOCSPNocheckExtNotIncluded) CheckApplies(c *x509.Certificate) bool {
	if util.HasEKU(c, x509.ExtKeyUsageOcspSigning) {
		return true
	}
	return false
}

func (l *OCSPIDPKIXOCSPNocheckExtNotIncluded) Execute(c *x509.Certificate) *lint.LintResult {
	if !l.CheckApplies(c) {
		// This point should never be reached, because this LINT only applies to certificates with this EKU
		return &lint.LintResult{Status: lint.Fatal}
	}

	if !util.IsExtInCert(c, util.OscpNoCheckOID) {
		if util.IsServerAuthCert(c) {
			// If the certificate is a TLS certificate, it is clear, that the BRGs apply and we have an ERROR
			return &lint.LintResult{Status: lint.Error}
		}

		// If the certificate is not a TLS certificate, the BRGs apply, if one of the sibling certificates or the parent certificate has the Server Auth EKU
		return &lint.LintResult{Status: lint.Warn, Details: "Check the sibling and parent certificate for the Server Auth EKU. If one of them contains the Server Auth EKU, this is an ERROR and not a WARN"}
	}

	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ocsp_id_pkix_ocsp_nocheck_ext_not_included",
		Description:   "OCSP signing Certificate MUST contain an extension of type id-pkixocsp-nocheck, as defined by RFC6960",
		Citation:      "BRs: 4.9.9",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.MozillaPolicy21Date,
		Lint:          &OCSPIDPKIXOCSPNocheckExtNotIncluded{},
	})
}
