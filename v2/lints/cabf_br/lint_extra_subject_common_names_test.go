package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestExtraSubjectCommonNames(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "One subject common name",
			InputFilename:  "commonNamesURL.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Multiple subject common names",
			InputFilename:  "extraCommonNames.pem",
			ExpectedResult: lint.Warn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_extra_subject_common_names", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
