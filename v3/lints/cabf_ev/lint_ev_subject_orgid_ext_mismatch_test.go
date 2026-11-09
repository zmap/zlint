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

package cabf_ev

import (
	"testing"

	"github.com/zmap/zcrypto/encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubjectOrgIdExtMismatchNilExtension(t *testing.T) {
	c := &x509.Certificate{
		Subject:           pkix.Name{OrganizationIDs: []string{"NTRGB-12345678"}},
		PolicyIdentifiers: []asn1.ObjectIdentifier{{2, 23, 140, 1, 1}},
		Extensions: []pkix.Extension{
			{Id: asn1.ObjectIdentifier{2, 23, 140, 3, 1}, Value: []byte{0x04, 0x03, 0x01, 0x02, 0x03}},
		},
		CABFOrganizationIdentifier: nil,
	}

	l := NewSubjectOrgIdExtMismatch()
	if l.CheckApplies(c) {
		t.Fatal("CheckApplies returned true for a certificate with an undecodable cabfOrganizationIdentifier extension")
	}
}

func TestSubjectOrgIdExtMismatch(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "NTR at country level, consistent",
			InputFilename:  "ev_orgid_ext_match_ntr.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NTR with subdivision, consistent",
			InputFilename:  "ev_orgid_ext_match_subdivision.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Three character alphanumeric subdivision, consistent",
			InputFilename:  "ev_orgid_ext_match_subdivision_alnum.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "PSD, NCA identifier is part of the reference, consistent",
			InputFilename:  "ev_orgid_ext_match_psd.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "VAT, consistent",
			InputFilename:  "ev_orgid_ext_match_vat.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "Registration Scheme identifier differs",
			InputFilename:   "ev_orgid_ext_mismatch_scheme.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "cabfOrganizationIdentifier extension is inconsistent with subject:organizationIdentifier 'NTRGB-12345678': registrationSchemeIdentifier is 'VAT' but the subject attribute has 'NTR'",
		},
		{
			Name:            "Registration country differs",
			InputFilename:   "ev_orgid_ext_mismatch_country.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "cabfOrganizationIdentifier extension is inconsistent with subject:organizationIdentifier 'NTRGB-12345678': registrationCountry is 'DE' but the subject attribute has 'GB'",
		},
		{
			Name:            "Registration state or province differs",
			InputFilename:   "ev_orgid_ext_mismatch_state.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "cabfOrganizationIdentifier extension is inconsistent with subject:organizationIdentifier 'NTRUS+CA-12345678': registrationStateOrProvince is 'NY' but the subject attribute has 'CA'",
		},
		{
			Name:            "Registration reference differs",
			InputFilename:   "ev_orgid_ext_mismatch_reference.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "cabfOrganizationIdentifier extension is inconsistent with subject:organizationIdentifier 'NTRGB-12345678': registrationReference is '87654321' but the subject attribute has '12345678'",
		},
		{
			Name:            "Extension omits the subdivision present in the subject",
			InputFilename:   "ev_orgid_ext_mismatch_missing_state.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "cabfOrganizationIdentifier extension is inconsistent with subject:organizationIdentifier 'NTRUS+CA-12345678': registrationStateOrProvince is '' but the subject attribute has 'CA'",
		},
		{
			Name:            "PSD extension drops the NCA identifier from the reference",
			InputFilename:   "ev_orgid_ext_mismatch_psd_reference.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "cabfOrganizationIdentifier extension is inconsistent with subject:organizationIdentifier 'PSDBE-NBB-1234.567.890': registrationReference is '1234.567.890' but the subject attribute has 'NBB-1234.567.890'",
		},
		{
			Name:            "Every component differs",
			InputFilename:   "ev_orgid_ext_mismatch_multiple.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "cabfOrganizationIdentifier extension is inconsistent with subject:organizationIdentifier 'NTRGB-12345678': registrationSchemeIdentifier is 'VAT' but the subject attribute has 'NTR'; registrationCountry is 'DE' but the subject attribute has 'GB'; registrationStateOrProvince is 'BY' but the subject attribute has ''; registrationReference is '87654321' but the subject attribute has '12345678'",
		},
		{
			Name:           "Not an EV certificate",
			InputFilename:  "ev_orgid_ext_not_ev.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Extension absent",
			InputFilename:  "ev_orgid_ext_absent.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Subject attribute does not conform to 7.1.4.2.8",
			InputFilename:  "ev_orgid_ext_unparseable_subject.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Issued before the effective date",
			InputFilename:  "ev_orgid_ext_not_effective.pem",
			ExpectedResult: lint.NE,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ev_subject_orgid_ext_mismatch", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if tc.ExpectedDetails != result.Details {
				t.Errorf("expected details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
