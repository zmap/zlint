package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestAIAInternalName(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - aia with valid names",
			InputFilename:  "aiaWithValidNames.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "warn - aia with internal names",
			InputFilename:  "aiaWithInternalNames.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "pass - aia with an IP address",
			InputFilename:  "aiaWithIP.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "na - aia is not present",
			InputFilename:  "akiCritical.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_sub_cert_aia_contains_internal_names", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
