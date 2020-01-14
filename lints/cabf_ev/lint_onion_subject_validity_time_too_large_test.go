package cabf_ev

import (
	"fmt"
	"testing"

	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/util"
)

func TestTorValidityTooLarge(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Onion subject, long expiry before util.OnionOnlyEVDate",
			InputFilename:  "onionSANLongExpiryPreBallot.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Onion subject, long expiry, after util.OnionOnlyEVDate",
			InputFilename:  "onionSANLongExpiry.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Onion subject, valid expiry",
			InputFilename:  "onionSANGoodExpiry.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", util.TestCaseDir, tc.InputFilename)
			result := lint.Lints["e_onion_subject_validity_time_too_large"].Execute(util.ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}

}
