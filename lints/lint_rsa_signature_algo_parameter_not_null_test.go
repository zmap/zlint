package lints

import (
	"fmt"
	"testing"
)

func TestRSAAlgIDNullParams(t *testing.T) {

	testCases := []struct {
		name           string
		filepath       string
		expectedStatus LintStatus
		details        string
	}{
		{
			name:           "cert with missing NULL params",
			filepath:       "rsaAlgIDNoNULLParams.pem",
			expectedStatus: Error,
			details:        "certificate contains RSA public key algorithm identifier missing required NULL parameter",
		},
		{
			name:           "cert with non NULL params",
			filepath:       "rsaKeyWithParameters.pem",
			expectedStatus: Error,
			details:        "certificate contains RSA public key algorithm identifier with non-NULL parameter",
		},
		{
			name:           "cert with NULL params",
			filepath:       "rsawithsha1after2016.pem",
			expectedStatus: Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.filepath)
			result := Lints["e_rsa_encryption_parameter_not_null"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v", tc.expectedStatus, result.Status)
			}

			if result.Details != tc.details {
				t.Errorf("expected error details %q was %q", tc.details, result.Details)
			}
		})
	}
}
