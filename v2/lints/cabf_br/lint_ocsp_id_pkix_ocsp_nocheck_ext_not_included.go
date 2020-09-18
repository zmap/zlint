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
	return util.HasEKU(c, x509.ExtKeyUsageOcspSigning)
}

func (l *OCSPIDPKIXOCSPNocheckExtNotIncluded) Execute(c *x509.Certificate) *lint.LintResult {
	// The OCSPNoCheckOID is included in the certficate this is a clear pass (at least for this lint)
	if util.IsExtInCert(c, util.OscpNoCheckOID) {
		return &lint.LintResult{Status: lint.Pass}
	}

	if util.IsServerAuthCert(c) {
		// If the certificate is a TLS certificate, it is clear, that the BRGs apply and we have an ERROR
		return &lint.LintResult{Status: lint.Error}
	}

	// If the certificate is not a TLS certificate, the BRGs only apply, if the parent certificate has the Server Auth EKU or could possibly issue a certificate that has the Server EKU
	return &lint.LintResult{Status: lint.Warn, Details: "If the parent (issuing) certificate contains the Server Auth EKU or no EKU at all this is an ERROR and not a WARN"}
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
