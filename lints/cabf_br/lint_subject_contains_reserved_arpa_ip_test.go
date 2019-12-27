package cabf_br

import (
	"fmt"
	"testing"
)

const (
	testCaseDir = "../testlint/testCerts/"
)

func TestSubjectReverseDNSARPA(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "IPv4 rDNS too few labels",
			InputFilename:  "subjectRDNSIPv4TooFewLabels.pem",
			ExpectedResult: lint.Pass, // this linter only cares about well formed rDNS for a reserved network address
		},
		{
			Name:           "IPv4 rDNS bad IP",
			InputFilename:  "subjectRDNSIPv4BadIP.pem",
			ExpectedResult: lint.Pass, // this linter only cares about well formed rDNS for a reserved network address
		},
		{
			Name:            "IPv4 rDNS reserved IP",
			InputFilename:   "subjectRDNSIPv4ReservedIP.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `the first 4 labels of name "1.1.168.192.in-addr.arpa" parsed as a reversed IP address in an IANA reserved IP space.`,
		},
		{
			Name:           "IPv4 rDNS OK",
			InputFilename:  "subjectRDNSIPv4GoodIP.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "IPv6 rDNS too few labels",
			InputFilename:  "subjectRDNSIPv6TooFewLabels.pem",
			ExpectedResult: lint.Pass, // this linter only cares about well formed rDNS for a reserved network address
		},
		{
			Name:           "IPv6 rDNS bad IP",
			InputFilename:  "subjectRDNSIPv6BadIP.pem",
			ExpectedResult: lint.Pass, // this linter only cares about well formed rDNS for a reserved network address
		},
		{
			Name:            "IPv6 rDNS reserved IP",
			InputFilename:   "subjectRDNSIPv6ReservedIP.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `the first 32 labels of name "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.e.f.ip6.arpa" parsed as a reversed IP address in an IANA reserved IP space.`,
		},
		{
			Name:           "IPv6 rDNS OK",
			InputFilename:  "subjectRDNSIPv6GoodIP.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.InputFilename)
			result := Lints["e_subject_contains_reserved_arpa_ip"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
