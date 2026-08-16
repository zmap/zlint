package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsMaxValidityPeriod460Days(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - code signing certificate with validity period 460 days or less",
			InputFilename:  "code_signing/validCodeSigningCertificateIssuedAfterMarch1st2026.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - code signing certificate with validity period longer than 460 days",
			InputFilename:  "code_signing/validityPeriodLongerThan460Days.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "NE - code signing certificate issued before March 1st, 2026",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.NE,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_max_validity_period_460_days", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
