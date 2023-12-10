package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSMIMEStrictAIAHasHTTPOnly(t *testing.T) {
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
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "warn - aia with internal names",
			InputFilename:  "smime/aiaWithLDAPOCSPStrict.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_smime_strict_aia_shall_have_http_only", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
