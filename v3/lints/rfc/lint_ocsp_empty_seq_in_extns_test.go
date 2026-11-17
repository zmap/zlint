/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestOCSPEmptySequenceInExtensions(t *testing.T) {
	testCases := []struct {
		desc string
		file string
		want lint.LintStatus
	}{
		{
			"ResponseExtensions and SingleExtensions both absent",
			"ocsp_resp_rxabs_sxabs.der",
			lint.Pass,
		},
		{
			"ResponseExtensions absent; SingleExtensions present and valid",
			"ocsp_resp_rxabs_sxval.der",
			lint.Pass,
		},
		{
			"ResponseExtensions any; SingleExtensions wrong (empty SEQUENCE)",
			"ocsp_resp_rxany_sxemp.der",
			lint.Error,
		},
		{
			"ResponseExtensions present and valid; SingleExtensions absent",
			"ocsp_resp_rxval_sxabs.der",
			lint.Pass,
		},
		{
			"ResponseExtensions and SingleExtensions both present and valid",
			"ocsp_resp_rxval_sxval.der",
			lint.Pass,
		},
		{
			"ResponseExtensions wrong (empty SEQUENCE); SingleExtensions any",
			"ocsp_resp_rxemp_sxany.der",
			lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.file, func(t *testing.T) {
			got := test.TestOCSPResponseLint(t, "e_ocsp_empty_seq_in_extns", tc.file)
			if got.Status != tc.want {
				t.Errorf("expected %s, got %s: %s", tc.want, got.Status, got.Details)
			}
		})
	}
}
