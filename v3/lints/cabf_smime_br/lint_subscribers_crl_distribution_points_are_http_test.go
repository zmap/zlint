package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubscriberCrlDistributionPointsAreHTTP(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "pass - strict cert with only HTTP CRL distribution points",
			InputFilename:  "smime/strict_subscriber_with_http_crl_distribution_point.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "error - strict cert with a non-HTTP CRL distribution point",
			InputFilename:   "smime/strict_subscriber_with_non_http_crl_distribution_point.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "SMIME certificate contains invalid URI scheme in CRL distribution point",
		},
		{
			Name:            "error - legacy cert with no HTTP CRL distribution points",
			InputFilename:   "smime/legacy_subscriber_with_non_http_crl_distribution_point.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "SMIME certificate contains no HTTP URI schemes as CRL distribution points",
		},
		{
			Name:           "pass - legacy cert with HTTP and non-HTTP CRL distribution points",
			InputFilename:  "smime/legacy_subscriber_with_mixed_crl_distribution_points.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "error - strict cert with HTTP and non-HTTP CRL distribution points",
			InputFilename:   "smime/strict_subscriber_with_mixed_crl_distribution_points.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "SMIME certificate contains invalid URI scheme in CRL distribution point",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_subscribers_crl_distribution_points_are_http", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}

			if tc.ExpectedDetails != "" && tc.ExpectedDetails != result.Details {
				t.Errorf("expected details: %s, was %s", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
