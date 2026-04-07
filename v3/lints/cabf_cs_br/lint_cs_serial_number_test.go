package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsSerialNumber(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - code signing certificate with valid serial number",
			InputFilename:  "code_signing/valid_code_signing_certificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - code signing certificate with serial number less than 64 bits",
			InputFilename:  "code_signing/serial_number_less_than_64.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - code signing certificate with negative serial number",
			InputFilename:  "code_signing/negative_serial_number.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_serial_number", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
