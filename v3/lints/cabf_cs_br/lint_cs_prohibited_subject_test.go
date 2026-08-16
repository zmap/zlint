package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsSubjectProhibited(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - code signing certificate with RSA key size >= 3072",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - non-ev code signing certificate with prohibited subject::domainComponent",
			InputFilename:  "code_signing/includes_prohibited_domain_component.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "fail - ev code signing certificate with prohibited subject::domainComponent",
			InputFilename:  "code_signing/ev_includes_prohibited_domain_component.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_subject_prohibited", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
