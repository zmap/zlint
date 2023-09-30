package cabf_smime_br

import (
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
	"testing"
)

func TestSubscriberCrlDistributionPoints(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - mailbox validated, legacy with commonName",
			InputFilename:  "smime/subscriber_with_crl_distribution_points.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - mailbox validated, multipurpose with commonName",
			InputFilename:  "smime/subscriber_no_crl_distribution_points.pem",
			ExpectedResult: lint.Error,
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
