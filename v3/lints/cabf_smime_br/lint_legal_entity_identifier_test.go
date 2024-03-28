package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestLegalEntityIdentifier(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "pass - mailbox validated, Legal Entity Identifier not present",
			InputFilename:  "smime/mailboxValidatedLegacyWithCommonName.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "error - mailbox validated, Legal Entity Identifier present",
			InputFilename:   "smime/mailbox_validated_with_lei.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "Legal Entity Identifier extension present",
		},
		{
			Name:            "error - individual validated, Legal Entity Identifier present",
			InputFilename:   "smime/individual_validated_with_lei.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "Legal Entity Identifier extension present",
		},
		{
			Name:            "error - organization validated, Legal Entity Identifier critical",
			InputFilename:   "smime/organization_validated_with_lei_critical.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "Legal Entity Identifier extension present and critical",
		},
		{
			Name:            "error - organization validated, Legal Entity Identifier Role present",
			InputFilename:   "smime/organization_validated_with_lei_role.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "Legal Entity Identifier Role extension present",
		},
		{
			Name:            "error - sponsor validated, Legal Entity Identifier critical",
			InputFilename:   "smime/sponsor_validated_with_lei_critical.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "Legal Entity Identifier extension present and critical",
		},
		{
			Name:            "error - sponsor validated, Legal Entity Identifier Role present",
			InputFilename:   "smime/sponsor_validated_with_lei_role_critical.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "Legal Entity Identifier Role extension present and critical",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_legal_entity_identifier", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}

			if tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %q, was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
