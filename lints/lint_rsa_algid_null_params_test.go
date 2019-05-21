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
	}{
		{
			name:           "cert with missing NULL params",
			filepath:       "rsaAlgIDNoNULLParams.pem",
			expectedStatus: Fatal,
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
			result := Lints["rsa_algid_null_params"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v", tc.expectedStatus, result.Status)
			}
		})
	}
}
