package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSingleEmailSubjectIfPresent(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "error - email address present in subjectDN with multiple values",
			InputFilename:  "smime/twoEmailAddressesInSubjectDN.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "pass - email address present in subjectDN with one value",
			InputFilename:  "smime/oneEmailAddressInSubjectDN.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "na - no email address present in subjectDN",
			InputFilename:  "smime/noEmailAddressInSubjectDN.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_single_email_subject_if_present", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
