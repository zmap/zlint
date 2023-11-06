package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestEdwardsPublicKeyKeyUsages(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with digitalSignature KU",
			InputFilename:  "smime/ed25519_legacy_digital_signature_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with digitalSignature and contentCommitment KUs",
			InputFilename:  "smime/ed25519_multipurpose_digital_signature_content_commitment_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - non-SMIME BR cert",
			InputFilename:  "smime/domainValidatedWithEmailCommonName.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NA - RSA cert",
			InputFilename:  "smime/rsa_strict_digital_signature_ku.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NE - certificate with KU extension dated before 2020-09-01",
			InputFilename:  "smime/ed25519_strict_valid_ku_august_2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Error - Certificate without digitalSignature KU",
			InputFilename:  "smime/ed25519_strict_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_edwardspublickey_key_usages", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
