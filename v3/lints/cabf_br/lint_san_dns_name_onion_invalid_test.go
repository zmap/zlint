package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestOnionNotInvalid(t *testing.T) {
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
			Name:            "non-V2/V3 onion subject, non-EV cert",
			InputFilename:   "onionSANNotEV.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `"zmap.onion" is not a v2 or v3 Tor address`,
		},
		{
			Name:            "non-V2/V3 onion subject, EV cert",
			InputFilename:   "invalidOnionAddress.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `"zmap.onion" is not a v2 or v3 Tor address`,
		},
		{
			Name:            "v2 onion address, non-EV",
			InputFilename:   "onionSANv2NameNonEV.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `"v2cbb2l4lsnpio4q.onion" is a v2 address, but the certificate is not EV`,
		},
		{
			Name:           "v2 onion address, EV",
			InputFilename:  "onionSANv2NameEV.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "misencoded v2 onion address, EV",
			InputFilename:   "onionSANv2NameInvalidEV.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `"v2cbb2l-lsnpio4q.onion" contains invalid characters not permitted within base-32`,
		},
		{
			Name:           "v3 onion address, non-EV",
			InputFilename:  "onionSANv3Name.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_san_dns_name_onion_invalid", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
