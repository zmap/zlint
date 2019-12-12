package lints

import (
	"fmt"
	"testing"
)

func TestECDSAInvalidKU(t *testing.T) {
	testCases := []struct {
		name            string
		filename        string
		expectedStatus  LintStatus
		expectedDetails string
	}{
		{
			name:           "non-ecdsa ee cert",
			filename:       "rsaKeyWithParameters.pem",
			expectedStatus: NA,
		},
		{
			name:           "ecdsa ee cert, valid key usage",
			filename:       "ecdsaP256ValidKUs.pem",
			expectedStatus: Pass,
		},
		{
			name:            "ecdsa ee cert, invalid key usage",
			filename:        "ecdsaP384InvalidKUs.pem",
			expectedStatus:  Notice,
			expectedDetails: "Certificate had unexpected key usage(s): KeyUsageKeyEncipherment",
		},
		{
			name:            "ecdsa ee cert, multiple invalid key usages",
			filename:        "ecdsaP256.pem",
			expectedStatus:  Notice,
			expectedDetails: "Certificate had unexpected key usage(s): KeyUsageCRLSign, KeyUsageCertSign",
		},
	}

	for _, tc := range testCases {
		inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.filename)
		result := Lints["n_ecdsa_ee_invalid_ku"].Execute(ReadCertificate(inputPath))
		if result.Status != tc.expectedStatus {
			t.Errorf("expected result %v. actual result was %v",
				tc.expectedStatus, result.Status)
		}
		if result.Details != tc.expectedDetails {
			t.Errorf("expected details %q. actual result was %q",
				tc.expectedDetails, result.Details)
		}
	}
}
