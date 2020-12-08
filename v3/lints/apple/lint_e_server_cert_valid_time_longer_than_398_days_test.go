package apple

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestServerCertValidityTooLong(t *testing.T) {
	// Test certificates were created using a small Go program:
	// See https://gist.github.com/cpu/96fad159e6e4db891ee69d225e8a61bc
	testCases := []struct {
		testCert string
		expected lint.LintStatus
	}{
		{
			// Cert issued before Sept 1, 2020 lifetime > 398 days.
			testCert: "eeServerCertValidOver398OldNotBefore.pem",
			expected: lint.NE,
		},
		{
			// Cert issued after Sept 1, 2020 with lifetime <= 397 days.
			testCert: "eeServerCertValidEqual397.pem",
			expected: lint.Pass,
		},
		{
			// Cert issued after Sept 1, 2020 with lifetime > 397 and < 398 days.
			testCert: "eeServerCertValidOver397.pem",
			expected: lint.Pass,
		},
		{
			// Cert issued after Sept 1, 2020 with lifetime == 398 days.
			testCert: "eeServerCertValidEqual398.pem",
			expected: lint.Pass,
		},
		{
			// Cert issued after Sept 1, 2020 with lifetime > 398 days.
			testCert: "eeServerCertValidOver398.pem",
			expected: lint.Error,
		},
		{
			// Cert containing CA basic constraint, should be Not Applicable
			testCert: "caBasicConstCrit.pem",
			expected: lint.NA,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCert, func(t *testing.T) {
			if result := test.TestLint(
				"e_tls_server_cert_valid_time_longer_than_398_days",
				tc.testCert); result.Status != tc.expected {
				t.Errorf("expected result %v was %v", tc.expected, result.Status)
			}
		})
	}
}
