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

package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCannotBeEmptyExtension(t *testing.T) {

	testCases := []struct {
		desc string
		path string
		want lint.LintStatus
	}{
		{
			desc: "Certificate with an assortment of valid extensions",
			path: "empty_seq_of_xxx.pem",
			want: lint.Pass,
		},
		{
			desc: "Certificate with empty AuthorityInformationAccess",
			path: "empty_seq_of_aia.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty CRLDistributionPoints",
			path: "empty_seq_of_aia.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty CertificatePolicies",
			path: "empty_seq_of_cps.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty ExtendedKeyUsage",
			path: "empty_seq_of_eku.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty FreshestCRL",
			path: "empty_seq_of_fre.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty IssuerAlternativeNames",
			path: "empty_seq_of_ian.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty PolicyMappings",
			path: "empty_seq_of_pms.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty SubjectAlternativeNames",
			path: "empty_seq_of_san.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty SubjectDirectoryAttributes",
			path: "empty_seq_of_sda.pem",
			want: lint.Error,
		},
		{
			desc: "Certificate with empty SubjectInformationAccess",
			path: "empty_seq_of_sia.pem",
			want: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := test.TestLint("e_ext_cannot_be_empty_sequence", tc.path)
			if out.Status != tc.want {
				t.Errorf("expected status %s for %s, got %s", tc.want, tc.path, out.Status)
			}
		})
	}
}
