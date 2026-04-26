package cabf_br

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCaIsCa(t *testing.T) {
	tests := []struct {
		id        string
		inputFile string
		expected  lint.LintStatus
	}{
		{"TestKeyCertSignNotCA", "keyCertSignNotCA.pem", lint.Error},
		{"TestKeyCertSignNotCAExplicitFalse", "keyCertSignNotCAExplicitly.pem", lint.Error},
		{"TestKeyCertSignCA", "keyCertSignCA.pem", lint.Pass},
		{"TestCaNoKeyUsageNotApplicable", "subCaNokeyUsage.pem", lint.NA},
		{"TestCaNoCertSignNotApplicable", "subCaNoCertSign.pem", lint.NA},
		{"TestKeyCertSignNoBasicConstraints", "subCertNoBcWithCertSign.pem", lint.NA},
		{"TestKeyCertSignNotCAButRequirementNotEffective", "subCaNotCaButOld.pem", lint.NE},
	}

	for _, testCase := range tests {
		t.Run(testCase.id, func(t *testing.T) {
			var out *lint.LintResult = test.TestLint("e_ca_is_ca", testCase.inputFile)
			if out.Status != testCase.expected {
				t.Errorf("%s: expected %s, got %s", testCase.inputFile, testCase.expected, out.Status)
			}
		})
	}
}
