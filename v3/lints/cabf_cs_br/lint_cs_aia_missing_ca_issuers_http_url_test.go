package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsAiaMissingCaIssuersHttpUrl(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - code signing certificate with HTTP caIssuers URL",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - code signing certificate with non-HTTP caIssuers URL",
			InputFilename:  "code_signing/cs_aia_ca_issuers_not_http.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - code signing certificate with no caIssuers URL",
			InputFilename:  "code_signing/cs_aia_no_ca_issuers.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - code signing certificate with caIssuers URL that can not be parsed",
			InputFilename:  "code_signing/cs_aia_malformed_ca_issuers.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_aia_missing_ca_issuers_http_url", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
