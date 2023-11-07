package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestECPublicKeyKeyUsage(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with digitalSignature KU",
			InputFilename:  "smime/ec_legacy_digital_signature_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with digitalSignature and contentCommitment KUs",
			InputFilename:  "smime/ec_multipurpose_digital_signature_content_commitment_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with keyAgreement KU",
			InputFilename:  "smime/ec_strict_key_agreement_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with keyAgreement and encipherOnly KUs",
			InputFilename:  "smime/ec_legacy_key_agreement_encipher_only_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with keyAgreement and decipherOnly KUs",
			InputFilename:  "smime/ec_multipurpose_key_agreement_decipher_only.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with digitalSignature, keyAgreement, contentCommitment, and encipherOnly KUs",
			InputFilename:  "smime/ec_strict_digital_signature_key_agreement_content_commitment_encipher_only_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - cert with digitalSignature, keyAgreement, contentCommitment, and decipherOnly KUs",
			InputFilename:  "smime/ec_legacy_digital_signature_key_agreement_content_commitment_decipher_only_ku.pem",
			ExpectedResult: lint.Pass,
		}, {
			Name:           "NA - cert without KUs",
			InputFilename:  "smime/without_subject_alternative_name.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NA - Certificate without digitalSignature or keyAgreement KUs",
			InputFilename:  "smime/ec_strict_cert_sign_ku.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NE - certificate with valid KUs dated before 2020-09-01",
			InputFilename:  "smime/ec_multipurpose_valid_ku_august_2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Error - Signing Certificate with unexpected KU",
			InputFilename:  "smime/ec_strict_digital_signature_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Error - Key Management Certificate with unexpected KU",
			InputFilename:  "smime/ec_legacy_key_agreement_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Error - Dual Use Certificate with unexpected KU",
			InputFilename:  "smime/ec_multipurpose_digital_signature_key_agreement_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ecpublickey_key_usages", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
