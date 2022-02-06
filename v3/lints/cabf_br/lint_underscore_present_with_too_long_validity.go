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

package cabf_br

import (
	"fmt"
	"strings"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:            "e_underscore_present_with_too_long_validity",
		Description:     "From December 10th 2018 to April 1st 2019 DNSNames may contain underscores if-and-only-if the certificate is valid for less than thirty days.",
		Citation:        "BR 7.1.4.2.1",
		Source:          lint.CABFBaselineRequirements,
		EffectiveDate:   util.CABFBRs_1_6_2_Date,
		IneffectiveDate: util.BALLOT_SC_12_Ineffective,
		Lint:            func() lint.LintInterface { return &UnderscorePresentWithTooLongValidity{} },
	})
}

type UnderscorePresentWithTooLongValidity struct{}

func (l *UnderscorePresentWithTooLongValidity) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *UnderscorePresentWithTooLongValidity) Execute(c *x509.Certificate) *lint.LintResult {
	validity := c.NotAfter.Sub(c.NotBefore)
	if validity <= time.Hour*24*30 {
		// Underscores are permissible if the cert is valid for less than thirty days
		return &lint.LintResult{Status: lint.Pass}
	}
	for _, dns := range c.DNSNames {
		if strings.Contains(dns, "_") {
			return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("The DNSName '%s' contains an "+
				"underscore character which is only permissible if the certiticate is valid for less than 30 days "+
				"(this certificate is valid for %d days)", dns, validity)}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}
