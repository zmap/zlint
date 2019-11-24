package lints

import (
	"fmt"
	"testing"
)

func TestExtraSubjectCommonNames(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult LintStatus
	}{
		{
			Name:           "One subject common name",
			InputFilename:  "commonNamesURL.pem",
			ExpectedResult: Pass,
		},
		{
			Name:           "Multiple subject common names",
			InputFilename:  "extraCommonNames.pem",
			ExpectedResult: Warn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.InputFilename)
			result := Lints["w_extra_subject_common_names"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
