package rfc

import (
	"testing"

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestSigAlgMismatch(t *testing.T) {
	testCases := []struct {
		name           string
		filepath       string
		expectedStatus lint.LintStatus
	}{
		{
			name:           "error cert with mismatching signature algorithms (bad OID)",
			filepath:       "mismatchingSigAlgsBadOID.pem",
			expectedStatus: lint.Error,
		},
		{
			name:           "error cert with mismatching signature algorithms (bad parameters)",
			filepath:       "mismatchingSigAlgsBadParams.pem",
			expectedStatus: lint.Error,
		},
		{
			name:           "pass cert with matching signature algorithms",
			filepath:       "ecdsaP256.pem",
			expectedStatus: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := test.TestLint("e_cert_sig_alg_not_match_tbs_sig_alg", tc.filepath)
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v", tc.expectedStatus, result.Status)
			}
		})
	}
}
