package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubjectMalformedDNSARPA(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:            "IPv4 rDNS too few labels",
			InputFilename:   "subjectRDNSIPv4TooFewLabels.pem",
			ExpectedResult:  lint.Warn,
			ExpectedDetails: `name "1.168.192.in-addr.arpa" has too few leading labels (3 vs 4) to be a reverse DNS entry in the ".in-addr.arpa" zone.`,
		},
		{
			Name:            "IPv4 rDNS bad IP",
			InputFilename:   "subjectRDNSIPv4BadIP.pem",
			ExpectedResult:  lint.Warn,
			ExpectedDetails: `the first 4 labels of name "a.b.c.d.in-addr.arpa" did not parse as a reversed IP address`,
		},
		{
			Name:           "IPv4 rDNS reserved IP",
			InputFilename:  "subjectRDNSIPv4ReservedIP.pem",
			ExpectedResult: lint.Pass, // This linter doesn't check that the IP isn't reserved.
		},
		{
			Name:           "IPv4 rDNS OK",
			InputFilename:  "subjectRDNSIPv4GoodIP.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "IPv6 rDNS too few labels",
			InputFilename:   "subjectRDNSIPv6TooFewLabels.pem",
			ExpectedResult:  lint.Warn,
			ExpectedDetails: `name "a.9.8.7.6.5.0.4.0.0.0.3.0.0.0.2.0.0.0.1.0.0.0.0.0.0.0.1.2.3.4.ip6.arpa" has too few leading labels (31 vs 32) to be a reverse DNS entry in the ".ip6.arpa" zone.`,
		},
		{
			Name:            "IPv6 rDNS bad IP",
			InputFilename:   "subjectRDNSIPv6BadIP.pem",
			ExpectedResult:  lint.Warn,
			ExpectedDetails: `the first 32 labels of name "j.a.9.8.7.6.5.0.4.0.0.0.3.0.0.0.2.0.0.0.1.0.0.0.0.0.0.0.1.2.3.4.ip6.arpa" did not parse as a reversed IP address`,
		},
		{
			Name:           "IPv6 rDNS reserved IP",
			InputFilename:  "subjectRDNSIPv6ReservedIP.pem",
			ExpectedResult: lint.Pass, // This linter doesn't check that the IP isn't reserved.
		},
		{
			Name:           "IPv6 rDNS OK",
			InputFilename:  "subjectRDNSIPv6GoodIP.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_subject_contains_malformed_arpa_ip", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
