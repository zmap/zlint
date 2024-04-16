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

func TestSubExtKeyUsageCheck(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - serverAuth EKU",
			InputFilename:  "subExtKeyUsageServerAuth.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - serverAuth and clientAuth EKU",
			InputFilename:  "subExtKeyUSageServerAndClientAuth.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "error - only clientAuth EKU with CA/B TLS BR policy OID",
			InputFilename:  "subExtKeyUsageClientAuth.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "ne - only clientAuth EKU, with CA/B TLS BR policy OID, NotBefore before effective date",
			InputFilename:  "subExtKeyUsageClientAuthPreBRv2.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "error - serverAuth and timeStamping EKU",
			InputFilename:  "subExtKeyUsageServerAuthAndTimeStamping.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "error - serverAuth and unknown PreCertSigCert EKU",
			InputFilename:  "subExtKeyUsageServerAuthAndPreCertSigCert.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "warn - serverAuth and MicrosoftDocumentSigning EKU",
			InputFilename:  "subExtKeyUsageServerAuthAndUnknown.pem",
			ExpectedResult: lint.Warn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_sub_cert_eku_check", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
