package mozilla

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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

/************************************************

https://www.mozilla.org/en-US/about/governance/policies/security-group/certs/policy/

Section 5.1.1 RSA

RSASSA-PSS with SHA-256, MGF-1 with SHA-256, and a salt length of 32 bytes.

The encoded AlgorithmIdentifier MUST match the following hex-encoded bytes:

304106092a864886f70d01010a3034a00f300d0609608648016503040201
0500a11c301a06092a864886f70d010108300d0609608648016503040201
0500a203020120

RSASSA-PSS with SHA-384, MGF-1 with SHA-384, and a salt length of 48 bytes.

The encoded AlgorithmIdentifier MUST match the following hex-encoded bytes:

304106092a864886f70d01010a3034a00f300d0609608648016503040202
0500a11c301a06092a864886f70d010108300d0609608648016503040202
0500a203020130

RSASSA-PSS with SHA-512, MGF-1 with SHA-512, and a salt length of 64 bytes.

The encoded AlgorithmIdentifier MUST match the following hex-encoded bytes:

304106092a864886f70d01010a3034a00f300d0609608648016503040203
0500a11c301a06092a864886f70d010108300d0609608648016503040203
0500a203020140
************************************************/

import (
	"bytes"
	"encoding/hex"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type rsaPssAidEncoding struct{}

var allowedAidEncodingsForPSS = [3]string{
	"304106092a864886f70d01010a3034a00f300d06096086480165030402010500a11c301a06092a864886f70d010108300d06096086480165030402010500a203020120",
	"304106092a864886f70d01010a3034a00f300d06096086480165030402020500a11c301a06092a864886f70d010108300d06096086480165030402020500a203020130",
	"304106092a864886f70d01010a3034a00f300d06096086480165030402030500a11c301a06092a864886f70d010108300d06096086480165030402030500a203020140"}

func (l *rsaPssAidEncoding) Initialize() error {
	return nil
}

func (l *rsaPssAidEncoding) CheckApplies(c *x509.Certificate) bool {
	return c.SignatureAlgorithmOID.Equal(util.OidRSASSAPSS)
}

func (l *rsaPssAidEncoding) Execute(c *x509.Certificate) *lint.LintResult {
	signatureAlgoID, err := util.GetSignatureAlgorithmInTBSEncoded(c)
	if err != nil {
		return &lint.LintResult{Status: lint.Error, Details: "error reading signatureAlgorithm from TBS"}
	}

	for _, encoding := range allowedAidEncodingsForPSS {
		expectedEncoding, _ := hex.DecodeString(encoding)
		if bytes.Equal(signatureAlgoID, expectedEncoding) {
			return &lint.LintResult{Status: lint.Pass}
		}
	}

	return &lint.LintResult{Status: lint.Error, Details: "RSASSA-PSS parameters are not properly encoded"}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_mp_rsassa-pss_parameters_encoding_correct",
		Description:   "The encoded AlgorithmIdentifier for RSASSA-PSS MUST match specific hex-encoded bytes",
		Citation:      "Mozilla Root Store Policy / Section 5.1.1",
		Source:        lint.MozillaRootStorePolicy,
		EffectiveDate: util.MozillaPolicy27Date,
		Lint:          &rsaPssAidEncoding{},
	})
}
