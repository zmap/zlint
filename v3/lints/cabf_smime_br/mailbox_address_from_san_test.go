package cabf_smime_br

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

func TestMailboxAddressFromSANLint(t *testing.T) {
	testCases := []struct {
		Name          string
		InputFilename string

		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:          "pass - subject:commonName email address matches san:otherName",
			InputFilename: "WithOtherNameMatched.pem",

			ExpectedResult: lint.Pass,
		},
		{
			Name:          "pass - subject:commonName email address matches san:emailAddress",
			InputFilename: "WithSANEmailMatched.pem",

			ExpectedResult: lint.Pass,
		},
		{
			Name:          "NA - only contains one san:emailAddress value",
			InputFilename: "WithOnlySANEmail.pem",

			ExpectedResult: lint.NA,
		},
		{
			Name:          "NA - only contains one san:otherName value",
			InputFilename: "WithOnlySANOtherName.pem",

			ExpectedResult: lint.NA,
		},
		{
			Name:          "NE - before effective date",
			InputFilename: "NotEffective.pem",

			ExpectedResult: lint.NE,
		},
		{
			Name:          "NA - does not contain smime certificate policy",
			InputFilename: "NotApplicable.pem",

			ExpectedResult: lint.NA,
		},
		{
			Name:          "fail - subject:commonName email address does not match san:otherName",
			InputFilename: "WithOtherNameUnmatched.pem",

			ExpectedResult:  lint.Error,
			ExpectedDetails: "all certificate mailbox addresses must be present in san:emailAddresses or san:otherNames in addition to any other field they may appear",
		},
		{
			Name:          "fail - subject:commonName email address does not match the email value under san:otherName",
			InputFilename: "WithOtherNameIncorrectType.pem",

			ExpectedResult:  lint.Error,
			ExpectedDetails: "all certificate mailbox addresses must be present in san:emailAddresses or san:otherNames in addition to any other field they may appear",
		},
		{
			Name:          "fail - subject:commonName email address does not match san:emailAddress",
			InputFilename: "WithSANEmailUnmatched.pem",

			ExpectedResult:  lint.Error,
			ExpectedDetails: "all certificate mailbox addresses must be present in san:emailAddresses or san:otherNames in addition to any other field they may appear",
		},
		{
			Name:          "fail - subject:commonName email address does not match san:emailAddress, certificate is sponsor validated",
			InputFilename: "sponsorValidatedMultipurposeEmailInSubjectNotInSAN.pem",

			ExpectedResult:  lint.Error,
			ExpectedDetails: "all certificate mailbox addresses must be present in san:emailAddresses or san:otherNames in addition to any other field they may appear",
		},
		{
			Name:          "NA - subject:commonName is personal name, san:emailAddress contains an email",
			InputFilename: "sponsorValidatedMultipurposePersonalNameInCN.pem",

			ExpectedResult: lint.NA,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mailbox_address_shall_contain_an_rfc822_name", "smime/MailboxAddressFromSAN/"+tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}

			if tc.ExpectedResult == lint.Error && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %q, was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
