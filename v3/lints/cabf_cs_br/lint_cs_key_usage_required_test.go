package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsKeyUsageCheck(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - valid code signing certificate with digital signature key usage",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - code signing certificate without required key usage",
			InputFilename:  "code_signing/noDigitalSignatureKeyUsage.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - code signing certificate with prohibited key usage",
			InputFilename:  "code_signing/containsProhibitedKeyUsage.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "warn - code signing certificate with not recommended key usage",
			InputFilename:  "code_signing/containsNotRecommendedKeyUsage.pem",
			ExpectedResult: lint.Warn,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_key_usage_required", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
