package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsMaxValidityPeriod39Months(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - code signing certificate with validity period 39 months or less",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - code signing certificate with validity period longer than 39 months",
			InputFilename:  "code_signing/validityPeriodLongerThan39Months.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "NE - code signing certificate issued on or after March 1st, 2026",
			InputFilename:  "code_signing/validCodeSigningCertificateIssuedAfterMarch1st2026.pem",
			ExpectedResult: lint.NE,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_max_validity_period_39_months", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
