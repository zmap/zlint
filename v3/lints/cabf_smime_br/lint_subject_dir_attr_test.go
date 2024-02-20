package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSMIMESubjectDirAttributes(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - no subject dir attributes extension",
			InputFilename:  "smime/mailboxValidatedStrictWithCommonName.pem",
			ExpectedResult: lint.Pass,
		},
		// A negative test case is hard to construct because neither the x509 package
		// nor OpenSSL support writing the subject directory attributes extension.
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_strict_multipurpose_smime_ext_subject_directory_attr", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
