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
	return util.IsDelegatedOCSPResponderCert(c)
}

func (l *OCSPIDPKIXOCSPNocheckExtNotIncluded) Execute(c *x509.Certificate) *lint.LintResult {
	// If the id-pkix-ocsp-nocheck extension, as specified in RFC 6960, Section 4.2.2.2.1, is present, then
	// the certificate complies.
	if util.IsExtInCert(c, util.OscpNoCheckOID) {
		return &lint.LintResult{Status: lint.Pass}
	}

	// If the certificate is a TLS certificate, the Baseline Requirements apply, which require the presence of
	// id-pkix-ocsp-nocheck as an extension.
	if util.IsServerAuthCert(c) {
		return &lint.LintResult{Status: lint.Error}
	}

	// If the certificate is not unambiguously a TLS certificate, then whether or not the OCSP responder is
	// in scope of the Baseline Requirements depends on whether the issuer of this certificate is capable of
	// being used for TLS. At present, this requires manual attention, so only return a warning.
	return &lint.LintResult{
		Status: lint.Warn,
		Details: "OCSP signing Certificate without id-pkix-ocsp-nocheck. If the associated CA certificate is" +
			" subject to the Baseline Requirements, this is an Error.",
	}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name: "e_ocsp_id_pkix_ocsp_nocheck_ext_not_included",
		Description: "OCSP signing Certificate MUST contain an extension of type id-pkixocsp-nocheck, as" +
			" defined by RFC6960",
		Citation:      "BRs: 4.9.9",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &OCSPIDPKIXOCSPNocheckExtNotIncluded{},
	})
}
