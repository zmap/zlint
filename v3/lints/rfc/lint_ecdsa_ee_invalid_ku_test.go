package rfc

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestECDSAInvalidKU(t *testing.T) {
	testCases := []struct {
		name            string
		filename        string
		expectedStatus  lint.LintStatus
		expectedDetails string
	}{
		{
			name:           "non-ecdsa ee cert",
			filename:       "rsaKeyWithParameters.pem",
			expectedStatus: lint.NA,
		},
		{
			name:           "ecdsa ee cert, valid key usage",
			filename:       "ecdsaP256ValidKUs.pem",
			expectedStatus: lint.Pass,
		},
		{
			name:            "ecdsa ee cert, invalid key usage",
			filename:        "ecdsaP384InvalidKUs.pem",
			expectedStatus:  lint.Notice,
			expectedDetails: "Certificate had unexpected key usage(s): KeyUsageKeyEncipherment",
		},
		{
			name:            "ecdsa ee cert, multiple invalid key usages",
			filename:        "ecdsaP256.pem",
			expectedStatus:  lint.Notice,
			expectedDetails: "Certificate had unexpected key usage(s): KeyUsageCRLSign, KeyUsageCertSign",
		},
	}

	for _, tc := range testCases {
		result := test.TestLint("n_ecdsa_ee_invalid_ku", tc.filename)
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
