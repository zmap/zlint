package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestOnionNotEV(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Onion subject, not EV cert, before util.OnionOnlyEVDate",
			InputFilename:  "dnsNameOnionTLD.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:            "Onion subject, not EV cert, after util.OnionOnlyEVDate",
			InputFilename:   "onionSANNotEV.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `certificate contains one or more .onion subject domains but is not an EV certificate`,
		},
		{
			Name:           "Onion subject, EV cert",
			InputFilename:  "onionSANEV.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_san_dns_name_onion_not_ev_cert", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
