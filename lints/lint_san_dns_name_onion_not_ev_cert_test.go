package lints

import (
	"fmt"
	"testing"
)

func TestOnionNotEV(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Onion subject, not EV cert, before util.OnionOnlyEVDate",
			InputFilename:  "dnsNameOnionTLD.pem",
			ExpectedResult: NE,
		},
		{
			Name:            "Onion subject, not EV cert, after util.OnionOnlyEVDate",
			InputFilename:   "onionSANNotEV.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `certificate contains one or more .onion subject domains but is not an EV certificate`,
		},
		{
			Name:           "Onion subject, EV cert",
			InputFilename:  "onionSANEV.pem",
			ExpectedResult: Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.InputFilename)
			result := Lints["e_san_dns_name_onion_not_ev_cert"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
