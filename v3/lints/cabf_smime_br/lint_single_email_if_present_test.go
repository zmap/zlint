package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSingleEmailIfPresent(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with one email address",
			InputFilename:  "smime/single_email_present.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - cert with no email addresses",
			InputFilename:  "smime/no_email_present.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Error - cert with multiple email addresses",
			InputFilename:  "smime/multiple_email_present.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_single_email_if_present", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
