package cabf_smime

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

import (
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type subCertValidTimeLongerThan825Days struct{}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sub_cert_valid_time_longer_than_825_days",
		Description:   "Validity period: confirming that initial version will seek a maximum validity SHOULD of 27 months",
		Citation:      "BRs: 6.3.2",
		Source:        lint.CABFSMIMEBaselineRequirements,
		EffectiveDate: min(time.Now(), util.CABSMIMEEffectiveDate), // work around to enable this lint with INFO before the Effective Date of the SMIME BRGs
		Lint:          &subCertValidTimeLongerThan825Days{},
	})
}

func (l *subCertValidTimeLongerThan825Days) Initialize() error {
	return nil
}

func (l *subCertValidTimeLongerThan825Days) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertValidTimeLongerThan825Days) Execute(c *x509.Certificate) *lint.LintResult {
	if c.NotBefore.AddDate(0, 0, 825).Before(c.NotAfter) {
		return &lint.LintResult{Status: lint.Error}
	}
	return &lint.LintResult{Status: lint.Pass}
}
