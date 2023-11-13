package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestRegistrationSchemeIDMatchesSubjectNameCountry(t *testing.T) {

	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "pass - organization validated certificate with subject:Name:Country matching subject:organizationIdentifier",
			InputFilename:  "smime/organization_validated_with_matching_country.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - sponsor validated certificate with subject:Name:Country matching subject:organizationIdentifier",
			InputFilename:  "smime/sponsor_validated_with_matching_country.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "error - individual validated certificate",
			InputFilename:   "smime/individual_validated_with_matching_country.pem",
			ExpectedResult:  lint.NA,
		},
		{
			Name:            "error - no country specified in certificate",
			InputFilename:   "smime/organization_validatged_with_no_country_specified.pem",
			ExpectedResult:  lint.NA,
		},
		{
			Name:            "error - organization validated certificate with subject:organizationIdentifier in incorrect format",
			InputFilename:   "smime/organization_validated_with_incorrect_format_identifier.pem",
			ExpectedResult:  lint.NA,
			},
		{
			Name:            "error - organization validated certificate with subject:Name:Country not matching subject:organizationIdentifier",
			InputFilename:   "smime/organization_validated_with_non_matching_country.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "the country code used in the Registration Scheme identifier SHALL match that of the subject:countryName",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_registration_scheme_id_matches_subject_country", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}

			if tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %q, was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}

}
