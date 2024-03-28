package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSMIMELegacyAIAHasOneHTTP(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - aia with one ldap URI and one HTTP in each method",
			InputFilename:  "smime/legacyAiaOneHTTPOneLdap.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "error - aia with only ldap URIs HTTP in each method",
			InputFilename:  "smime/legacyAiaLdapOnly.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_smime_legacy_aia_shall_have_one_http", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
