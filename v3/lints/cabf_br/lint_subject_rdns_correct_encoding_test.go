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

package cabf_br

import (
	"strings"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubjectRdnsCorrectEncoding(t *testing.T) {
	data := []struct {
		file    string
		want    lint.LintStatus
		details string
	}{
		{
			"subjectDCWrongEncoding.pem",
			lint.Error,
			"Attribute domainComponent in subjectDN has the wrong encoding UTF8String",
		},
		{
			"subjectCWrongEncoding.pem",
			lint.Error,
			"Attribute countryName in subjectDN has the wrong encoding UTF8String",
		},
		{
			"subjectSTWrongEncoding.pem",
			lint.Error,
			"Attribute stateOrProvinceName in subjectDN has the wrong encoding TeletexString",
		},
		{
			"subjectLWrongEncoding.pem",
			lint.Error,
			"Attribute localityName in subjectDN has the wrong encoding IA5String",
		},
		{
			"subjectPostalCodeWrongEncoding.pem",
			lint.Error,
			"Attribute postalCode in subjectDN has the wrong encoding UniversalString",
		},
		{
			"subjectStreetWrongEncoding.pem",
			lint.Error,
			"Attribute streetAddress in subjectDN has the wrong encoding BMPString",
		},
		{
			"subjectOWrongEncoding.pem",
			lint.Error,
			"Attribute organizationName in subjectDN has the wrong encoding TeletexString",
		},
		{
			"subjectSurnameWrongEncoding.pem",
			lint.Error,
			"Attribute surname in subjectDN has the wrong encoding IA5String",
		},
		{
			"subjectGivenNameWrongEncoding.pem",
			lint.Error,
			"Attribute givenName in subjectDN has the wrong encoding BMPString",
		},
		{
			"subjectOUWrongEncoding.pem",
			lint.Error,
			"Attribute organizationalUnitName in subjectDN has the wrong encoding BMPString",
		},
		{
			"subjectCNWrongEncoding.pem",
			lint.Error,
			"Attribute commonName in subjectDN has the wrong encoding UniversalString",
		},
		{
			"subjectBusinessCategoryWrongEncoding.pem",
			lint.Error,
			"Attribute businessCategory in subjectDN has the wrong encoding TeletexString",
		},
		{
			"subjectjurCWrongEncoding.pem",
			lint.Error,
			"Attribute jurisdictionCountry in subjectDN has the wrong encoding BMPString",
		},
		{
			"subjectjurSTWrongEncoding.pem",
			lint.Error,
			"Attribute jurisdictionStateOrProvince in subjectDN has the wrong encoding IA5String",
		},
		{
			"subjectjurLWrongEncoding.pem",
			lint.Error,
			"Attribute jurisdictionLocality in subjectDN has the wrong encoding BMPString",
		},
		{
			"subjectSerialNumberWrongEncoding.pem",
			lint.Error,
			"Attribute serialNumber in subjectDN has the wrong encoding UniversalString",
		},
		{
			"subjectOrganizationIdentifierWrongEncoding.pem",
			lint.Error,
			"Attribute organizationIdentifier in subjectDN has the wrong encoding TeletexString",
		},
		{
			"subjectDCCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectCCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectSTCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectLCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectPostalCodeCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectStreetCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectOCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectSurnameCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectGivenNameCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectOUCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectCNCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectBusinessCategoryCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectjurCCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectjurSTCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectjurLCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectSerialNumberCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectOrganizationIdentifierCorrectEncoding.pem",
			lint.Pass,
			"",
		},
		{
			"subjectValidCountry.pem",
			lint.NE,
			"",
		},
	}
	for _, d := range data {
		file := d.file
		want := d.want
		details := d.details
		t.Run(file, func(t *testing.T) {
			got := test.TestLint("e_subject_rdns_correct_encoding", file)
			if got.Status != want {
				t.Errorf("expected %v got %v", want, got)
			}
			if !strings.Contains(got.Details, details) {
				t.Errorf("expected the returned details to contain '%s' but got %s", details, got.Details)
			}
		})
	}
}
