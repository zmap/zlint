package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubscriberCrlDistributionPoints(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with a CRL distribution point",
			InputFilename:  "smime/subscriber_with_crl_distribution_points.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "error - cert without a CRL distribution point",
			InputFilename:  "smime/subscriber_no_crl_distribution_points.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "na - certificate has no SMIME BR policy",
			InputFilename:  "smime/with_subject_alternative_name_no_br.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_subscribers_shall_have_crl_distribution_points", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
