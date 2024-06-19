package community

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

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type serialNumberLowEntropy struct{}

func (l *serialNumberLowEntropy) Initialize() error {
	return nil
}

func (l *serialNumberLowEntropy) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *serialNumberLowEntropy) Execute(c *x509.Certificate) *lint.LintResult {
	if len(c.SerialNumber.Bytes()) <= 8 {
		return &lint.LintResult{Status: lint.Notice}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "n_serial_number_low_entropy",
			Description:   "Using a serial of length 64 or lower is inadvisable as it is right on the edge of the BRs 7.1 limit of 64 bits minimum of entropy.",
			Citation:      "BRs: 7.1",
			Source:        lint.Community,
			EffectiveDate: util.CABSerialNumberEntropyDate,
		},
		Lint: func() lint.CertificateLintInterface {
			return &serialNumberLowEntropy{}
		},
	})
}
