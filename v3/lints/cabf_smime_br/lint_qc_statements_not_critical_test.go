package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSMIMEQCStatementsNotCritical(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "N/A - no qcStatements extension",
			InputFilename:  "smime/legacyAiaOneHTTPOneLdap.pem",
			ExpectedResult: lint.NA,
		},
		// Test cases are hard to construct because:
		//  - while the zcrypto x509.Certificate struct has an QCStatements field, it is ignored when creating a certificate
		//  - OpenSSL doesn't support writing the qcStatements extension.
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_smime_qc_statements_must_not_be_critical", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
