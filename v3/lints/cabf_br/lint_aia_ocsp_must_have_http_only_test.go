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

func TestBRAIAOCSPHasHTTPOnly(t *testing.T) {
	testCases := []struct {
		Name          string
		InputFilename string

		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:          "NA - AIA has only one HTTP URI for id-ad-caIssuers accessMethod.",
			InputFilename: "aiaCaIssuersHTTPOnly.pem",

			ExpectedResult: lint.NA,
		},
		{
			Name:          "NA - AIA is missing",
			InputFilename: "subjectOCorrectEncoding.pem",

			ExpectedResult: lint.NA,
		},
		{
			Name:          "pass - AIA has one HTTP URL for id-ad-ocsp accessMethod.",
			InputFilename: "aiaCaIssuersHttpOnlyNoCAIssuers.pem",

			ExpectedResult: lint.Pass,
		},
		{
			Name:          "error - AIA has two HTTP URLs for id-ad-ocsp accessMethod, one is HTTP the other is LDAP.",
			InputFilename: "aiaOCSPOneHTTPOneLDAP.pem",

			ExpectedResult:  lint.Error,
			ExpectedDetails: "Found scheme ldap in OCSP URL of AIA, which is not allowed.",
		},
		{
			Name:          "error - AIA has one HTTPS URL for id-ad-ocsp accessMethod",
			InputFilename: "aiaOCSPWithHTTPSURL.pem",

			ExpectedResult:  lint.Error,
			ExpectedDetails: "Found scheme https in OCSP URL of AIA, which is not allowed.",
		},
		{
			Name:          "NE - AIA has only one HTTP URI for id-ad-ocsp accessMethod and it is issued before September 15th 2023.",
			InputFilename: "aiaOCSPHttpOnlyNE.pem",

			ExpectedResult: lint.NE,
		},
		{
			Name:          "NA - CA certificate issued on September 15th 2023.",
			InputFilename: "caCertificateAfter15092023.pem",

			ExpectedResult: lint.NA,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_aia_ocsp_must_have_http_only", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}

			if tc.ExpectedResult == lint.Error && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %q, was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
