package lints

import (
	"fmt"
	"testing"
)

func TestTorValidityTooLarge(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult LintStatus
	}{
		{
			Name:           "Onion subject, long expiry before util.OnionOnlyEVDate",
			InputFilename:  "onionSANLongExpiryPreBallot.pem",
			ExpectedResult: NE,
		},
		{
			Name:           "Onion subject, long expiry, after util.OnionOnlyEVDate",
			InputFilename:  "onionSANLongExpiry.pem",
			ExpectedResult: Error,
		},
		{
			Name:           "Onion subject, valid expiry",
			InputFilename:  "onionSANGoodExpiry.pem",
			ExpectedResult: Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.InputFilename)
			result := Lints["e_onion_subject_validity_time_too_large"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}

}
