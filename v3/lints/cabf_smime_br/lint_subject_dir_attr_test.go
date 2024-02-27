package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSMIMESubjectDirAttributes(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "pass - no subject dir attributes extension",
			InputFilename:  "smime/mailboxValidatedStrictWithCommonName.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "fail - subject dir attributes extension present",
			InputFilename:   "smime/subject_dir_attributes_present.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "subject direcotry attribute present",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_strict_multipurpose_smime_ext_subject_directory_attr", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}

			if tc.ExpectedDetails != "" && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %s, was %s", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
