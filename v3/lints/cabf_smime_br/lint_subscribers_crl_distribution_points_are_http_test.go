package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubscriberCrlDistributionPointsAreHTTP(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with HTTP CRL distribution point",
			InputFilename:  "smime/subscriber_with_http_crl_distribution_point.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "error - cert without a non-HTTP CRL distribution point",
			InputFilename:  "smime/subscriber_with_non_http_crl_distribution_point.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "error - cert without no HTTP CRL distribution points",
			InputFilename:  "smime/legacy_subscriber_with_non_http_crl_distribution_point.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_subscribers_crl_distribution_points_are_http", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
