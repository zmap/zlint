package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsProhibitedEcdsaCurve(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - valid code signing certificate with p256 ecdsa curve",
			InputFilename:  "code_signing/valid_ecdsa_code_signing_certificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - valid code signing certificate with rsa key",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "fail - code signing certificate with p224 ecdsa curve",
			InputFilename:  "code_signing/prohibited_p224_ecdsa_curve.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_ecdsa_prohibited_curve", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
