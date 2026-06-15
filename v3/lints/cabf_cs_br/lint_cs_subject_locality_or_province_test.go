package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsSubjectLocalityOrProvince(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - locality and province both present",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - locality present, province absent",
			InputFilename:  "code_signing/cs_locality_only.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - province present, locality absent",
			InputFilename:  "code_signing/cs_province_only.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - locality and province both absent",
			InputFilename:  "code_signing/cs_no_locality_no_province.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_subject_locality_or_province_required", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
