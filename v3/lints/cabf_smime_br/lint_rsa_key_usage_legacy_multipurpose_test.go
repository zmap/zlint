package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestRSAKeyUsageLegacyMultipurpose(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with digitalSignature KU",
			InputFilename:  "smime/rsa_legacy_digital_signature_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with digitalSignature and contentCommitment KUs",
			InputFilename:  "smime/rsa_multipurpose_digital_signature_content_commitment_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with keyEncipherment KU",
			InputFilename:  "smime/rsa_legacy_key_encipherment_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with keyEncipherment and dataEncipherment KU",
			InputFilename:  "smime/rsa_multipurpose_key_encipherment_data_encipherment_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with digitalSignature, keyEncipherment, contentCommitment, and dataEncipherment KUs",
			InputFilename:  "smime/rsa_legacy_digital_signature_key_encipherment_content_commitment_data_encipherment_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - cert without KUs",
			InputFilename:  "smime/without_subject_alternative_name.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NA - certificate without digitalSignature or keyEncipherment KUs",
			InputFilename:  "smime/rsa_multipurpose_cert_sign_ku.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NE - certificate with valid KUs dated before 2020-09-01",
			InputFilename:  "smime/rsa_multipurpose_valid_ku_august_2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Error - Signing Certificate with unexpected KU",
			InputFilename:  "smime/rsa_legacy_digital_signature_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Error - Key Management Certificate with unexpected KU",
			InputFilename:  "smime/rsa_multipurpose_key_encipherment_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Error - Dual Use Certificate with unexpected KU",
			InputFilename:  "smime/rsa_legacy_digital_signature_key_encipherment_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_rsa_key_usage_legacy_multipurpose", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
