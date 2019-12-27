package cabf_br

import (
	"fmt"
	"testing"

	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/util"
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
			inputPath := fmt.Sprintf("%s%s", util.TestCaseDir, tc.InputFilename)
			result := lint.Lints["w_extra_subject_common_names"].Execute(util.ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
