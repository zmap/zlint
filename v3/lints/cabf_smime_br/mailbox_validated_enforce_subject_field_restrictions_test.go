package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestMailboxValidatedEnforceSubjectFieldRestrictions(t *testing.T) {
	testCases := []struct {
		Name          string
		InputFilename string

		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "ok - certificate with commonName",
			InputFilename:  "mailboxValidatedLegacyWithCommonName.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "ok - certificate without mailbox validated policy",
			InputFilename:  "domainValidatedWithEmailCommonName.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "ok - certificate with NotBefore before effective date of lint",
			InputFilename:  "mailboxValidatedLegacyWithCommonNameMay2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:            "error - certificate with countryName",
			InputFilename:   "mailboxValidatedLegacyWithCountryName.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "subject DN contains forbidden field: subject:countryName (2.5.4.6)",
		},
		{
			Name:            "error - certificate containing nonsense subject field (1.2.3.4.5.6.7.8.9.0)",
			InputFilename:   "mailboxValidatedMultipurposeWithNonsenseSubjectField.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "subject DN contains forbidden field: 1.2.3.4.5.6.7.8.9.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mailbox_validated_enforce_subject_field_restrictions", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}

			if tc.ExpectedDetails != "" && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %s, was %s", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
