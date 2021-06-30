/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

/*
 * This lint contributed by Adriano Santoni <adriano.santoni@staff.aruba.it>
 * of ACTALIS S.p.A. (an ARUBA company). Last revised June 30, 2021.
 * 
 * Checks that the EKU extension in a Subordinate CA certificate meets the 
 * "MUST NOT" requirement set out in the BRs 7.1.2.2, letter g):
 * 
 *     For Subordinate CA Certificates that will be used to issue 
 *     TLS certificates, the value id-kp-serverAuth[RFC5280] MUST be present. 
 *     The value id-kp-clientAuth[RFC5280] MAY be present. 
 *     The values id-kp-emailProtection[RFC5280], id-kp-codeSigning[RFC5280], 
 *     id-kp-timeStamping[RFC5280], and anyExtendedKeyUsage[RFC5280] MUST NOT 
 *     be present. Other values SHOULD NOT be present. 
 *     For Subordinate CA Certificates that are not used to issue TLS certificates, 
 *     then the value id-kp-serverAuth[RFC5280] MUST NOT be present. 
 *     Other values MAY be present, but SHOULD NOT combine multiple independent usages 
 *     (e.g.including id-kp-timeStamping[RFC5280] with id-kp-codeSigning[RFC5280]).
 */

package cabf_br

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sub_ca_eku_incompatible_values",
		Description:   "Subordinate CA extkeyUsage: if serverAuth is present, then emailProtection, codeSigning, timeStamping, and anyExtendedKeyUsage MUST NOT be present.",
		Citation:      "BRs: 7.1.2.2",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABFBRs_1_7_1_Date,
		Lint:          NewSubCAEkuIncompatibleValues,
	})
}

type subCAEkuIncompatibleValues struct{}

func NewSubCAEkuIncompatibleValues() lint.LintInterface {
    return &subCAEkuIncompatibleValues{}
}

func (l *subCAEkuIncompatibleValues) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func (l *subCAEkuIncompatibleValues) Execute(c *x509.Certificate) *lint.LintResult {
	
	serverAuthPresent := false
	emailProtectionPresent := false
	codeSigningPresent := false
	timeStampingPresent := false
	anyEkuPresent := false
	
	for _, ekuValue := range c.ExtKeyUsage {
		if ekuValue == x509.ExtKeyUsageServerAuth {
			serverAuthPresent = true
		}
		if ekuValue == x509.ExtKeyUsageEmailProtection {
			emailProtectionPresent = true
		}
		if ekuValue == x509.ExtKeyUsageCodeSigning {
			codeSigningPresent = true
		}
		if ekuValue == x509.ExtKeyUsageTimeStamping {
			timeStampingPresent = true
		}
		if ekuValue == x509.ExtKeyUsageAny {
			anyEkuPresent = true
		}
	}
	
	if serverAuthPresent && (emailProtectionPresent || codeSigningPresent || timeStampingPresent || anyEkuPresent) {
		return &lint.LintResult{Status: lint.Error}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}

