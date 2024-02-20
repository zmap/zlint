package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestAuthorityKeyInfoCorrect(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert has keyIdentifier",
			InputFilename:  "smime/authority_key_identifier_valid.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "Error - cert has serial and DirName",
			InputFilename:  "smime/authority_key_identifier_invalid.pem",
			ExpectedResult: lint.Error,
		}}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_authority_key_identifier_correct", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
