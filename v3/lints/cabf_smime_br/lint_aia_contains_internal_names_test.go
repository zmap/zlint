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
			Name:           "warn - aia with internal names in AIA OCSP ",
			InputFilename:  "smime/aiaWithInternalNamesStrict.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "warn - aia with internal names in AIA CA issuers ",
			InputFilename:  "smime/aiaWithInternalNamesCaIssuersStrict.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "warn - aia with valid names, one is ldap",
			InputFilename:  "smime/aiaWithLDAPOCSPStrict.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - aia with IP address in host part of the URL",
			InputFilename:  "smime/aiaWithIPAddress.pem",
			ExpectedResult: lint.Pass,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_smime_aia_contains_internal_names", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
