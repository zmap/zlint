package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsSubjectCountryName(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - code signing certificate with RSA key size >= 3072",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - code signing certificate with no countryName in subject",
			InputFilename:  "code_signing/no_country_name_in_subject.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - code signing certificate with non-iso countryName in subject",
			InputFilename:  "code_signing/non_iso_country_name_in_subject.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - code signing certificate with only metadata in countryName",
			InputFilename:  "code_signing/only_meta_data_in_country_name.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "NA - code signing certificate with valid EV code signing",
			InputFilename:  "code_signing/valid_ev_code_signing.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_subject_country_name", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
