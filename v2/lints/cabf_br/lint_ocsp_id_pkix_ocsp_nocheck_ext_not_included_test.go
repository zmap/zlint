package cabf_br

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

import (
	"testing"

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestOCSPIDPKIXOCSPNocheckExtNotIncluded(t *testing.T) {
	testCases := []test.TestCase{
		{
			Name:           "SMIME CA Wrong",
			Filename:       "ocspidpkixocspnocheckextnotincluded_SMIME_CA_wrong.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "SMIME CA Correct",
			Filename:       "ocspidpkixocspnocheckextnotincluded_SMIME_CA_correct.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "TLS CA Wrong",
			Filename:       "ocspidpkixocspnocheckextnotincluded_TLS_CA_wrong.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "TLS CA Correct",
			Filename:       "ocspidpkixocspnocheckextnotincluded_TLS_CA_correct.pem",
			ExpectedResult: lint.NA,
		}, {
			Name:           "Delegated Responder",
			Filename:       "ocspidpkixocspnocheckextnotincluded_delegated_responder.pem",
			ExpectedResult: lint.Pass,
		},
	}

	lintName := "e_ocsp_id_pkix_ocsp_nocheck_ext_not_included"

	test.RunTest(lintName, testCases, t)
}
