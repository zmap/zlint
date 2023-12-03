package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestKeyUsagePresence(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with KU extension",
			InputFilename:  "smime/rsa_strict_digital_signature_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - non-SMIME BR cert",
			InputFilename:  "smime/domainValidatedWithEmailCommonName.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NE - certificate with KU extension dated before 2020-09-01",
			InputFilename:  "smime/rsa_strict_valid_ku_august_2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Error - certificate without KU extension",
			InputFilename:  "smime/mailboxValidatedLegacyWithCommonName.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_key_usage_presence", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
