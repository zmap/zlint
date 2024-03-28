package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubscriberSubjectAlternativeNameShallBePresent(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with SAN",
			InputFilename:  "smime/with_subject_alternative_name.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "error - cert without SAN",
			InputFilename:  "smime/without_subject_alternative_name.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "na - certificate has no SMIME BR policy",
			InputFilename:  "smime/with_subject_alternative_name_no_br.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_san_shall_be_present", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
