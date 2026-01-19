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

package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
	"github.com/zmap/zlint/v3/util"
)

/*
   SMOKE TEST CASES - file naming convention
   ===================================
   sm 1/0      Certificate for S/MIME: yes/no
   sub 1/0/x   Subscriber certificate: yes/no/don't care
   cp 2/1/0/x  How many CABF S/MIME BR reserved policy OIDs: two, one, zero, don't care
   ef 1/0/x    Certificate issued after Effective Date: one/no/don't care
*/

func TestCABFPolicyMissing(t *testing.T) {

	type Data struct {
		input string
		want  lint.LintStatus
	}

	data := []Data{
		{
			input: "smime/sm0_subx_cpx_efx.pem",
			want:  lint.NA,
		},
		{
			input: "smime/sm1_sub0_cpx_efx.pem",
			want:  lint.NA,
		},
		{
			input: "smime/sm1_sub1_cp0_ef0.pem",
			want:  lint.NE,
		},
		{
			input: "smime/sm1_sub1_cp0_ef1.pem",
			want:  lint.Error,
		},
		{
			input: "smime/sm1_sub1_cp1_ef1.pem",
			want:  lint.Pass,
		},
		{
			input: "smime/sm1_sub1_cp2_ef1.pem",
			want:  lint.Error,
		},
	}

	for _, testData := range data {
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLint("e_exactly_one_smime_policy", testData.input)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}

func TestExactlyOneSMIMEPolicy(t *testing.T) {
	tests := []struct {
		name     string
		policies []asn1.ObjectIdentifier
		want     bool
	}{
		{
			name:     "empty slice returns false",
			policies: []asn1.ObjectIdentifier{},
			want:     false,
		},
		{
			name:     "nil slice returns false",
			policies: nil,
			want:     false,
		},
		{
			name: "single mailbox validated legacy policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRMailboxValidatedLegacyOID,
			},
			want: true,
		},
		{
			name: "single mailbox validated multipurpose policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRMailboxValidatedMultipurposeOID,
			},
			want: true,
		},
		{
			name: "single mailbox validated strict policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRMailboxValidatedStrictOID,
			},
			want: true,
		},
		{
			name: "single organization validated legacy policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBROrganizationValidatedLegacyOID,
			},
			want: true,
		},
		{
			name: "single organization validated multipurpose policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBROrganizationValidatedMultipurposeOID,
			},
			want: true,
		},
		{
			name: "single organization validated strict policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBROrganizationValidatedStrictOID,
			},
			want: true,
		},
		{
			name: "single sponsor validated legacy policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRSponsorValidatedLegacyOID,
			},
			want: true,
		},
		{
			name: "single sponsor validated multipurpose policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRSponsorValidatedMultipurposeOID,
			},
			want: true,
		},
		{
			name: "single sponsor validated strict policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRSponsorValidatedStrictOID,
			},
			want: true,
		},
		{
			name: "single individual validated legacy policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRIndividualValidatedLegacyOID,
			},
			want: true,
		},
		{
			name: "single individual validated multipurpose policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRIndividualValidatedMultipurposeOID,
			},
			want: true,
		},
		{
			name: "single individual validated strict policy returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRIndividualValidatedStrictOID,
			},
			want: true,
		},
		{
			name: "two different SMIME policies returns false",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRMailboxValidatedLegacyOID,
				util.SMIMEBROrganizationValidatedLegacyOID,
			},
			want: false,
		},
		{
			name: "two same SMIME policies returns false",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRMailboxValidatedLegacyOID,
				util.SMIMEBRMailboxValidatedLegacyOID,
			},
			want: false,
		},
		{
			name: "three SMIME policies returns false",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRMailboxValidatedLegacyOID,
				util.SMIMEBROrganizationValidatedMultipurposeOID,
				util.SMIMEBRIndividualValidatedStrictOID,
			},
			want: false,
		},
		{
			name: "one SMIME policy with non-SMIME policies returns true",
			policies: []asn1.ObjectIdentifier{
				util.BRDomainValidatedOID,
				util.SMIMEBRMailboxValidatedLegacyOID,
				util.AnyPolicyOID,
			},
			want: true,
		},
		{
			name: "two SMIME policies with non-SMIME policies returns false",
			policies: []asn1.ObjectIdentifier{
				util.BRDomainValidatedOID,
				util.SMIMEBRMailboxValidatedLegacyOID,
				util.SMIMEBROrganizationValidatedMultipurposeOID,
				util.AnyPolicyOID,
			},
			want: false,
		},
		{
			name: "only non-SMIME policies returns false",
			policies: []asn1.ObjectIdentifier{
				util.BRDomainValidatedOID,
				util.BROrganizationValidatedOID,
				util.BRExtendedValidatedOID,
			},
			want: false,
		},
		{
			name: "SMIME policy at end of list returns true",
			policies: []asn1.ObjectIdentifier{
				util.BRDomainValidatedOID,
				util.BROrganizationValidatedOID,
				util.SMIMEBRIndividualValidatedStrictOID,
			},
			want: true,
		},
		{
			name: "SMIME policy at beginning of list returns true",
			policies: []asn1.ObjectIdentifier{
				util.SMIMEBRSponsorValidatedMultipurposeOID,
				util.BRDomainValidatedOID,
				util.BROrganizationValidatedOID,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.ContainsExactlyOneSMIMEPolicy(tt.policies)
			if got != tt.want {
				t.Errorf("ContainsExactlyOneSMIMEPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}
