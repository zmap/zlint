package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsAiaOcspNotHttp(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - code signing certificate with HTTP OCSP URL",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - code signing certificate with non-HTTP OCSP URL",
			InputFilename:  "code_signing/cs_aia_ocsp_not_http.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - code signing certificate with OCSP URL that can not be parsed",
			InputFilename:  "code_signing/cs_aia_malformed_ocsp_url.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "na - code signing certificate with no OCSP URL",
			InputFilename:  "code_signing/cs_aia_no_ocsp.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_aia_ocsp_not_http", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
