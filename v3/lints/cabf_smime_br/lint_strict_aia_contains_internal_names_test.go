package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSMIMEStrictAIAInternalName(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - aia with valid names",
			InputFilename:  "smime/aiaWithValidNamesStrict.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "warn - aia with internal names",
			InputFilename:  "smime/aiaWithInternalNamesStrict.pem",
			ExpectedResult: lint.Warn,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_smime_strict_aia_contains_internal_names", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
